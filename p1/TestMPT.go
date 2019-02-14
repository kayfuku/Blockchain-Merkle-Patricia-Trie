/*
	For testing MerklePatriciaTrie.go
	Author: Kei Fukutani
	Date  : February 13, 2019
*/
package p1

import (
	"fmt"
)

// Test convert_string_to_hex()
func Test_convert_string_to_hex(key string) []uint8 {
	ret := convert_string_to_hex(key)
	fmt.Println(ret)
	return ret
}

// Test compact_encode()
func Test_compact_encode(hex_array []uint8) []uint8 {
	ret := compact_encode(hex_array)
	fmt.Println(ret)
	return ret
}

// Test test_compact_encode()
func Test_test_compact_encode() {
	fmt.Println("Test_test_compact_encode() start.")
	test_compact_encode()
	return
}

// Test prefixLen()
func Test_prefixLen(a []uint8, b []uint8) int {
	ret := prefixLen(a, b)
	fmt.Println(ret)
	return ret
}

// Test getFirstDigitOfAscii()
func Test_getFirstDigitOfAscii(a []uint8) uint8 {
	ret := getFirstDigitOfAscii(a)
	fmt.Println(ret)
	return ret
}

// Test getOnlyOneValueInBranch()
func Test_getOnlyOneValueInBranch(array [17]string) bool {
	node := Node{node_type: 1, branch_value: array}
	ret, str, index := getOnlyOneValueInBranch(node)
	fmt.Printf("ret: %v str: %s index: %d\n", ret, str, index)
	return ret
}
