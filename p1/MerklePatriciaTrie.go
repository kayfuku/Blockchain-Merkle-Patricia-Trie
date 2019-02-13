package p1

import (
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"

	"golang.org/x/crypto/sha3"
)

type Flag_value struct {
	// ASCII value array.
	encoded_prefix []uint8
	// If the node is Ext, 'value' is hash of the next node.
	// If the node is Leaf, 'value' is the string value inserted.
	value string
}

type Node struct {
	// 0: Null, 1: Branch, 2: Ext or Leaf
	node_type int
	// If the node is not Branch, 'branch_value' is default.
	branch_value [17]string
	// If the node is Branch, 'flag_value' is default.
	flag_value Flag_value
}

type MerklePatriciaTrie struct {
	// K: Node's hash value, V: Node
	db map[string]Node
	// hash value of the root node
	root string
}

func NewMPT() *MerklePatriciaTrie {
	// Initialize node.
	nullNode := createNewLeafOrExtNode(0, nil, "")
	// flagValue := Flag_value{encoded_prefix: nil, value: ""}
	// node := Node{node_type: 0, branch_value: [17]string{}, flag_value: flagValue}

	mpt := &MerklePatriciaTrie{}
	mpt.db = map[string]Node{}
	mpt.root = putNodeInDb(nullNode, mpt.db)

	return mpt
}

// Convert key string to hex value array and append 16.
// If the key is "", then key_hex is 16.
func convert_string_to_hex(key string) []uint8 {
	length := 2*len(key) + 1
	key_hex := make([]uint8, length)
	for i, r := range key {
		key_hex[i*2] = uint8(r / 16)
		key_hex[i*2+1] = uint8(r % 16)
	}
	key_hex[length-1] = 16
	return key_hex
}

func compact_encode(hex_array []uint8) []uint8 {
	// TODO
	if len(hex_array) == 0 {
		return []uint8{}
	}
	var isLeaf uint8 = 0
	if hex_array[len(hex_array)-1] == 16 {
		isLeaf = 1
	}
	if isLeaf == 1 {
		hex_array = hex_array[:len(hex_array)-1]
	}
	var isOdd uint8 = uint8(len(hex_array) % 2)
	var flagInHexArray uint8 = 2*isLeaf + isOdd
	if isOdd == 1 {
		hex_array = append([]uint8{flagInHexArray}, hex_array...)
	} else {
		hex_array = append(append([]uint8{flagInHexArray}, 0), hex_array...)
	}
	// 'hex_array' now has an even length whose first nibble is the 'flagInHexArray'.
	length := len(hex_array) / 2
	encoded_prefix := make([]uint8, length)
	p := 0
	for i := 0; i < len(hex_array); i += 2 {
		encoded_prefix[p] = 16*hex_array[i] + hex_array[i+1]
		p++
	}

	return encoded_prefix
}

// If Leaf, ignore 16 at the end
func compact_decode(encoded_arr []uint8) []uint8 {
	// TODO
	if len(encoded_arr) == 0 {
		return []uint8{}
	}

	length := len(encoded_arr) * 2
	hex_array := make([]uint8, length)
	for i, ascii := range encoded_arr {
		hex_array[i*2] = ascii / 16
		hex_array[i*2+1] = ascii % 16
	}

	// Remove prefix and return hex array without 16.
	// If hex hex_array[0] is even, then cut first two. If hex hex_array[0] is odd, then cut first one.
	cut := 2 - hex_array[0]&1
	return hex_array[cut:]
}

// Get function returning two values for testing.
func (mpt *MerklePatriciaTrie) Get(key string) (string, error) {
	// TODO
	if key == "" {
		return "", nil
	}

	keySearch := convert_string_to_hex(key)

	rootNode := mpt.db[mpt.root]
	encodedPrefix := rootNode.flag_value.encoded_prefix
	keyMPT := compact_decode(encodedPrefix)

	return get_helper(rootNode, keyMPT, keySearch, mpt.db), nil
}

