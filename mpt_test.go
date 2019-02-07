package main

import (
	"cs686/cs686-project-1/p1"
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	key := "do"
	hex_array := p1.Test_convert_string_to_hex(key) // [6 4 6 15 16]
	p1.Test_compact_encode(hex_array)               // [32 100 111]

	hex_array = []uint8{1, 6, 1}
	p1.Test_compact_encode(hex_array) // [17 97]

	hex_array = []uint8{2}
	p1.Test_compact_encode(hex_array) // [18]

	hex_array2 := []uint8{2, 16}
	p1.Test_compact_encode(hex_array2) // [50]

	hex_array2 = []uint8{6, 1}
	p1.Test_compact_encode(hex_array2) // [0 97]

	key = "a"
	hex_array = p1.Test_convert_string_to_hex(key) // [6 1 16]
	p1.Test_compact_encode(hex_array)              // [32 97]

	key = "ab"
	hex_array = p1.Test_convert_string_to_hex(key) // [6 1 6 2 16]
	p1.Test_compact_encode(hex_array)              // [32 97 98]

	p1.Test_test_compact_encode()

	// Test Insert(), Get()
	fmt.Println("Test Insert(), Get(): ")
	mpt := p1.NewMPT()
	var ret string

	// // Insert("a", "apple"), Get("a")
	// fmt.Println("Insert(\"a\", \"apple\")")
	// mpt.Insert("a", "apple")
	// fmt.Println("Get(\"a\")")
	// ret = mpt.Get("a")
	// fmt.Println(ret) // apple
	// if ret != "apple" {
	// 	t.Errorf("Expected %s, but was %s", "apple", ret)
	// }

	// // Case A:
	// // Insert("a", "apple")
	// // Insert("a", "orange")
	// // Get("a")
	// fmt.Println("Case A: ")
	// fmt.Println("Insert(\"a\", \"orange\")")
	// mpt.Insert("a", "orange")
	// fmt.Println("Get(\"a\")")
	// ret = mpt.Get("a")
	// fmt.Println(ret) // orange
	// if ret != "orange" {
	// 	t.Errorf("Expected %s, but was %s", "orange", ret)
	// }

	// // Test prefixLen()
	// a := []uint8{2, 2, 3, 4}
	// b := []uint8{2, 2, 3, 5}
	// p1.Test_prefixLen(a, b) // 3
	// b = []uint8{0, 2, 3, 5}
	// p1.Test_prefixLen(a, b) // 0
	// b = []uint8{2, 2, 3, 4}
	// p1.Test_prefixLen(a, b) // 4

	// Case B:
	// Insert("a", "apple")
	// Insert("aa", "orange")
	// Get("a")
	fmt.Println("Case B: ")
	mpt = p1.NewMPT()
	fmt.Println("Insert(\"a\", \"apple\")")
	mpt.Insert("a", "apple")
	fmt.Println("Insert(\"aa\", \"orange\")")
	mpt.Insert("aa", "orange")
	fmt.Println("Get(\"a\")")
	ret = mpt.Get("a")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("Expected %s, but was %s", "apple", ret)
	}
	// // Insert("a", "apple")
	// // Insert("aa", "orange")
	// // Get("aa")
	// fmt.Println("Get(\"aa\")")
	// ret = mpt.Get("aa")
	// fmt.Println(ret) // ??
	// if ret != "orange" {
	// 	t.Errorf("Expected %s, but was %s", "orange", ret)
	// }

}
