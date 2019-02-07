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

	hash := node.hash_node()

	mpt := &MerklePatriciaTrie{}
	mpt.db = map[string]Node{hash: node}
	mpt.root = hash

	return mpt
}

// Convert key string to hex value array and append 16.
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
	keySearch := convert_string_to_hex(key)

	rootNode := mpt.db[mpt.root]
	encodedPrefix := rootNode.flag_value.encoded_prefix
	keyMPT := compact_decode(encodedPrefix)

	return get_helper(rootNode, keyMPT, keySearch, mpt.db)
}
func get_helper(node Node, keyMPT, keySearch []uint8, db map[string]Node) string {
	// if keySearch[0] == 16 {
	// 	return node.flag_value.value
	// }
	nodeType := node.node_type
	switch nodeType {
	case 0:
		// Null node
		return ""
	case 1:
		// Branch

	case 2:
		// Ext or Leaf

		matchLen := prefixLen(keySearch, keyMPT)

		// if matchLen == len(keyMPT) {
		if matchLen != 0 {

			if keySearch[matchLen] == 16 {
				// Case A. keyMPT: [6 1], keySearch: [6 1 16]
				if node, ok := db[node.flag_value.value]; ok {
					return node.branch_value[16]
				}
				return node.flag_value.value
			}

			// Case B. keyMPT: [6 1], keySearch: [6 1 6 1 16]
			// extNode := createNewLeafOrExtNode(2, keySearch[:matchLen], node.flag_value.value)
			// branchNode := insert_helper(extNode, keyMPT[matchLen:], keySearch[matchLen:], new_value)
			// extNode.flag_value.value = branchNode.hash_node()
			// return extNode

		}

		// Ext
		// nextNode := db[node.flag_value.value]
		// return get_helper(nextNode, keySearch[len(decodedPrefix):], db)

	}

	return ""
}

func (mpt *MerklePatriciaTrie) Insert(key string, new_value string) {
	// TODO
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
	// if keyMPT == nil && keySearch == nil {
	// 	return node
	// }
	// if keySearch[0] == 16 {
	// 	node.flag_value.value = new_value
	// 	updateMPT(mpt, node)
	// 	return
	// }

	nodeType := node.node_type
	switch nodeType {
	case 0:
		// Null node
		// Create a new Leaf node.
		node = createNewLeafOrExtNode(2, keySearch, new_value)
		return node
	case 1:
		// Branch

	case 2:
		// Ext or Leaf

		matchLen := prefixLen(keySearch, keyMPT)
		// stack 1: 2
		// stack 2: 0

		// if matchLen == len(keyMPT) {
		if matchLen != 0 {
			if keySearch[matchLen] == 16 {
				// Case A. keyMPT: [6 1], keySearch: [6 1 16]
				node.flag_value.value = new_value
				return node
			}

			// Case B.
			// stack 1. keyMPT: [6 1], keySearch: [6 1 6 1 16]
			extNode := createNewLeafOrExtNode(2, keyMPT[:matchLen], "")
			branchNode := insert_helper(extNode, keyMPT[matchLen:], keySearch[matchLen:], new_value, db)
			branchNode.branch_value[16] = node.flag_value.value
			hash := branchNode.hash_node()
			db[hash] = branchNode
			extNode.flag_value.value = branchNode.hash_node()
			return extNode

		} else if matchLen == 0 {
			// Case B.
			// stack 2. keyMPT: [], keySearch: [6 1 16]
			branchNode := Node{node_type: 1, branch_value: [17]string{}}
			leafNode := createNewLeafOrExtNode(2, keySearch[matchLen+1:], new_value)
			hash := leafNode.hash_node()
			db[hash] = leafNode
			// leafNode = insert_helper(leafNode, nil, nil, new_value)
			branchNode.branch_value[keySearch[matchLen]] = hash
			return branchNode
		}

		return node

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