// // Get function.
// func (mpt *MerklePatriciaTrie) Get(key string) string {
// 	// TODO
// 	if key == "" {
// 		return ""
// 	}

// 	keySearch := convert_string_to_hex(key)

// 	rootNode := mpt.db[mpt.root]
// 	encodedPrefix := rootNode.flag_value.encoded_prefix
// 	keyMPT := compact_decode(encodedPrefix)

// 	return get_helper(rootNode, keyMPT, keySearch, mpt.db)
// }
func get_helper(node Node, keyMPT, keySearch []uint8, db map[string]Node) string {

	nodeType := node.node_type
	switch nodeType {
	case 0:
		// Null node

		return ""
	case 1:
		// Branch

		if keySearch[0] == 16 {
			// Case B-1. Insert("a"), Insert("aa"), Get("a"), stack 2. keyMPT: [], keySearch: [16], matchLen: 2
			// Case B-3. Insert("aa"), Insert("a"), Get("aa"), stack 2. keyMPT: [], keySearch: [16], matchLen: 2

			return node.branch_value[16]
		}

		if nextNode, ok := db[node.branch_value[keySearch[0]]]; ok {
			// There is the next node in the Branch.
			// 'node' is now the next node.
			// Case B-1. Insert("a"), Insert("aa"), Get("aa"), stack 2. keyMPT: [], keySearch: [6 1 16]
			// Case B-2. Insert("a"), Insert("b"), Get("a"), stack 2. keyMPT: [], keySearch: [1 16]
			// Case B-3. Insert("aa"), Insert("a"), Get("aa"), stack 2. keyMPT: [], keySearch: [6 1 16]
			// Case C. Insert("a"), Insert("p"), Get("a"), stack 1. keyMPT: [], keySearch: [6 1 16]
			// Case D-1. Insert("a"), Insert("p"), Insert("abc"), Get("abc") stack 1. keyMPT: [], keySearch: [6 1 6 2 6 3 16]
			// Case D-1. Insert("a"), Insert("p"), Insert("abc"), Get("abc") stack 3. keyMPT: [], keySearch: [6 2 6 3 16]
			// Case D-3. Insert("a"), Insert("p"), Insert("A"), Get("A") stack 1. keyMPT: [], keySearch: [4 1 16]
			encodedPrefix := nextNode.flag_value.encoded_prefix
			keyMPT := compact_decode(encodedPrefix)
			return get_helper(nextNode, keyMPT, keySearch[1:], db)
		}

		// There is no link in the Branch.
		return ""

	case 2:
		// Ext or Leaf

		matchLen := prefixLen(keySearch, keyMPT)

		if matchLen != 0 {

			if len(keySearch) <= len(keyMPT) {
				// keySearch is shorter than keyMPT.
				return ""
			}

			if firstDigit := getFirstDigitOfAscii(node.flag_value.encoded_prefix); firstDigit == 0 || firstDigit == 1 || firstDigit == 2 {
				// 'node' is Ext node.
				// Case B-1. Insert("a"), Insert("aa"), Get("aa"), stack 1. keyMPT: [6 1], keySearch: [6 1 6 1 16], matchLen: 2
				// Case B-1. Insert("a"), Insert("aa"), Get("a"), stack 1. keyMPT: [6 1], keySearch: [6 1 16], matchLen: 2
				// Case B-2. Insert("a"), Insert("b"), Get("a"), stack 1. keyMPT: [6], keySearch: [6 1 16], matchLen: 1
				// Case B-3. Insert("aa"), Insert("a"), Get("aa"), stack 1. keyMPT: [6 1], keySearch: [6 1 6 1 16], matchLen: 2
				// Case D-1. Insert("a"), Insert("p"), Insert("abc"), Get("abc") stack 2. keyMPT: [1], keySearch: [1 6 2 6 3 16], matchLen: 1

				node = db[node.flag_value.value]
				// 'node' is now Branch node next to the Ext node.
				return get_helper(node, keyMPT[matchLen:], keySearch[matchLen:], db)
			}

			// 'node' is Leaf node.

			if keySearch[matchLen] == 16 && len(keyMPT) == matchLen {
				// Exact match.
				// Case A. Insert("a"), Get("a"). keyMPT: [6 1], keySearch: [6 1 16], matchLen: 2
				// Case B-1. Insert("a"), Insert("aa"), Get("aa"), stack 3. keyMPT: [1], keySearch: [1 16], matchLen: 1
				// Case B-3. Insert("aa"), Insert("a"), Get("aa"), stack 3. keyMPT: [1], keySearch: [1 16], matchLen: 1
				// Case C.   Insert("a"), Insert("p"), Get("a"), stack 2. keyMPT: [1], keySearch: [1 16], matchLen: 1
				// Case D-1. Insert("a"), Insert("p"), Insert("abc"), Get("abc") stack 4. keyMPT: [2 6 3], keySearch: [2 6 3 16], matchLen: 3
				// Case D-3. Insert("a"), Insert("p"), Insert("A"), Get("A") stack 2. keyMPT: [1], keySearch: [1 16]
				return node.flag_value.value
			}

			// 'node' is Leaf node and keySearch is shorter or longer than keyMPT.
			return ""

		} else if matchLen == 0 {

			if keySearch[matchLen] == 16 {
				// Case B-2. Insert("a"), Insert("b"), Get("a"), stack 3. keyMPT: [], keySearch: [16], matchLen: 0
				return node.flag_value.value
			}
			return ""
		}

	}

	return ""
}

