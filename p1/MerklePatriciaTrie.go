package p1

import (
	"encoding/hex"
	"fmt"
	"reflect"

	"golang.org/x/crypto/sha3"
)

type Flag_value struct {
	encoded_prefix []uint8
	value          string
}

type Node struct {
	node_type    int // 0: Null, 1: Branch, 2: Ext or Leaf
	branch_value [17]string
	flag_value   Flag_value
}

type MerklePatriciaTrie struct {
	// K: Node's hash value, V: Node
	db map[string]Node
	// hash value of the root node
	root string
}

func NewMPT() *MerklePatriciaTrie {
	node := Node{}
	hash := node.hash_node()

	mpt := &MerklePatriciaTrie{}
	mpt.db = map[string]Node{hash: node}
	mpt.root = hash

	return mpt
}

func (mpt *MerklePatriciaTrie) Get(key string) string {
	// TODO
	return ""
}

func (mpt *MerklePatriciaTrie) Insert(key string, new_value string) {
	// TODO
	fmt.Println("hello")

}

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

func (mpt *MerklePatriciaTrie) Delete(key string) {
	// TODO
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
	var firstHexValue uint8 = 2*isLeaf + isOdd
	if isOdd == 1 {
		hex_array = append([]uint8{firstHexValue}, hex_array...)
	} else {
		hex_array = append(append([]uint8{firstHexValue}, 0), hex_array...)
	}
	// 'hex_array' now has an even length whose first nibble is the 'firstHexValue'.
	var encoded_prefix []uint8
	for i := 0; i < len(hex_array); i = i + 2 {
		encoded_prefix = append(encoded_prefix, 16*hex_array[i]+hex_array[i+1])
	}

	return encoded_prefix
}

// If Leaf, ignore 16 at the end
func compact_decode(encoded_arr []uint8) []uint8 {
	// TODO
	return []uint8{}
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

// Print the output of compact_encode()
func Test_compact_encode(hex_array []uint8) {
	ret := compact_encode(hex_array)
	fmt.Println(ret)
}

// Print the output of convert_string_to_hex()
func Test_convert_string_to_hex(key string) {
	ret := convert_string_to_hex(key)
	fmt.Println(ret)
}
