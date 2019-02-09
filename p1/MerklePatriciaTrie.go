package p1

import (
	"encoding/hex"
	"fmt"
	"reflect"

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
	flagValue := Flag_value{encoded_prefix: nil, value: ""}
	node := Node{node_type: 0, branch_value: [17]string{}, flag_value: flagValue}

	mpt := &MerklePatriciaTrie{}
	mpt.db = map[string]Node{}
	mpt.root = putNodeInDb(node, mpt.db)

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
	// fmt.Println("isLeaf: ", isLeaf)
	// fmt.Println("isOdd: ", isOdd)
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

// If Leaf, ignore 16 at the end << why?
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

	// Append 16 if it is Leaf.
	// if hex_array[0] >= 2 {
	// 	hex_array = append(hex_array, 16)
	// }

	cut := 2 - hex_array[0]&1
	return hex_array[cut:]
}

func (mpt *MerklePatriciaTrie) Get(key string) string {
	// TODO
	if key == "" {
		return ""
	}

	keySearch := convert_string_to_hex(key)

	rootNode := mpt.db[mpt.root]
	encodedPrefix := rootNode.flag_value.encoded_prefix
	keyMPT := compact_decode(encodedPrefix)

	return get_helper(rootNode, keyMPT, keySearch, mpt.db)
}
func get_helper(node Node, keyMPT, keySearch []uint8, db map[string]Node) string {
	nodeType := node.node_type
	switch nodeType {
	case 0:
		// Null node
		return ""
	case 1:
		// Branch

		if node, ok := db[node.branch_value[keySearch[0]]]; ok {
			// 'node' is now the next node.
			// Case B-1. Insert("a"), Insert("aa"), Get("aa"), stack 2. keyMPT: [], keySearch: [6 1 16]
			// Case B-2. Insert("a"), Insert("b"), Get("a"), stack 2. keyMPT: [], keySearch: [1 16]
			// Case B-3. Insert("aa"), Insert("a"), Get("aa"), stack 2. keyMPT: [], keySearch: [6 1 16]
			// Case C. Insert("a"), Insert("p"), Get("a"), stack 1. keyMPT: [], keySearch: [6 1 16]
			// Case D-1. Insert("a"), Insert("p"), Insert("abc"), Get("abc") stack 1. keyMPT: [], keySearch: [6 1 6 2 6 3 16]
			// Case D-1. Insert("a"), Insert("p"), Insert("abc"), Get("abc") stack 3. keyMPT: [], keySearch: [6 2 6 3 16]
			// Case D-3. Insert("a"), Insert("p"), Insert("A"), Get("A") stack 1. keyMPT: [], keySearch: [4 1 16]
			encodedPrefix := node.flag_value.encoded_prefix
			keyMPT := compact_decode(encodedPrefix)
			return get_helper(node, keyMPT, keySearch[1:], db)
		}
		// If there is no link in the Branch.
		return "Not found"

	case 2:
		// Ext or Leaf

		matchLen := prefixLen(keySearch, keyMPT)

		if matchLen != 0 {

			if keySearch[matchLen] == 16 {

				if firstDigit := getFirstDigitOfAscii(node.flag_value.encoded_prefix); firstDigit == 0 || firstDigit == 1 || firstDigit == 2 {
					// 'node' is Ext node.
					// Case B-1. Insert("a"), Insert("aa"), Get("a"), stack 1. keyMPT: [6 1], keySearch: [6 1 16], matchLen: 2

					node = db[node.flag_value.value]
					// 'node' is now Branch node next to the Ext node.
					return node.branch_value[16]
				}
				// 'node' is Leaf node.
				// Case A. Insert("a"), Get("a"). keyMPT: [6 1], keySearch: [6 1 16], matchLen: 2
				// Case B-1. Insert("a"), Insert("aa"), Get("aa"), stack 3. keyMPT: [1], keySearch: [1 16], matchLen: 1
				// Case B-3. Insert("aa"), Insert("a"), Get("aa"), stack 3. keyMPT: [1], keySearch: [1 16], matchLen: 1
				// Case C.   Insert("a"), Insert("p"), Get("a"), stack 2. keyMPT: [1], keySearch: [1 16], matchLen: 1
				// Case D-1. Insert("a"), Insert("p"), Insert("abc"), Get("abc") stack 4. keyMPT: [2 6 3], keySearch: [2 6 3 16], matchLen: 3
				// Case D-3. Insert("a"), Insert("p"), Insert("A"), Get("A") stack 2. keyMPT: [1], keySearch: [1 16]
				return node.flag_value.value
			}

			if firstDigit := getFirstDigitOfAscii(node.flag_value.encoded_prefix); firstDigit == 0 || firstDigit == 1 || firstDigit == 2 {
				// 'node' is Ext node.
				// Case B-1. Insert("a"), Insert("aa"), Get("aa"), stack 1. keyMPT: [6 1], keySearch: [6 1 6 1 16], matchLen: 2
				// Case B-2. Insert("a"), Insert("b"), Get("a"), stack 1. keyMPT: [6], keySearch: [6 1 16], matchLen: 1
				// Case B-3. Insert("aa"), Insert("a"), Get("aa"), stack 1. keyMPT: [6 1], keySearch: [6 1 6 1 16], matchLen: 2
				// Case D-1. Insert("a"), Insert("p"), Insert("abc"), Get("abc") stack 2. keyMPT: [1], keySearch: [1 6 2 6 3 16], matchLen: 1

				node = db[node.flag_value.value]
				// 'node' is now Branch node next to the Ext node.
				return get_helper(node, keyMPT[matchLen:], keySearch[matchLen:], db)
			}
			// 'node' is Leaf node and keySearch still has hex number.
			return "Not found."

		} else if matchLen == 0 {

			if keySearch[matchLen] == 16 {
				// Case B-2. Insert("a"), Insert("b"), Get("a"), stack 3. keyMPT: [], keySearch: [16], matchLen: 0
				return node.flag_value.value
			}
			return "Not found."
		}

	}

	return "Not found."
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

	rootNode = insert_helper(rootNode, keyMPT, keySearch, new_value, db)
	updateMPT(mpt, rootNode)

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
			// 'nextNode' is the node next to the Branch.
			// Case D-1. Insert("a"), Insert("p"), Insert("abc"), Get("abc"),
			// stack 1. keyMPT: [], keySearch: [6 1 6 2 6 3 16]
			// Case E-1. Insert("p"), Insert("aaaaa"), Insert("aaaap"), Insert("aa"), Get("aa"),
			// stack 1. keyMPT: [], keySearch: [6 1 6 1 16]
			// Case E-2. Insert("p"), Insert("aaaaa"), Insert("aaaap"), Insert("aaaa"), Get("aaaa"),
			// stack 1. keyMPT: [], keySearch: [6 1 6 1 6 1 6 1 16]
			encodedPrefix := nextNode.flag_value.encoded_prefix
			keyMPT := compact_decode(encodedPrefix)
			nextNode = insert_helper(nextNode, keyMPT, keySearch[1:], new_value, db)
			node.branch_value[keySearch[0]] = putNodeInDb(nextNode, db)
			return node
		}
		// There is no link in the Branch.
		// Case D-3. Insert("a"), Insert("p"), Insert("A"), Get("A"). keyMPT: [], keySearch: [4 1 16]
		leafNode := createNewLeafOrExtNode(2, keySearch[1:], new_value)
		node.branch_value[keySearch[0]] = putNodeInDb(leafNode, db)
		return node

	case 2:
		// Ext or Leaf

		matchLen := prefixLen(keySearch, keyMPT)

		if matchLen != 0 {

			if firstDigit := getFirstDigitOfAscii(node.flag_value.encoded_prefix); firstDigit == 0 || firstDigit == 1 || firstDigit == 2 {
				// 'node' is Ext.

				if keySearch[matchLen] == 16 && len(keySearch) < len(keyMPT) {
					// Case E-1. stack 2. keyMPT: [1 6 1 6 1 6 1], keySearch: [1 6 1 16], matchLen: 3
					extNode1 := createNewLeafOrExtNode(2, keyMPT[:matchLen], "")
					branchNode := Node{node_type: 1, branch_value: [17]string{}}
					extNode2 := createNewLeafOrExtNode(2, keyMPT[:matchLen], node.flag_value.value)
					branchNode.branch_value[keyMPT[matchLen]] = putNodeInDb(extNode2, db)
					branchNode.branch_value[16] = new_value
					extNode1.flag_value.value = putNodeInDb(branchNode, db)

					return extNode1
				}

				if keySearch[matchLen] == 16 && len(keyMPT) == matchLen {
					// Case E-2. stack 2. keyMPT: [1 6 1 6 1 6 1], keySearch: [1 6 1 6 1 6 1 16], matchLen: 7
					branchNode := db[node.flag_value.value]
					branchNode = insert_helper(branchNode, nil, keySearch[matchLen:], new_value, db)
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
			if matchLen == len(keyMPT) {
				// Case B-1.
				// Case D-1.
				branchNode.branch_value[16] = node.flag_value.value
			} else if keySearch[matchLen] == 16 {
				// Case B-3.
				branchNode.branch_value[16] = new_value
			}

			extNode.flag_value.value = putNodeInDb(branchNode, db)
			return extNode

		} else if matchLen == 0 {

			branchNode := Node{node_type: 1, branch_value: [17]string{}}
			if len(keyMPT) == 0 {
				// Case B-1 (Prefix match).
				// stack 2. keyMPT: [], keySearch: [6 1 16], matchLen: 0
				// Case D-1. stack 3. keyMPT: [], keySearch: [6 2 6 3 16], matchLen: 0
				leafNode := createNewLeafOrExtNode(2, keySearch[matchLen+1:], new_value)
				branchNode.branch_value[keySearch[matchLen]] = putNodeInDb(leafNode, db)

			} else if keySearch[matchLen] == 16 {
				// Case B-3 (Prefix match).
				// stack 2. keyMPT: [6 1], keySearch: [16], matchLen: 0
				leafNode := createNewLeafOrExtNode(2, append(keyMPT[matchLen+1:], 16), node.flag_value.value)
				branchNode.branch_value[keyMPT[matchLen]] = putNodeInDb(leafNode, db)

			} else {
				// Case B-2 (Prefix match).
				// stack 2. keyMPT: [1], keySearch: [2 16], matchLen: 0
				// Case C (Mismatch).
				// keyMPT: [6 1], keySearch: [7 0 16], matchLen: 0
				leafNode := createNewLeafOrExtNode(2, keySearch[matchLen+1:], new_value)
				branchNode.branch_value[keySearch[matchLen]] = putNodeInDb(leafNode, db)

				leafNode = createNewLeafOrExtNode(2, append(keyMPT[matchLen+1:], 16), node.flag_value.value)
				branchNode.branch_value[keyMPT[matchLen]] = putNodeInDb(leafNode, db)
			}

			return branchNode
		}
	default:

	}

	return node
}

func (mpt *MerklePatriciaTrie) Delete(key string) {
	// TODO
}

func updateMPT(mpt *MerklePatriciaTrie, node Node) {
	hash := node.hash_node()
	mpt.db[hash] = node
	mpt.root = hash
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