func (mpt *MerklePatriciaTrie) Insert(key string, new_value string) {
	// TODO
	if key == "" {
		return
	}

	keySearch := convert_string_to_hex(key)

	db := mpt.db
	rootNode := db[mpt.root]
	encodedPrefix := rootNode.flag_value.encoded_prefix
	keyMPT := compact_decode(encodedPrefix)

	newRootNode := insert_helper(rootNode, keyMPT, keySearch, new_value, db)
	// delete(db, mpt.root)
	mpt.root = putNodeInDb(newRootNode, db)

	return
}
func insert_helper(node Node, keyMPT, keySearch []uint8, new_value string, db map[string]Node) Node {

	nodeType := node.node_type
	switch nodeType {
	case 0:
		// Null node

		// Create a new Leaf node.
		node = createNewLeafOrExtNode(2, keySearch, new_value)
		return node
	case 1:
		// Branch

		if keySearch[0] == 16 {
			// Case E-2. stack 3. keyMPT: [], keySearch: [16]
			node.branch_value[16] = new_value
			return node
		}

		if nextNode, ok := db[node.branch_value[keySearch[0]]]; ok {
			// There is a link in the Branch.
			// 'nextNode' is the node next to the Branch. It could be Leaf, Ext, or Branch.
			// Case D-1. Insert("a"), Insert("p"), Insert("abc"), Get("abc"),
			// stack 1. keyMPT: [], keySearch: [6 1 6 2 6 3 16]
			// Case E-1. Insert("p"), Insert("aaaaa"), Insert("aaaap"), Insert("aa"), Get("aa"),
			// stack 1. keyMPT: [], keySearch: [6 1 6 1 16]
			// Case E-2. Insert("p"), Insert("aaaaa"), Insert("aaaap"), Insert("aaaa"), Get("aaaa"),
			// stack 1. keyMPT: [], keySearch: [6 1 6 1 6 1 6 1 16]
			// Case C-2. stack 2. C-3. stack 2. E-4b. stack 1.
			// Case D-2. stack 1.
			encodedPrefix := nextNode.flag_value.encoded_prefix
			keyMPT := compact_decode(encodedPrefix)

			nextNode = insert_helper(nextNode, keyMPT, keySearch[1:], new_value, db)
			delete(db, node.hash_node())
			node.branch_value[keySearch[0]] = putNodeInDb(nextNode, db)
			return node
		}

		// There is no link in the Branch.
		// Case D-3. Insert("a"), Insert("p"), Insert("A"), Get("A"). keyMPT: [], keySearch: [4 1 16]
		leafNode := createNewLeafOrExtNode(2, keySearch[1:], new_value)
		delete(db, node.hash_node())
		node.branch_value[keySearch[0]] = putNodeInDb(leafNode, db)
		return node

	case 2:
		// Ext or Leaf

		matchLen := prefixLen(keySearch, keyMPT)

		if matchLen != 0 {

			if firstDigit := getFirstDigitOfAscii(node.flag_value.encoded_prefix); firstDigit == 0 || firstDigit == 1 || firstDigit == 2 {
				// 'node' is Ext.

				if keySearch[matchLen] == 16 && len(keyMPT) > matchLen {
					// Partial match. keySearch is done and keyMPT is left.
					// Case E-1. stack 2. keyMPT: [1 6 1 6 1 6 1], keySearch: [1 6 1 16], matchLen: 3
					// E-1b. stack 2. keyMPT: [1 6 1 6 1], keySearch: [1 16], matchLen: 1
					extNode1 := createNewLeafOrExtNode(2, keyMPT[:matchLen], "")
					branchNode := Node{node_type: 1, branch_value: [17]string{}}
					extNode2 := createNewLeafOrExtNode(2, keyMPT[matchLen+1:], node.flag_value.value)
					delete(db, node.hash_node())
					branchNode.branch_value[keyMPT[matchLen]] = putNodeInDb(extNode2, db)
					branchNode.branch_value[16] = new_value
					extNode1.flag_value.value = putNodeInDb(branchNode, db)

					return extNode1
				}

				if len(keyMPT) == matchLen {
					// Partial match. keyMPT is done.
					// Case E-2. stack 2. keyMPT: [1 6 1 6 1 6 1], keySearch: [1 6 1 6 1 6 1 16], matchLen: 7
					// Case E-3. stack 2. keyMPT: [1 6 1 6 1 6 1], keySearch: [1 6 1 6 1 6 1 4 1 16], matchLen: 7
					// Case C-2. stack 1. C-3.
					branchNode := db[node.flag_value.value]
					branchNode = insert_helper(branchNode, nil, keySearch[matchLen:], new_value, db)
					delete(db, node.hash_node())
					node.flag_value.value = putNodeInDb(branchNode, db)

					return node
				}

			}

			// 'node' is Leaf.

			if keySearch[matchLen] == 16 && len(keyMPT) == matchLen {
				// Case A (Exact match). keyMPT: [6 1], keySearch: [6 1 16], matchLen: 2
				node.flag_value.value = new_value
				return node
			}

			// Case B-1 (Prefix match).
			// stack 1. keyMPT: [6 1], keySearch: [6 1 6 1 16], matchLen: 2
			// Case B-2 (Prefix match).
			// stack 1. keyMPT: [6 1], keySearch: [6 2 16], matchLen: 1
			// Case B-3 (Prefix match).
			// stack 1. keyMPT: [6 1 6 1] , keySearch: [6 1 16], matchLen: 2
			// Case D-1. stack 2. keyMPT: [1], keySearch: [1 6 2 6 3 16], matchLen: 1
			extNode := createNewLeafOrExtNode(2, keyMPT[:matchLen], node.flag_value.value)

			branchNode := insert_helper(extNode, keyMPT[matchLen:], keySearch[matchLen:], new_value, db)
			extNode.flag_value.value = putNodeInDb(branchNode, db)
			return extNode

		} else if matchLen == 0 {

			if keySearch[matchLen] == 16 && len(keyMPT) == 0 {
				// C-2. stack 3.
				node.flag_value.value = new_value
				delete(db, node.hash_node())
				return node
			}

			branchNode := Node{node_type: 1, branch_value: [17]string{}}
			if len(keyMPT) == 0 {
				// Case B-1 (Prefix match).
				// stack 2. keyMPT: [], keySearch: [6 1 16], matchLen: 0
				// Case D-1.
				// stack 3. keyMPT: [], keySearch: [6 2 6 3 16], matchLen: 0
				// D-3. stack 3.
				leafNode := createNewLeafOrExtNode(2, keySearch[matchLen+1:], new_value)
				branchNode.branch_value[keySearch[matchLen]] = putNodeInDb(leafNode, db)
				branchNode.branch_value[16] = node.flag_value.value

			} else if keySearch[matchLen] == 16 {
				// Case B-3 (Prefix match).
				// stack 2. keyMPT: [6 1], keySearch: [16], matchLen: 0
				leafNode := createNewLeafOrExtNode(2, append(keyMPT[matchLen+1:], 16), node.flag_value.value)
				branchNode.branch_value[keyMPT[matchLen]] = putNodeInDb(leafNode, db)
				branchNode.branch_value[16] = new_value

			} else {
				// Case B-2 (Prefix match).
				// stack 2. keyMPT: [1], keySearch: [2 16], matchLen: 0
				// Case C (Mismatch).
				// keyMPT: [6 1], keySearch: [7 0 16], matchLen: 0
				// D-2. stack 2.
				leafNode := createNewLeafOrExtNode(2, keySearch[matchLen+1:], new_value)
				delete(db, node.hash_node())
				branchNode.branch_value[keySearch[matchLen]] = putNodeInDb(leafNode, db)

				if _, ok := db[node.flag_value.value]; ok {
					if len(keyMPT[1:]) != 0 {
						// E-4.
						extNode := createNewLeafOrExtNode(2, keyMPT[1:], node.flag_value.value)
						branchNode.branch_value[keyMPT[0]] = putNodeInDb(extNode, db)
					} else {
						// E-4b
						branchNode.branch_value[keyMPT[0]] = node.flag_value.value
					}

				} else {
					// B-2, C. D-2.
					leafNode = createNewLeafOrExtNode(2, append(keyMPT[matchLen+1:], 16), node.flag_value.value)
					branchNode.branch_value[keyMPT[matchLen]] = putNodeInDb(leafNode, db)
				}
			}

			return branchNode
		}
	default:

	}

	return node
}

func (mpt *MerklePatriciaTrie) Delete(key string) string {
	// TODO
	if key == "" {
		return ""
	}

	keySearch := convert_string_to_hex(key)

	db := mpt.db
	rootNode := db[mpt.root]
	encodedPrefix := rootNode.flag_value.encoded_prefix
	keyMPT := compact_decode(encodedPrefix)

	newRootNode, ret := delete_helper(rootNode, keyMPT, keySearch, db)
	mpt.root = putNodeInDb(newRootNode, db)

	return ret
}
func delete_helper(node Node, keyMPT, keySearch []uint8, db map[string]Node) (Node, string) {

	nodeType := node.node_type
	switch nodeType {
	case 0:
		// 'node' is Null node.

		return node, ""
	case 1:
		// 'node' is Branch.

		if keySearch[0] == 16 {
			// Del-4. B-3. stack 2. Insert("aa"), Insert("a"), Delete("a"), stack 2. keyMPT: [], keySearch: [16], matchLen: 2
			// Del-5. stack 2.
			// Del-7.
			node.branch_value[16] = ""

			if b, oneValue, index := getOnlyOneValueInBranch(node); b {
				// Only one value in the Branch. Rebalance.
				// The value is a link to the next node.
				// Del-4. stack 2.
				// Del-7.

				leftNode := db[oneValue]
				// 'leftNode' could be Leaf, Ext, or Branch.
				if leftNode.node_type == 1 {
					// 'leftNode' is Branch.
					// Del-7.
					delete(db, node.hash_node())
					extNode := createNewLeafOrExtNode(2, []uint8{index}, oneValue)
					return extNode, ""

				} else if firstDigit := getFirstDigitOfAscii(leftNode.flag_value.encoded_prefix); firstDigit == 3 || firstDigit == 4 || firstDigit == 5 {
					// leftNode is Leaf.
					leftNode.flag_value.encoded_prefix = compact_encode(
						append([]uint8{index},
							append(compact_decode(leftNode.flag_value.encoded_prefix), 16)...))
				} else {
					// leftNode is Ext or Branch.
					// Del-5.
					leftNode.flag_value.encoded_prefix = compact_encode(
						append([]uint8{index}, compact_decode(leftNode.flag_value.encoded_prefix)...))

				}

				return leftNode, ""
			}

			return node, ""
		}

		if nextNode, ok := db[node.branch_value[keySearch[0]]]; ok {
			// There is the next node in the Branch.
			// 'nextNode' is the node next to the Branch.
			// Del-3. Insert("a"), Insert("aa"), Delete("aa"), stack 2. keyMPT: [], keySearch: [6 1 16]
			// Del-8. stack 1. stack 2.
			encodedPrefix := nextNode.flag_value.encoded_prefix
			keyMPT := compact_decode(encodedPrefix)

			retNode, ret := delete_helper(nextNode, keyMPT, keySearch[1:], db)
			if retNode.node_type == 0 {
				// 'retNode' is Null node.
				node.branch_value[keySearch[0]] = ""

				if b, oneValue, index := getOnlyOneValueInBranch(node); b {
					// Only one value in the Branch. Rebalance.
					if node.branch_value[16] != "" {
						// The value is in the last 16th elem.
						// Del-3. stack 2.
						leafNode := createNewLeafOrExtNode(2, []uint8{16}, oneValue)
						return leafNode, ret
					}
					// The value is a link to the next node.
					// Del-1. stack 2. Del-8.

					leftNode := db[oneValue]
					// 'leftNode' could be Leaf, Ext, or Branch.
					if leftNode.node_type == 1 {
						// 'leftNode' is Branch.
						// Del-6. Del-8.
						delete(db, node.hash_node())
						extNode := createNewLeafOrExtNode(2, []uint8{index}, oneValue)
						return extNode, ret

					} else if firstDigit := getFirstDigitOfAscii(leftNode.flag_value.encoded_prefix); firstDigit == 3 || firstDigit == 4 || firstDigit == 5 {
						// leftNode is Leaf.
						leftNode.flag_value.encoded_prefix = compact_encode(
							append([]uint8{index},
								append(compact_decode(leftNode.flag_value.encoded_prefix), 16)...))
					} else {
						// leftNode is Ext.
						// Del-2.
						leftNode.flag_value.encoded_prefix = compact_encode(
							append([]uint8{index}, compact_decode(leftNode.flag_value.encoded_prefix)...))
					}

					return leftNode, ret
				}

				return node, ret
			}

			// Del-8. 'retNode' is Ext.
			node.branch_value[keySearch[0]] = putNodeInDb(retNode, db)

			return node, ret
		}

		// There is no link in the Branch.
		return node, "path_not_found"

	case 2:
		// 'node' is Ext or Leaf.

		matchLen := prefixLen(keySearch, keyMPT)

		if matchLen != 0 {

			if firstDigit := getFirstDigitOfAscii(node.flag_value.encoded_prefix); firstDigit == 0 || firstDigit == 1 || firstDigit == 2 {
				// 'node' is Ext node.
				// Del-3. Insert("a"), Insert("aa"), Delete("aa"), stack 1. keyMPT: [6 1], keySearch: [6 1 6 1 16], matchLen: 2
				// Del-1. Insert("a"), Insert("b"), Delete("b"), stack 1.
				// Del-4. stack 1.
				// Del-2. stack 1.
				branchNode := db[node.flag_value.value]
				// 'node' is now Branch node next to the Ext node.

				retNode, ret := delete_helper(branchNode, keyMPT[matchLen:], keySearch[matchLen:], db)
				// 'retNode' could be Leaf, Ext, or Branch.
				if retNode.node_type == 1 {
					// 'retNode' is Branch.
					delete(db, node.hash_node())
					node.flag_value.value = putNodeInDb(retNode, db)
					return node, ret

				} else if firstDigit := getFirstDigitOfAscii(retNode.flag_value.encoded_prefix); firstDigit == 3 || firstDigit == 4 || firstDigit == 5 {
					// 'retNode' is Leaf.
					retNode.flag_value.encoded_prefix = compact_encode(
						append(keyMPT,
							append(compact_decode(retNode.flag_value.encoded_prefix), 16)...))

				} else {
					// 'retNode' is Ext.
					// Del-2. Del-7.
					retNode.flag_value.encoded_prefix = compact_encode(
						append(keyMPT, compact_decode(retNode.flag_value.encoded_prefix)...))
				}
				delete(db, node.hash_node())
				node.flag_value.value = putNodeInDb(retNode, db)
				return retNode, ret
			}

			// 'node' is Leaf node.

			if keySearch[matchLen] == 16 && len(keyMPT) == matchLen {
				// Exact match.
				// Del-0. Just one node.
				// Del-3. Insert("a"), Insert("aa"), Delete("aa"), stack 3. keyMPT: [1], keySearch: [1 16], matchLen: 1
				delete(db, node.hash_node())
				nullNode := createNewLeafOrExtNode(0, nil, "")
				return nullNode, ""

				// flagValue := Flag_value{encoded_prefix: nil, value: ""}
				// node = Node{node_type: 0, branch_value: [17]string{}, flag_value: flagValue}
				// return node, ""

			}

			// 'node' is Leaf node and keySearch is shorter or longer than keyMPT.
			return node, "path_not_found"

		} else if matchLen == 0 {

			if keySearch[matchLen] == 16 {
				// Del-1. Insert("a"), Insert("b"), Delete("b"), stack 3. keyMPT: [], keySearch: [16], matchLen: 0
				delete(db, node.hash_node())
				nullNode := createNewLeafOrExtNode(0, nil, "")
				// flagValue := Flag_value{encoded_prefix: nil, value: ""}
				// node = Node{node_type: 0, branch_value: [17]string{}, flag_value: flagValue}
				// return node, ""
				return nullNode, ""
			}

			return node, "path_not_found"
		}

	}

	return node, "path_not_found"
}

func getOnlyOneValueInBranch(node Node) (bool, string, uint8) {
	count := 0
	var index uint8 = 0
	oneValue := ""
	for i, str := range node.branch_value {
		if str != "" {
			index = uint8(i)
			oneValue = str
			count++
		}
	}
	return count <= 1, oneValue, index
}

func createNewLeafOrExtNode(nodeType int, keyHex []uint8, newValue string) Node {
	encodedPrefix := compact_encode(keyHex)
	flagValue := Flag_value{encoded_prefix: encodedPrefix, value: newValue}
	node := Node{node_type: nodeType, flag_value: flagValue}
	return node
}

func putNodeInDb(node Node, db map[string]Node) string {
	hash := node.hash_node()
	db[hash] = node
	return hash
}

func prefixLen(a []uint8, b []uint8) int {
	length := len(a)
	if len(b) < length {
		length = len(b)
	}
	i := 0
	for i < length {
		if a[i] != b[i] {
			break
		}
		i++
	}
	return i
}

func getFirstDigitOfAscii(encodedPrefix []uint8) uint8 {
	firstDigit := encodedPrefix[0] / 10
	return firstDigit
}

func test_compact_encode() {
	fmt.Println(reflect.DeepEqual(compact_decode(compact_encode([]uint8{1, 2, 3, 4, 5})), []uint8{1, 2, 3, 4, 5}))
	fmt.Println(reflect.DeepEqual(compact_decode(compact_encode([]uint8{0, 1, 2, 3, 4, 5})), []uint8{0, 1, 2, 3, 4, 5}))
	fmt.Println(reflect.DeepEqual(compact_decode(compact_encode([]uint8{0, 15, 1, 12, 11, 8, 16})), []uint8{0, 15, 1, 12, 11, 8}))
	fmt.Println(reflect.DeepEqual(compact_decode(compact_encode([]uint8{15, 1, 12, 11, 8, 16})), []uint8{15, 1, 12, 11, 8}))
}

func (node *Node) hash_node() string {
	var str string
	switch node.node_type {
	case 0:
		str = ""
	case 1:
		str = "branch_"
		for _, v := range node.branch_value {
			str += v
		}
	case 2:
		str = node.flag_value.value
	}

	sum := sha3.Sum256([]byte(str))
	return "HashStart_" + hex.EncodeToString(sum[:]) + "_HashEnd"
}

// Additional given code.
func (node *Node) String() string {
	str := "empty string"
	switch node.node_type {
	case 0:
		str = "[Null Node]"
	case 1:
		str = "Branch["
		for i, v := range node.branch_value[:16] {
			str += fmt.Sprintf("%d=\"%s\", ", i, v)
		}
		str += fmt.Sprintf("value=%s]", node.branch_value[16])
	case 2:
		encoded_prefix := node.flag_value.encoded_prefix
		node_name := "Leaf"
		if is_ext_node(encoded_prefix) {
			node_name = "Ext"
		}
		ori_prefix := strings.Replace(fmt.Sprint(compact_decode(encoded_prefix)), " ", ", ", -1)
		str = fmt.Sprintf("%s<%v, value=\"%s\">", node_name, ori_prefix, node.flag_value.value)
	}
	return str
}

func node_to_string(node Node) string {
	return node.String()
}

func (mpt *MerklePatriciaTrie) Initial() {
	mpt.db = make(map[string]Node)
}

func is_ext_node(encoded_arr []uint8) bool {
	return encoded_arr[0]/16 < 2
}

func TestCompact() {
	test_compact_encode()
}

func (mpt *MerklePatriciaTrie) String() string {
	content := fmt.Sprintf("ROOT=%s\n", mpt.root)
	for hash := range mpt.db {
		content += fmt.Sprintf("%s: %s\n", hash, node_to_string(mpt.db[hash]))
	}
	return content
}

func (mpt *MerklePatriciaTrie) Order_nodes() string {
	raw_content := mpt.String()
	content := strings.Split(raw_content, "\n")
	root_hash := strings.Split(strings.Split(content[0], "HashStart")[1], "HashEnd")[0]
	queue := []string{root_hash}
	i := -1
	rs := ""
	cur_hash := ""
	for len(queue) != 0 {
		last_index := len(queue) - 1
		cur_hash, queue = queue[last_index], queue[:last_index]
		i += 1
		line := ""
		for _, each := range content {
			if strings.HasPrefix(each, "HashStart"+cur_hash+"HashEnd") {
				line = strings.Split(each, "HashEnd: ")[1]
				rs += each + "\n"
				rs = strings.Replace(rs, "HashStart"+cur_hash+"HashEnd", fmt.Sprintf("Hash%v", i), -1)
			}
		}
		temp2 := strings.Split(line, "HashStart")
		flag := true
		for _, each := range temp2 {
			if flag {
				flag = false
				continue
			}
			queue = append(queue, strings.Split(each, "HashEnd")[0])
		}
	}
	return rs
}
