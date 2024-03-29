/*
	For testing MerklePatriciaTrie.go
	Author: Kei Fukutani
	Date  : February 13, 2019
*/
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

	// Ext, even
	hex_array2 := []uint8{6, 1}
	p1.Test_compact_encode(hex_array2) // [0 97]
	hex_array2 = []uint8{7, 0}
	p1.Test_compact_encode(hex_array2) // [0 112]
	hex_array2 = []uint8{3, 0}
	p1.Test_compact_encode(hex_array2) // [0 48]
	hex_array2 = []uint8{7, 10}
	p1.Test_compact_encode(hex_array2) // [0 122]

	// Ext, odd
	hex_array = []uint8{2}
	p1.Test_compact_encode(hex_array) // [18]
	hex_array = []uint8{1, 6, 1}
	p1.Test_compact_encode(hex_array) // [17 97]

	// Leaf, even
	hex_array2 = []uint8{6, 4, 16}
	p1.Test_compact_encode(hex_array2) // [32 100]
	hex_array2 = []uint8{6, 4, 7, 0, 16}
	p1.Test_compact_encode(hex_array2) // [32 100 112]
	hex_array2 = []uint8{3, 0, 16}
	p1.Test_compact_encode(hex_array2) // [32 48]
	hex_array2 = []uint8{7, 10, 16}
	p1.Test_compact_encode(hex_array2) // [32 122]

	// Leaf, odd
	hex_array2 = []uint8{2, 16}
	p1.Test_compact_encode(hex_array2) // [50]
	hex_array2 = []uint8{2, 5, 4, 16}
	p1.Test_compact_encode(hex_array2) // [50 84]

	key = "a"
	hex_array = p1.Test_convert_string_to_hex(key) // [6 1 16]
	p1.Test_compact_encode(hex_array)              // [32 97]

	key = "ab"
	hex_array = p1.Test_convert_string_to_hex(key) // [6 1 6 2 16]
	p1.Test_compact_encode(hex_array)              // [32 97 98]

	key = "0"
	hex_array = p1.Test_convert_string_to_hex(key) // [3 0 16]
	p1.Test_compact_encode(hex_array)              // [32 48]

	key = "A"
	hex_array = p1.Test_convert_string_to_hex(key) // [4 1 16]
	p1.Test_compact_encode(hex_array)              // [32 65]

	key = "z"
	hex_array = p1.Test_convert_string_to_hex(key) // [7 10 16]
	p1.Test_compact_encode(hex_array)              // [32 122]

	key = ""
	hex_array = p1.Test_convert_string_to_hex(key) // [16]
	p1.Test_compact_encode(hex_array)              // [32]

	// Test getFirstDigitOfAscii()
	ascii := []uint8{0, 97}
	p1.Test_getFirstDigitOfAscii(ascii) // 0
	ascii = []uint8{16, 97}
	p1.Test_getFirstDigitOfAscii(ascii) // 1
	ascii = []uint8{25, 97}
	p1.Test_getFirstDigitOfAscii(ascii) // 2
	ascii = []uint8{32, 97}
	p1.Test_getFirstDigitOfAscii(ascii) // 3
	ascii = []uint8{48, 97}
	p1.Test_getFirstDigitOfAscii(ascii) // 4
	ascii = []uint8{57, 97}
	p1.Test_getFirstDigitOfAscii(ascii) // 5

	// Test test_compact_encode()
	p1.Test_test_compact_encode()

	// Test prefixLen()
	a := []uint8{2, 2, 3, 4}
	b := []uint8{2, 2, 3, 5}
	p1.Test_prefixLen(a, b) // 3
	b = []uint8{0, 2, 3, 5}
	p1.Test_prefixLen(a, b) // 0
	b = []uint8{2, 2, 3, 4}
	p1.Test_prefixLen(a, b) // 4

	// Test getOnlyOneValueInBranch()
	fmt.Println("Test getOnlyOneValueInBranch(): ")
	array := [17]string{}
	p1.Test_getOnlyOneValueInBranch(array) // ret: true str:  index: 0
	array[6] = "***"
	p1.Test_getOnlyOneValueInBranch(array) // ret: true str: *** index: 6
	array[16] = "last"
	p1.Test_getOnlyOneValueInBranch(array) // ret: false str: last index: 16
	array[6] = ""
	p1.Test_getOnlyOneValueInBranch(array) // ret: true str: last index: 16

	// Test Insert(), Get() start.
	fmt.Println("Test Insert(), Get(): ")

	mpt := p1.NewMPT()
	var ret string

	ret = mpt.Get("a")
	fmt.Println(ret) // ""

	// Insert("a", "apple")
	// Get("a"), Delete("a")
	fmt.Println("Insert(\"a\", \"apple\")")
	mpt.Insert("a", "apple")
	fmt.Println("Get(\"a\")")
	ret = mpt.Get("a")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("Expected %s, but was %s", "apple", ret)
	}
	fmt.Println("Delete(\"a\")")
	mpt.Delete("a")
	ret = mpt.Get("a")
	fmt.Println(ret) // ""
	if ret != "" {
		t.Errorf("Expected %s, but was %s", "", ret)
	}

	// Case A (Exact match):
	// Insert("a", "apple")
	// Insert("a", "orange")
	// Get("a"), Delete("a")
	fmt.Println("Case A: ")
	fmt.Println("Insert(\"a\", \"orange\")")
	mpt.Insert("a", "orange")
	fmt.Println("Get(\"a\")")
	ret = mpt.Get("a")
	fmt.Println(ret) // orange
	if ret != "orange" {
		t.Errorf("Case A (Exact match): Insert(\"a\", \"apple\"), Insert(\"a\", \"orange\"), Get(\"a\") \n"+
			"Expected %s, but was %s", "orange", ret)
	}
	fmt.Println("Delete(\"a\")")
	mpt.Delete("a")
	ret = mpt.Get("a")
	fmt.Println(ret) // ""
	if ret != "" {
		t.Errorf("Expected %s, but was %s", "", ret)
	}

	// Case B-1 (Partial match):
	// Insert("a", "apple")
	// Insert("aa", "orange")
	// Get("a")
	fmt.Println("Case B-1: ")
	mpt = p1.NewMPT()
	fmt.Println("Insert(\"a\", \"apple\")")
	mpt.Insert("a", "apple")
	fmt.Println("Insert(\"aa\", \"orange\")")
	mpt.Insert("aa", "orange")
	fmt.Println("Get(\"a\")")
	ret = mpt.Get("a")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("Case B-1 (Partial match): Insert(\"a\", \"apple\"), Insert(\"aa\", \"orange\"), Get(\"a\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}

	// Case B-1 (Partial match):
	// Insert("a", "apple")
	// Insert("aa", "orange")
	// Get("aa")
	fmt.Println("Get(\"aa\")")
	ret = mpt.Get("aa")
	fmt.Println(ret) // orange
	if ret != "orange" {
		t.Errorf("Case B-1 (Partial match): Insert(\"a\", \"apple\"), Insert(\"aa\", \"orange\"), Get(\"aa\") \n"+
			"Expected %s, but was %s", "orange", ret)
	}

	// Del-3.
	fmt.Println("Del-3. Delete(\"aa\")")
	mpt.Delete("aa")
	fmt.Println("Get(\"a\")")
	ret = mpt.Get("a")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("Del-3. Delete(\"aa\"): Insert(\"a\", \"apple\"), Insert(\"aa\", \"orange\"), Get(\"a\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}
	fmt.Println("Get(\"aa\")")
	ret = mpt.Get("aa")
	fmt.Println(ret) // ""
	if ret != "" {
		t.Errorf("Del-3. Delete(\"aa\"): Insert(\"a\", \"apple\"), Insert(\"aa\", \"orange\"), Get(\"aa\") \n"+
			"Expected %s, but was %s", "", ret)
	}

	// Case B-2 (Partial match):
	// Insert("a", "apple")
	// Insert("b", "orange")
	// Get("a"), Get("b")
	fmt.Println("Case B-2: ")
	mpt = p1.NewMPT()
	fmt.Println("Insert(\"a\", \"apple\")")
	mpt.Insert("a", "apple")
	fmt.Println("Insert(\"b\", \"orange\")")
	mpt.Insert("b", "orange")
	fmt.Println("Get(\"a\")")
	ret = mpt.Get("a")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("Case B-2 (Partial match): Insert(\"a\", \"apple\"), Insert(\"b\", \"orange\"), Get(\"a\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}
	fmt.Println("Get(\"b\")")
	ret = mpt.Get("b")
	fmt.Println(ret) // orange
	if ret != "orange" {
		t.Errorf("Case B-2 (Partial match): Insert(\"a\", \"apple\"), Insert(\"b\", \"orange\"), Get(\"b\") \n"+
			"Expected %s, but was %s", "orange", ret)
	}

	// Del-1.
	fmt.Println("Del-1. Delete(\"b\")")
	mpt.Delete("b")
	fmt.Println("Get(\"a\")")
	ret = mpt.Get("a")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("Del-1. Delete(\"b\"): Insert(\"a\", \"apple\"), Insert(\"b\", \"orange\"), Get(\"a\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}
	fmt.Println("Get(\"b\")")
	ret = mpt.Get("b")
	fmt.Println(ret) // ""
	if ret != "" {
		t.Errorf("Del-1. Delete(\"b\"): Insert(\"a\", \"apple\"), Insert(\"b\", \"orange\"), Get(\"b\") \n"+
			"Expected %s, but was %s", "", ret)
	}

	// Case B-3 (Partial match):
	// Insert("aa", "apple")
	// Insert("a", "orange")
	// Get("a")
	fmt.Println("Case B-3: ")
	mpt = p1.NewMPT()
	fmt.Println("Insert(\"aa\", \"apple\")")
	mpt.Insert("aa", "apple")
	fmt.Println("Insert(\"a\", \"orange\")")
	mpt.Insert("a", "orange")
	fmt.Println("Get(\"a\")")
	ret = mpt.Get("a")
	fmt.Println(ret) // orange
	if ret != "orange" {
		t.Errorf("Case B-3 (Partial match): Insert(\"aa\", \"apple\"), Insert(\"a\", \"orange\"), Get(\"a\") \n"+
			"Expected %s, but was %s", "orange", ret)
	}

	// Case B-3 (Partial match):
	// Insert("aa", "apple")
	// Insert("a", "orange")
	// Get("aa")
	fmt.Println("Get(\"aa\")")
	ret = mpt.Get("aa")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("Case B-3 (Partial match): Insert(\"aa\", \"apple\"), Insert(\"a\", \"orange\"), Get(\"aa\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}

	// Del-4.
	fmt.Println("Del-4. Delete(\"a\")")
	mpt.Delete("a")
	fmt.Println("Get(\"a\")")
	ret = mpt.Get("a")
	fmt.Println(ret) // ""
	if ret != "" {
		t.Errorf("Del-4. Delete(\"a\"): Insert(\"aa\", \"apple\"), Insert(\"a\", \"orange\"), Get(\"a\") \n"+
			"Expected %s, but was %s", "", ret)
	}
	fmt.Println("Get(\"aa\")")
	ret = mpt.Get("aa")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("Del-4. Delete(\"a\"): Insert(\"aa\", \"apple\"), Insert(\"a\", \"orange\"), Get(\"aa\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}

	// Case C (Mismatch).
	// Insert("a", "apple")
	// Insert("p", "orange")
	// Get("a"), Get("p")
	fmt.Println("Case C: ")
	mpt = p1.NewMPT()
	fmt.Println("Insert(\"a\", \"apple\")")
	mpt.Insert("a", "apple")
	fmt.Println("Insert(\"p\", \"orange\")")
	mpt.Insert("p", "orange")
	fmt.Println("Get(\"a\")")
	ret = mpt.Get("a")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("Case C (Mismatch): Insert(\"a\", \"apple\"), Insert(\"p\", \"orange\"), Get(\"a\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}
	fmt.Println("Get(\"p\")")
	ret = mpt.Get("p")
	fmt.Println(ret) // orange
	if ret != "orange" {
		t.Errorf("Case C (Mismatch): Insert(\"a\", \"apple\"), Insert(\"p\", \"orange\"), Get(\"p\") \n"+
			"Expected %s, but was %s", "orange", ret)
	}

	// Del-1. C.
	fmt.Println("Del-1. C. Delete(\"p\")")
	mpt.Delete("p")
	fmt.Println("Get(\"a\")")
	ret = mpt.Get("a")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("Del-1. C. Insert(\"a\", \"apple\"), Insert(\"p\", \"orange\"), Delete(\"p\"), Get(\"a\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}
	fmt.Println("Get(\"p\")")
	ret = mpt.Get("p")
	fmt.Println(ret) // ""
	if ret != "" {
		t.Errorf("Del-1. C. Insert(\"a\", \"apple\"), Insert(\"p\", \"orange\"), Delete(\"p\"), Get(\"p\") \n"+
			"Expected %s, but was %s", "", ret)
	}

	// Case C-2.
	// Insert("a", "apple")
	// Insert("b", "orange")
	// Insert("a", "banana")
	// Get("a"), Get("b")
	fmt.Println("Case C-2: ")
	mpt = p1.NewMPT()
	fmt.Println("Insert(\"a\", \"apple\")")
	mpt.Insert("a", "apple")
	fmt.Println("Insert(\"b\", \"orange\")")
	mpt.Insert("b", "orange")
	fmt.Println("Insert(\"a\", \"banana\")")
	mpt.Insert("a", "banana")
	fmt.Println("Get(\"a\")")
	ret = mpt.Get("a")
	fmt.Println(ret) // banana
	if ret != "banana" {
		t.Errorf("Case C-2: Insert(\"a\", \"apple\"), Insert(\"b\", \"orange\"), Insert(\"a\", \"banana\"), Get(\"a\") \n"+
			"Expected %s, but was %s", "banana", ret)
	}
	fmt.Println("Get(\"b\")")
	ret = mpt.Get("b")
	fmt.Println(ret) // orange
	if ret != "orange" {
		t.Errorf("Case C-2: Insert(\"a\", \"apple\"), Insert(\"b\", \"orange\"), Insert(\"a\", \"banana\"), Get(\"b\") \n"+
			"Expected %s, but was %s", "orange", ret)
	}

	// Case C-3.
	// Insert("a", "apple")
	// Insert("b", "orange")
	// Insert("ab", "banana")
	// Get("a"), Get("b"), Get("ab")
	fmt.Println("Case C-3: ")
	mpt = p1.NewMPT()
	fmt.Println("Insert(\"a\", \"apple\")")
	mpt.Insert("a", "apple")
	fmt.Println("Insert(\"b\", \"orange\")")
	mpt.Insert("b", "orange")
	fmt.Println("Insert(\"ab\", \"banana\")")
	mpt.Insert("ab", "banana")
	fmt.Println("Get(\"a\")")
	ret = mpt.Get("a")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("Case C-3: Insert(\"a\", \"apple\"), Insert(\"b\", \"orange\"), Insert(\"ab\", \"banana\"), Get(\"a\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}
	fmt.Println("Get(\"b\")")
	ret = mpt.Get("b")
	fmt.Println(ret) // orange
	if ret != "orange" {
		t.Errorf("Case C-3: Insert(\"a\", \"apple\"), Insert(\"b\", \"orange\"), Insert(\"ab\", \"banana\"), Get(\"b\") \n"+
			"Expected %s, but was %s", "orange", ret)
	}
	fmt.Println("Get(\"ab\")")
	ret = mpt.Get("ab")
	fmt.Println(ret) // banana
	if ret != "banana" {
		t.Errorf("Case C-3: Insert(\"a\", \"apple\"), Insert(\"b\", \"orange\"), Insert(\"ab\", \"banana\"), Get(\"ab\") \n"+
			"Expected %s, but was %s", "banana", ret)
	}

	// Instructor's test cases.
	// Insert("a", "apple")
	// Insert("b", "banana")
	// Get("a")
	fmt.Println("Instructor's test cases: ")
	mpt = p1.NewMPT()
	fmt.Println("Insert(\"a\", \"apple\")")
	mpt.Insert("a", "apple")
	fmt.Println("Insert(\"b\", \"banana\")")
	mpt.Insert("b", "banana")
	fmt.Println("Get(\"a\")")
	ret = mpt.Get("a")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("Instructor's test case. Insert(\"a\", \"apple\"), Insert(\"b\", \"banana\"), Get(\"a\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}

	// Instructor's test cases.
	// Insert("a", "apple")
	// Insert("b", "banana")
	// Get("b")
	fmt.Println("Get(\"b\")")
	ret = mpt.Get("b")
	fmt.Println(ret) // banana
	if ret != "banana" {
		t.Errorf("Instructor's test case. Insert(\"a\", \"apple\"), Insert(\"b\", \"banana\"), Get(\"b\") \n"+
			"Expected %s, but was %s", "banana", ret)
	}

	// Instructor's test cases. (Case D-1)
	// Insert("a", "apple")
	// Insert("p", "banana")
	// Insert("abc", "new")
	// Get("a"), Get("p"), Get("abc")
	fmt.Println("Instructor's test cases (Case D-1): ")
	mpt = p1.NewMPT()
	fmt.Println("Insert(\"a\", \"apple\")")
	mpt.Insert("a", "apple")
	fmt.Println("Insert(\"p\", \"banana\")")
	mpt.Insert("p", "banana")
	fmt.Println("Insert(\"abc\", \"new\")")
	mpt.Insert("abc", "new")
	fmt.Println("Get(\"a\")")
	ret = mpt.Get("a")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("Instructor's test case. Insert(\"a\", \"apple\"), Insert(\"p\", \"banana\"), Insert(\"abc\", \"new\"), Get(\"a\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}
	fmt.Println("Get(\"p\")")
	ret = mpt.Get("p")
	fmt.Println(ret) // banana
	if ret != "banana" {
		t.Errorf("Instructor's test case. Insert(\"a\", \"apple\"), Insert(\"p\", \"banana\"), Insert(\"abc\", \"new\"), Get(\"p\") \n"+
			"Expected %s, but was %s", "banana", ret)
	}
	fmt.Println("Get(\"abc\")")
	ret = mpt.Get("abc")
	fmt.Println(ret) // new
	if ret != "new" {
		t.Errorf("Instructor's test case. Insert(\"a\", \"apple\"), Insert(\"p\", \"banana\"), Insert(\"abc\", \"new\"), Get(\"abc\") \n"+
			"Expected %s, but was %s", "new", ret)
	}

	// Case D-2
	// Insert("a", "apple")
	// Insert("p", "banana")
	// Insert("b", "new")
	// Get("a"), Get("p"), Get("b")
	fmt.Println("Case D-2: ")
	mpt = p1.NewMPT()
	fmt.Println("Insert(\"a\", \"apple\")")
	mpt.Insert("a", "apple")
	fmt.Println("Insert(\"p\", \"banana\")")
	mpt.Insert("p", "banana")
	fmt.Println("Insert(\"b\", \"new\")")
	mpt.Insert("b", "new")
	fmt.Println("Get(\"a\")")
	ret = mpt.Get("a")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("Instructor's test case. Insert(\"a\", \"apple\"), Insert(\"p\", \"banana\"), Insert(\"b\", \"new\"), Get(\"a\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}
	fmt.Println("Get(\"p\")")
	ret = mpt.Get("p")
	fmt.Println(ret) // banana
	if ret != "banana" {
		t.Errorf("Instructor's test case. Insert(\"a\", \"apple\"), Insert(\"p\", \"banana\"), Insert(\"b\", \"new\"), Get(\"p\") \n"+
			"Expected %s, but was %s", "banana", ret)
	}
	fmt.Println("Get(\"b\")")
	ret = mpt.Get("b")
	fmt.Println(ret) // new
	if ret != "new" {
		t.Errorf("Instructor's test case. Insert(\"a\", \"apple\"), Insert(\"p\", \"banana\"), Insert(\"b\", \"new\"), Get(\"b\") \n"+
			"Expected %s, but was %s", "new", ret)
	}

	// Del-6.
	fmt.Println("Del-6.")
	fmt.Println("Delete(\"p\")")
	mpt.Delete("p")
	fmt.Println("Get(\"a\")")
	ret = mpt.Get("a")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("Del-6. Insert(\"a\", \"apple\"), Insert(\"p\", \"banana\"), Insert(\"b\", \"new\"), Delete(\"p\"), Get(\"a\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}
	fmt.Println("Get(\"b\")")
	ret = mpt.Get("b")
	fmt.Println(ret) // new
	if ret != "new" {
		t.Errorf("Del-6. Insert(\"a\", \"apple\"), Insert(\"p\", \"banana\"), Insert(\"b\", \"new\"), Delete(\"p\"), Get(\"b\") \n"+
			"Expected %s, but was %s", "new", ret)
	}
	fmt.Println("Get(\"p\")")
	ret = mpt.Get("p")
	fmt.Println(ret) // ""
	if ret != "" {
		t.Errorf("Del-6. Insert(\"a\", \"apple\"), Insert(\"p\", \"banana\"), Insert(\"b\", \"new\"), Delete(\"p\"), Get(\"p\") \n"+
			"Expected %s, but was %s", "", ret)
	}

	// Case D-3
	// Insert("a", "apple")
	// Insert("p", "banana")
	// Insert("A", "new")
	// Get("a"), Get("p"), Get("A")
	fmt.Println("Case D-3: ")
	mpt = p1.NewMPT()
	fmt.Println("Insert(\"a\", \"apple\")")
	mpt.Insert("a", "apple")
	fmt.Println("Insert(\"p\", \"banana\")")
	mpt.Insert("p", "banana")
	fmt.Println("Insert(\"b\", \"new\")")
	mpt.Insert("A", "new")
	fmt.Println("Get(\"a\")")
	ret = mpt.Get("a")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("Instructor's test case. Insert(\"a\", \"apple\"), Insert(\"p\", \"banana\"), Insert(\"A\", \"new\"), Get(\"a\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}
	fmt.Println("Get(\"p\")")
	ret = mpt.Get("p")
	fmt.Println(ret) // banana
	if ret != "banana" {
		t.Errorf("Instructor's test case. Insert(\"a\", \"apple\"), Insert(\"p\", \"banana\"), Insert(\"A\", \"new\"), Get(\"p\") \n"+
			"Expected %s, but was %s", "banana", ret)
	}
	fmt.Println("Get(\"A\")")
	ret = mpt.Get("A")
	fmt.Println(ret) // new
	if ret != "new" {
		t.Errorf("Instructor's test case. Insert(\"a\", \"apple\"), Insert(\"p\", \"banana\"), Insert(\"A\", \"new\"), Get(\"A\") \n"+
			"Expected %s, but was %s", "new", ret)
	}

	// Instructor's test cases.
	// Insert("p", "apple")
	// Insert("aa", "banana")
	// Insert("ap", "orange")
	// Get("p"), Get("aa"), Get("ap")
	fmt.Println("Instructor's test cases: ")
	mpt = p1.NewMPT()
	fmt.Println("Insert(\"p\", \"apple\")")
	mpt.Insert("p", "apple")
	fmt.Println("Insert(\"aa\", \"banana\")")
	mpt.Insert("aa", "banana")
	fmt.Println("Insert(\"ap\", \"orange\")")
	mpt.Insert("ap", "orange")
	fmt.Println("Get(\"p\")")
	ret = mpt.Get("p")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("Instructor's test case. Insert(\"p\", \"apple\"), Insert(\"aa\", \"banana\"), Insert(\"ap\", \"orange\"), Get(\"p\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}
	fmt.Println("Get(\"aa\")")
	ret = mpt.Get("aa")
	fmt.Println(ret) // banana
	if ret != "banana" {
		t.Errorf("Instructor's test case. Insert(\"p\", \"apple\"), Insert(\"aa\", \"banana\"), Insert(\"ap\", \"orange\"), Get(\"aa\") \n"+
			"Expected %s, but was %s", "banana", ret)
	}
	fmt.Println("Get(\"ap\")")
	ret = mpt.Get("ap")
	fmt.Println(ret) // orange
	if ret != "orange" {
		t.Errorf("Instructor's test case. Insert(\"p\", \"apple\"), Insert(\"aa\", \"banana\"), Insert(\"ap\", \"orange\"), Get(\"ap\") \n"+
			"Expected %s, but was %s", "orange", ret)
	}

	// Instructor's test cases. Case E-1
	// Insert("p", "apple")
	// Insert("aaaaa", "banana")
	// Insert("aaaap", "orange")
	// Insert("aa", "new")
	// Get("p"), Get("aaaaa"), Get("aaaap"), Get("aa")
	fmt.Println("Instructor's test cases (Case E-1): ")
	mpt = p1.NewMPT()
	fmt.Println("Insert(\"p\", \"apple\")")
	mpt.Insert("p", "apple")
	fmt.Println("Insert(\"aaaaa\", \"banana\")")
	mpt.Insert("aaaaa", "banana")
	fmt.Println("Insert(\"aaaap\", \"orange\")")
	mpt.Insert("aaaap", "orange")
	fmt.Println("Insert(\"aa\", \"new\")")
	mpt.Insert("aa", "new")
	fmt.Println("Get(\"p\")")
	ret = mpt.Get("p")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("Instructor's test case (Case E-1). Insert(\"p\", \"apple\"), Insert(\"aaaaa\", \"banana\"), Insert(\"aaaap\", \"orange\"), Insert(\"aa\", \"new\"), Get(\"p\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}
	fmt.Println("Get(\"aaaaa\")")
	ret = mpt.Get("aaaaa")
	fmt.Println(ret) // banana
	if ret != "banana" {
		t.Errorf("Instructor's test case (Case E-1). Insert(\"p\", \"apple\"), Insert(\"aaaaa\", \"banana\"), Insert(\"aaaap\", \"orange\"), Insert(\"aa\", \"new\"), Get(\"aaaaa\") \n"+
			"Expected %s, but was %s", "banana", ret)
	}
	fmt.Println("Get(\"aaaap\")")
	ret = mpt.Get("aaaap")
	fmt.Println(ret) // orange
	if ret != "orange" {
		t.Errorf("Instructor's test case (Case E-1). Insert(\"p\", \"apple\"), Insert(\"aaaaa\", \"banana\"), Insert(\"aaaap\", \"orange\"), Insert(\"aa\", \"new\"), Get(\"aaaap\") \n"+
			"Expected %s, but was %s", "orange", ret)
	}
	fmt.Println("Get(\"aa\")")
	ret = mpt.Get("aa")
	fmt.Println(ret) // new
	if ret != "new" {
		t.Errorf("Instructor's test case (Case E-1). Insert(\"p\", \"apple\"), Insert(\"aaaaa\", \"banana\"), Insert(\"aaaap\", \"orange\"), Insert(\"aa\", \"new\"), Get(\"aa\") \n"+
			"Expected %s, but was %s", "new", ret)
	}

	// Case E-1b
	// Insert("p", "apple")
	// Insert("aaaa", "banana")
	// Insert("aaap", "orange")
	// Insert("a", "new")
	// Get("p"), Get("aaaa"), Get("aaap"), Get("a")
	fmt.Println("Instructor's test cases (Case E-1b): ")
	mpt = p1.NewMPT()
	fmt.Println("Insert(\"p\", \"apple\")")
	mpt.Insert("p", "apple")
	fmt.Println("Insert(\"aaaa\", \"banana\")")
	mpt.Insert("aaaa", "banana")
	fmt.Println("Insert(\"aaap\", \"orange\")")
	mpt.Insert("aaap", "orange")
	fmt.Println("Insert(\"a\", \"new\")")
	mpt.Insert("a", "new")
	fmt.Println("Get(\"p\")")
	ret = mpt.Get("p")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("Case E-1b. Insert(\"p\", \"apple\"), Insert(\"aaaa\", \"banana\"), Insert(\"aaap\", \"orange\"), Insert(\"a\", \"new\"), Get(\"p\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}
	fmt.Println("Get(\"aaaa\")")
	ret = mpt.Get("aaaa")
	fmt.Println(ret) // banana
	if ret != "banana" {
		t.Errorf("Case E-1b. Insert(\"p\", \"apple\"), Insert(\"aaaa\", \"banana\"), Insert(\"aaap\", \"orange\"), Insert(\"a\", \"new\"), Get(\"aaaa\") \n"+
			"Expected %s, but was %s", "banana", ret)
	}
	fmt.Println("Get(\"aaap\")")
	ret = mpt.Get("aaap")
	fmt.Println(ret) // orange
	if ret != "orange" {
		t.Errorf("Case E-1b. Insert(\"p\", \"apple\"), Insert(\"aaaa\", \"banana\"), Insert(\"aaap\", \"orange\"), Insert(\"a\", \"new\"), Get(\"aaap\") \n"+
			"Expected %s, but was %s", "orange", ret)
	}
	fmt.Println("Get(\"a\")")
	ret = mpt.Get("a")
	fmt.Println(ret) // new
	if ret != "new" {
		t.Errorf("Case E-1b. Insert(\"p\", \"apple\"), Insert(\"aaaa\", \"banana\"), Insert(\"aaap\", \"orange\"), Insert(\"a\", \"new\"), Get(\"a\") \n"+
			"Expected %s, but was %s", "new", ret)
	}

	// Case E-2
	// Insert("p", "apple")
	// Insert("aaaaa", "banana")
	// Insert("aaaap", "orange")
	// Insert("aaaa", "new")
	// Get("p"), Get("aaaaa"), Get("aaaap"), Get("aaaa")
	fmt.Println("Case E-2: ")
	mpt = p1.NewMPT()
	fmt.Println("Insert(\"p\", \"apple\")")
	mpt.Insert("p", "apple")
	fmt.Println("Insert(\"aaaaa\", \"banana\")")
	mpt.Insert("aaaaa", "banana")
	fmt.Println("Insert(\"aaaap\", \"orange\")")
	mpt.Insert("aaaap", "orange")
	fmt.Println("Insert(\"aaaa\", \"new\")")
	mpt.Insert("aaaa", "new")
	fmt.Println("Get(\"p\")")
	ret = mpt.Get("p")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("Instructor's test case. Insert(\"p\", \"apple\"), Insert(\"aaaaa\", \"banana\"), Insert(\"aaaap\", \"orange\"), Insert(\"aaaa\", \"new\"), Get(\"p\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}
	fmt.Println("Get(\"aaaaa\")")
	ret = mpt.Get("aaaaa")
	fmt.Println(ret) // banana
	if ret != "banana" {
		t.Errorf("Instructor's test case. Insert(\"p\", \"apple\"), Insert(\"aaaaa\", \"banana\"), Insert(\"aaaap\", \"orange\"), Insert(\"aaaa\", \"new\"), Get(\"aaaaa\") \n"+
			"Expected %s, but was %s", "banana", ret)
	}
	fmt.Println("Get(\"aaaap\")")
	ret = mpt.Get("aaaap")
	fmt.Println(ret) // orange
	if ret != "orange" {
		t.Errorf("Instructor's test case. Insert(\"p\", \"apple\"), Insert(\"aaaaa\", \"banana\"), Insert(\"aaaap\", \"orange\"), Insert(\"aaaa\", \"new\"), Get(\"aaaap\") \n"+
			"Expected %s, but was %s", "orange", ret)
	}
	fmt.Println("Get(\"aaaa\")")
	ret = mpt.Get("aaaa")
	fmt.Println(ret) // new
	if ret != "new" {
		t.Errorf("Instructor's test case. Insert(\"p\", \"apple\"), Insert(\"aaaaa\", \"banana\"), Insert(\"aaaap\", \"orange\"), Insert(\"aaaa\", \"new\"), Get(\"aaaa\") \n"+
			"Expected %s, but was %s", "new", ret)
	}

	// Case E-3
	// Insert("p", "apple")
	// Insert("aaaaa", "banana")
	// Insert("aaaap", "orange")
	// Insert("aaaaA", "new")
	// Get("p"), Get("aaaaa"), Get("aaaap"), Get("aaaaA")
	fmt.Println("Case E-3: ")
	mpt = p1.NewMPT()
	fmt.Println("Insert(\"p\", \"apple\")")
	mpt.Insert("p", "apple")
	fmt.Println("Insert(\"aaaaa\", \"banana\")")
	mpt.Insert("aaaaa", "banana")
	fmt.Println("Insert(\"aaaap\", \"orange\")")
	mpt.Insert("aaaap", "orange")
	fmt.Println("Insert(\"aaaaA\", \"new\")")
	mpt.Insert("aaaaA", "new")
	fmt.Println("Get(\"p\")")
	ret = mpt.Get("p")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("Instructor's test case. Insert(\"p\", \"apple\"), Insert(\"aaaaa\", \"banana\"), Insert(\"aaaap\", \"orange\"), Insert(\"aaaaA\", \"new\"), Get(\"p\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}
	fmt.Println("Get(\"aaaaa\")")
	ret = mpt.Get("aaaaa")
	fmt.Println(ret) // banana
	if ret != "banana" {
		t.Errorf("Instructor's test case. Insert(\"p\", \"apple\"), Insert(\"aaaaa\", \"banana\"), Insert(\"aaaap\", \"orange\"), Insert(\"aaaaA\", \"new\"), Get(\"aaaaa\") \n"+
			"Expected %s, but was %s", "banana", ret)
	}
	fmt.Println("Get(\"aaaap\")")
	ret = mpt.Get("aaaap")
	fmt.Println(ret) // orange
	if ret != "orange" {
		t.Errorf("Instructor's test case. Insert(\"p\", \"apple\"), Insert(\"aaaaa\", \"banana\"), Insert(\"aaaap\", \"orange\"), Insert(\"aaaaA\", \"new\"), Get(\"aaaap\") \n"+
			"Expected %s, but was %s", "orange", ret)
	}
	fmt.Println("Get(\"aaaaA\")")
	ret = mpt.Get("aaaaA")
	fmt.Println(ret) // new
	if ret != "new" {
		t.Errorf("Instructor's test case. Insert(\"p\", \"apple\"), Insert(\"aaaaa\", \"banana\"), Insert(\"aaaap\", \"orange\"), Insert(\"aaaaA\", \"new\"), Get(\"aaaaA\") \n"+
			"Expected %s, but was %s", "new", ret)
	}

	// Case E-4
	// Insert("p", "apple")
	// Insert("aaaaa", "banana")
	// Insert("aaaap", "orange")
	// Insert("b", "new")
	// Get("p"), Get("aaaaa"), Get("aaaap"), Get("b")
	fmt.Println("Case E-4: ")
	mpt = p1.NewMPT()
	fmt.Println("Insert(\"p\", \"apple\")")
	mpt.Insert("p", "apple")
	fmt.Println("Insert(\"aaaaa\", \"banana\")")
	mpt.Insert("aaaaa", "banana")
	fmt.Println("Insert(\"aaaap\", \"orange\")")
	mpt.Insert("aaaap", "orange")
	fmt.Println("Insert(\"b\", \"new\")")
	mpt.Insert("b", "new")
	fmt.Println("Get(\"p\")")
	ret = mpt.Get("p")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("Instructor's test case. Insert(\"p\", \"apple\"), Insert(\"aaaaa\", \"banana\"), Insert(\"aaaap\", \"orange\"), Insert(\"b\", \"new\"), Get(\"p\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}
	fmt.Println("Get(\"aaaaa\")")
	ret = mpt.Get("aaaaa")
	fmt.Println(ret) // banana
	if ret != "banana" {
		t.Errorf("Instructor's test case. Insert(\"p\", \"apple\"), Insert(\"aaaaa\", \"banana\"), Insert(\"aaaap\", \"orange\"), Insert(\"b\", \"new\"), Get(\"aaaaa\") \n"+
			"Expected %s, but was %s", "banana", ret)
	}
	fmt.Println("Get(\"aaaap\")")
	ret = mpt.Get("aaaap")
	fmt.Println(ret) // orange
	if ret != "orange" {
		t.Errorf("Instructor's test case. Insert(\"p\", \"apple\"), Insert(\"aaaaa\", \"banana\"), Insert(\"aaaap\", \"orange\"), Insert(\"b\", \"new\"), Get(\"aaaap\") \n"+
			"Expected %s, but was %s", "orange", ret)
	}
	fmt.Println("Get(\"b\")")
	ret = mpt.Get("b")
	fmt.Println(ret) // new
	if ret != "new" {
		t.Errorf("Instructor's test case. Insert(\"p\", \"apple\"), Insert(\"aaaaa\", \"banana\"), Insert(\"aaaap\", \"orange\"), Insert(\"b\", \"new\"), Get(\"b\") \n"+
			"Expected %s, but was %s", "new", ret)
	}

	// Case E-4b
	// Insert("p", "apple")
	// Insert("aa", "banana")
	// Insert("ap", "orange")
	// Insert("b", "new")
	// Get("p"), Get("aa"), Get("ap"), Get("b")
	fmt.Println("Case E-4b: ")
	mpt = p1.NewMPT()
	fmt.Println("Insert(\"p\", \"apple\")")
	mpt.Insert("p", "apple")
	fmt.Println("Insert(\"aa\", \"banana\")")
	mpt.Insert("aa", "banana")
	fmt.Println("Insert(\"ap\", \"orange\")")
	mpt.Insert("ap", "orange")
	fmt.Println("Insert(\"b\", \"new\")")
	mpt.Insert("b", "new")
	fmt.Println("Get(\"p\")")
	ret = mpt.Get("p")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("Case E-4b. Insert(\"p\", \"apple\"), Insert(\"aa\", \"banana\"), Insert(\"ap\", \"orange\"), Insert(\"b\", \"new\"), Get(\"p\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}
	fmt.Println("Get(\"aa\")")
	ret = mpt.Get("aa")
	fmt.Println(ret) // banana
	if ret != "banana" {
		t.Errorf("Case E-4b. Insert(\"p\", \"apple\"), Insert(\"aa\", \"banana\"), Insert(\"ap\", \"orange\"), Insert(\"b\", \"new\"), Get(\"aa\") \n"+
			"Expected %s, but was %s", "banana", ret)
	}
	fmt.Println("Get(\"ap\")")
	ret = mpt.Get("ap")
	fmt.Println(ret) // orange
	if ret != "orange" {
		t.Errorf("Case E-4b. Insert(\"p\", \"apple\"), Insert(\"aa\", \"banana\"), Insert(\"ap\", \"orange\"), Insert(\"b\", \"new\"), Get(\"ap\") \n"+
			"Expected %s, but was %s", "orange", ret)
	}
	fmt.Println("Get(\"b\")")
	ret = mpt.Get("b")
	fmt.Println(ret) // new
	if ret != "new" {
		t.Errorf("Case E-4b. Insert(\"p\", \"apple\"), Insert(\"aa\", \"banana\"), Insert(\"ap\", \"orange\"), Insert(\"b\", \"new\"), Get(\"b\") \n"+
			"Expected %s, but was %s", "new", ret)
	}

	// Del-8
	fmt.Println("Del-8.")
	fmt.Println("Delete(\"b\")")
	mpt.Delete("b")
	fmt.Println("Get(\"p\")")
	ret = mpt.Get("p")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("Del-8. Insert(\"p\", \"apple\"), Insert(\"aa\", \"banana\"), Insert(\"ap\", \"orange\"), Insert(\"b\", \"new\"), Get(\"p\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}
	fmt.Println("Get(\"aa\")")
	ret = mpt.Get("aa")
	fmt.Println(ret) // banana
	if ret != "banana" {
		t.Errorf("Del-8. Insert(\"p\", \"apple\"), Insert(\"aa\", \"banana\"), Insert(\"ap\", \"orange\"), Insert(\"b\", \"new\"), Get(\"aa\") \n"+
			"Expected %s, but was %s", "banana", ret)
	}
	fmt.Println("Get(\"ap\")")
	ret = mpt.Get("ap")
	fmt.Println(ret) // orange
	if ret != "orange" {
		t.Errorf("Del-8. Insert(\"p\", \"apple\"), Insert(\"aa\", \"banana\"), Insert(\"ap\", \"orange\"), Insert(\"b\", \"new\"), Get(\"ap\") \n"+
			"Expected %s, but was %s", "orange", ret)
	}
	fmt.Println("Get(\"b\")")
	ret = mpt.Get("b")
	fmt.Println(ret) // ""
	if ret != "" {
		t.Errorf("Del-8. Insert(\"p\", \"apple\"), Insert(\"aa\", \"banana\"), Insert(\"ap\", \"orange\"), Insert(\"b\", \"new\"), Get(\"b\") \n"+
			"Expected %s, but was %s", "", ret)
	}

	// Del-8b.
	// Insert("p", "apple")
	// Insert("aa", "banana")
	// Insert("ap", "orange")
	// Insert("b", "new")
	// Delete("aa")
	// Get("p"), Get("aa"), Get("ap"), Get("b")
	fmt.Println("Del-8b: ")
	mpt = p1.NewMPT()
	fmt.Println("Insert(\"p\", \"apple\")")
	mpt.Insert("p", "apple")
	fmt.Println("Insert(\"aa\", \"banana\")")
	mpt.Insert("aa", "banana")
	fmt.Println("Insert(\"ap\", \"orange\")")
	mpt.Insert("ap", "orange")
	fmt.Println("Insert(\"b\", \"new\")")
	mpt.Insert("b", "new")
	fmt.Println("Delete(\"aa\")")
	mpt.Delete("aa")
	fmt.Println("Get(\"p\")")
	ret = mpt.Get("p")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("Del-8b. Insert(\"p\", \"apple\"), Insert(\"aa\", \"banana\"), Insert(\"ap\", \"orange\"), Insert(\"aa\", \"new\"), Get(\"p\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}
	fmt.Println("Get(\"aa\")")
	ret = mpt.Get("aa")
	fmt.Println(ret) // ""
	if ret != "" {
		t.Errorf("Del-8b. Insert(\"p\", \"apple\"), Insert(\"aa\", \"banana\"), Insert(\"ap\", \"orange\"), Insert(\"aa\", \"new\"), Get(\"aa\") \n"+
			"Expected %s, but was %s", "", ret)
	}
	fmt.Println("Get(\"ap\")")
	ret = mpt.Get("ap")
	fmt.Println(ret) // orange
	if ret != "orange" {
		t.Errorf("Del-8b. Insert(\"p\", \"apple\"), Insert(\"aa\", \"banana\"), Insert(\"ap\", \"orange\"), Insert(\"aa\", \"new\"), Get(\"ap\") \n"+
			"Expected %s, but was %s", "orange", ret)
	}
	fmt.Println("Get(\"b\")")
	ret = mpt.Get("b")
	fmt.Println(ret) // new
	if ret != "new" {
		t.Errorf("Del-8b. Insert(\"p\", \"apple\"), Insert(\"aa\", \"banana\"), Insert(\"ap\", \"orange\"), Insert(\"aa\", \"new\"), Get(\"b\") \n"+
			"Expected %s, but was %s", "new", ret)
	}

	// Case E-5
	// Insert("p", "apple")
	// Insert("aaaaa", "banana")
	// Insert("aaaap", "orange")
	// Insert("aaA", "new")
	// Get("p"), Get("aaaaa"), Get("aaaap"), Get("aaA")
	fmt.Println("Case E-5: ")
	mpt = p1.NewMPT()
	fmt.Println("Insert(\"p\", \"apple\")")
	mpt.Insert("p", "apple")
	fmt.Println("Insert(\"aaaaa\", \"banana\")")
	mpt.Insert("aaaaa", "banana")
	fmt.Println("Insert(\"aaaap\", \"orange\")")
	mpt.Insert("aaaap", "orange")
	fmt.Println("Insert(\"aaA\", \"new\")")
	mpt.Insert("aaA", "new")
	fmt.Println("Get(\"p\")")
	ret = mpt.Get("p")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("Instructor's test case. Insert(\"p\", \"apple\"), Insert(\"aaaaa\", \"banana\"), Insert(\"aaaap\", \"orange\"), Insert(\"aaA\", \"new\"), Get(\"p\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}
	fmt.Println("Get(\"aaaaa\")")
	ret = mpt.Get("aaaaa")
	fmt.Println(ret) // banana
	if ret != "banana" {
		t.Errorf("Instructor's test case. Insert(\"p\", \"apple\"), Insert(\"aaaaa\", \"banana\"), Insert(\"aaaap\", \"orange\"), Insert(\"aaA\", \"new\"), Get(\"aaaaa\") \n"+
			"Expected %s, but was %s", "banana", ret)
	}
	fmt.Println("Get(\"aaaap\")")
	ret = mpt.Get("aaaap")
	fmt.Println(ret) // orange
	if ret != "orange" {
		t.Errorf("Instructor's test case. Insert(\"p\", \"apple\"), Insert(\"aaaaa\", \"banana\"), Insert(\"aaaap\", \"orange\"), Insert(\"aaA\", \"new\"), Get(\"aaaap\") \n"+
			"Expected %s, but was %s", "orange", ret)
	}
	fmt.Println("Get(\"aaA\")")
	ret = mpt.Get("aaA")
	fmt.Println(ret) // new
	if ret != "new" {
		t.Errorf("Instructor's test case. Insert(\"p\", \"apple\"), Insert(\"aaaaa\", \"banana\"), Insert(\"aaaap\", \"orange\"), Insert(\"aaA\", \"new\"), Get(\"aaA\") \n"+
			"Expected %s, but was %s", "new", ret)
	}

	// For Del-2.
	// Insert("aaa", "apple")
	// Insert("aap", "orange")
	// Insert("aaaa", "banana")
	// Get("aaa"), Get("aap"), Get("aaaa")
	fmt.Println("For Del-2: ")
	mpt = p1.NewMPT()
	fmt.Println("Insert(\"aaa\", \"apple\")")
	mpt.Insert("aaa", "apple")
	fmt.Println("Insert(\"aap\", \"orange\")")
	mpt.Insert("aap", "orange")
	fmt.Println("Insert(\"aaaa\", \"banana\")")
	mpt.Insert("aaaa", "banana")
	fmt.Println("Get(\"aaa\")")
	ret = mpt.Get("aaa")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("For Del-2. Insert(\"aaa\", \"apple\"), Insert(\"aap\", \"orange\"), Insert(\"aaaa\", \"banana\"), Get(\"aaa\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}
	fmt.Println("Get(\"aap\")")
	ret = mpt.Get("aap")
	fmt.Println(ret) // orange
	if ret != "orange" {
		t.Errorf("For Del-2. Insert(\"aaa\", \"apple\"), Insert(\"aap\", \"orange\"), Insert(\"aaaa\", \"banana\"), Get(\"aap\") \n"+
			"Expected %s, but was %s", "orange", ret)
	}
	fmt.Println("Get(\"aaaa\")")
	ret = mpt.Get("aaaa")
	fmt.Println(ret) // banana
	if ret != "banana" {
		t.Errorf("For Del-2. Insert(\"aaa\", \"apple\"), Insert(\"aap\", \"orange\"), Insert(\"aaaa\", \"banana\"), Get(\"aaaa\") \n"+
			"Expected %s, but was %s", "banana", ret)
	}

	// Del-2.
	fmt.Println("Del-2.")
	fmt.Println("Delete(\"aap\")")
	mpt.Delete("aap")
	fmt.Println("Get(\"aaa\")")
	ret = mpt.Get("aaa")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("Del-2. Insert(\"aaa\", \"apple\"), Insert(\"aap\", \"orange\"), Insert(\"aaaa\", \"banana\"), Delete(\"aap\"), Get(\"aaa\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}
	fmt.Println("Get(\"aaaa\")")
	ret = mpt.Get("aaaa")
	fmt.Println(ret) // banana
	if ret != "banana" {
		t.Errorf("Del-2. Insert(\"aaa\", \"apple\"), Insert(\"aap\", \"orange\"), Insert(\"aaaa\", \"banana\"), Delete(\"aap\"), Get(\"aaaa\") \n"+
			"Expected %s, but was %s", "", ret)
	}
	fmt.Println("Get(\"aap\")")
	ret = mpt.Get("aap")
	fmt.Println(ret) // ""
	if ret != "" {
		t.Errorf("Del-2. Insert(\"aaa\", \"apple\"), Insert(\"aap\", \"orange\"), Insert(\"aaaa\", \"banana\"), Delete(\"aap\"), Get(\"aap\") \n"+
			"Expected %s, but was %s", "", ret)
	}

	// For Del-5.
	// Insert("aaa", "apple")
	// Insert("aap", "orange")
	// Insert("aaaa", "banana")
	// Insert("aa", "kiwi")
	// Get("aaa"), Get("aap"), Get("aaaa"), Get("aa")
	fmt.Println("For Del-5: ")
	mpt = p1.NewMPT()
	fmt.Println("Insert(\"aaa\", \"apple\")")
	mpt.Insert("aaa", "apple")
	fmt.Println("Insert(\"aap\", \"orange\")")
	mpt.Insert("aap", "orange")
	fmt.Println("Insert(\"aaaa\", \"banana\")")
	mpt.Insert("aaaa", "banana")
	fmt.Println("Insert(\"aa\", \"kiwi\")")
	mpt.Insert("aa", "kiwi")
	fmt.Println("Get(\"aaa\")")
	ret = mpt.Get("aaa")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("For Del-5. Insert(\"aaa\", \"apple\"), Insert(\"aap\", \"orange\"), Insert(\"aaaa\", \"banana\"), Insert(\"aa\", \"kiwi\"), Get(\"aaa\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}
	fmt.Println("Get(\"aap\")")
	ret = mpt.Get("aap")
	fmt.Println(ret) // orange
	if ret != "orange" {
		t.Errorf("For Del-5. Insert(\"aaa\", \"apple\"), Insert(\"aap\", \"orange\"), Insert(\"aaaa\", \"banana\"), Insert(\"aa\", \"kiwi\"), Get(\"aap\") \n"+
			"Expected %s, but was %s", "orange", ret)
	}
	fmt.Println("Get(\"aaaa\")")
	ret = mpt.Get("aaaa")
	fmt.Println(ret) // banana
	if ret != "banana" {
		t.Errorf("For Del-5. Insert(\"aaa\", \"apple\"), Insert(\"aap\", \"orange\"), Insert(\"aaaa\", \"banana\"), Insert(\"aa\", \"kiwi\"), Get(\"aaaa\") \n"+
			"Expected %s, but was %s", "banana", ret)
	}
	fmt.Println("Get(\"aa\")")
	ret = mpt.Get("aa")
	fmt.Println(ret) // kiwi
	if ret != "kiwi" {
		t.Errorf("For Del-5. Insert(\"aaa\", \"apple\"), Insert(\"aap\", \"orange\"), Insert(\"aaaa\", \"banana\"), Insert(\"aa\", \"kiwi\"), Get(\"aa\") \n"+
			"Expected %s, but was %s", "kiwi", ret)
	}

	// Del-5.
	fmt.Println("Del-5.")
	fmt.Println("Delete(\"aap\")")
	mpt.Delete("aap")
	ret = mpt.Get("aaa")
	fmt.Println(ret) // apple
	ret = mpt.Get("aaaa")
	fmt.Println(ret) // banana
	ret = mpt.Get("aap")
	fmt.Println(ret) // ""
	ret = mpt.Get("aa")
	fmt.Println(ret) // kiwi

	fmt.Println("Delete(\"aa\")")
	mpt.Delete("aa")
	fmt.Println("Get(\"aaa\")")
	ret = mpt.Get("aaa")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("Del-5. Insert(\"aaa\", \"apple\"), Insert(\"aap\", \"orange\"), Insert(\"aaaa\", \"banana\"), Delete(\"aap\"), Delete(\"aa\"), Get(\"aaa\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}
	fmt.Println("Get(\"aaaa\")")
	ret = mpt.Get("aaaa")
	fmt.Println(ret) // banana
	if ret != "banana" {
		t.Errorf("Del-5. Insert(\"aaa\", \"apple\"), Insert(\"aap\", \"orange\"), Insert(\"aaaa\", \"banana\"), Delete(\"aap\"), Delete(\"aa\"), Get(\"aaaa\") \n"+
			"Expected %s, but was %s", "banana", ret)
	}
	fmt.Println("Get(\"aap\")")
	ret = mpt.Get("aap")
	fmt.Println(ret) // ""
	if ret != "" {
		t.Errorf("Del-5. Insert(\"aaa\", \"apple\"), Insert(\"aap\", \"orange\"), Insert(\"aaaa\", \"banana\"), Delete(\"aap\"), Delete(\"aa\"), Get(\"aap\") \n"+
			"Expected %s, but was %s", "", ret)
	}
	fmt.Println("Get(\"aa\")")
	ret = mpt.Get("aa")
	fmt.Println(ret) // ""
	if ret != "" {
		t.Errorf("Del-5. Insert(\"aaa\", \"apple\"), Insert(\"aap\", \"orange\"), Insert(\"aaaa\", \"banana\"), Delete(\"aap\"), Delete(\"aa\"), Get(\"aa\") \n"+
			"Expected %s, but was %s", "", ret)
	}

	// For Del-2b.
	// Insert("p", "apple")
	// Insert("aaa", "orange")
	// Insert("aap", "banana")
	// Get("p"), Get("aaa"), Get("aap")
	fmt.Println("For Del-2b: ")
	mpt = p1.NewMPT()
	fmt.Println("Insert(\"p\", \"apple\")")
	mpt.Insert("p", "apple")
	fmt.Println("Insert(\"aaa\", \"orange\")")
	mpt.Insert("aaa", "orange")
	fmt.Println("Insert(\"aap\", \"banana\")")
	mpt.Insert("aap", "banana")
	fmt.Println("Get(\"p\")")
	ret = mpt.Get("p")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("For Del-2b. Insert(\"p\", \"apple\"), Insert(\"aaa\", \"orange\"), Insert(\"aap\", \"banana\"), Get(\"p\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}
	fmt.Println("Get(\"aaa\")")
	ret = mpt.Get("aaa")
	fmt.Println(ret) // orange
	if ret != "orange" {
		t.Errorf("For Del-2b. Insert(\"p\", \"apple\"), Insert(\"aaa\", \"orange\"), Insert(\"aap\", \"banana\"), Get(\"aaa\") \n"+
			"Expected %s, but was %s", "orange", ret)
	}
	fmt.Println("Get(\"aap\")")
	ret = mpt.Get("aap")
	fmt.Println(ret) // banana
	if ret != "banana" {
		t.Errorf("For Del-2b. Insert(\"p\", \"apple\"), Insert(\"aaa\", \"orange\"), Insert(\"aap\", \"banana\"), Get(\"aap\") \n"+
			"Expected %s, but was %s", "banana", ret)
	}

	// Del-2b.
	fmt.Println("Del-2b.")
	fmt.Println("Delete(\"p\")")
	mpt.Delete("p")
	fmt.Println("Get(\"aaa\")")
	ret = mpt.Get("aaa")
	fmt.Println(ret) // orange
	if ret != "orange" {
		t.Errorf("Del-2b. Insert(\"p\", \"apple\"), Insert(\"aaa\", \"orange\"), Insert(\"aap\", \"banana\"), Delete(\"p\"), Get(\"aaa\") \n"+
			"Expected %s, but was %s", "orange", ret)
	}
	fmt.Println("Get(\"aap\")")
	ret = mpt.Get("aap")
	fmt.Println(ret) // banana
	if ret != "banana" {
		t.Errorf("Del-2b. Insert(\"p\", \"apple\"), Insert(\"aaa\", \"orange\"), Insert(\"aap\", \"banana\"), Delete(\"p\"), Get(\"aap\") \n"+
			"Expected %s, but was %s", "banana", ret)
	}
	fmt.Println("Get(\"p\")")
	ret = mpt.Get("p")
	fmt.Println(ret) // ""
	if ret != "" {
		t.Errorf("Del-2b. Insert(\"p\", \"apple\"), Insert(\"aaa\", \"orange\"), Insert(\"aap\", \"banana\"), Delete(\"p\"), Get(\"p\") \n"+
			"Expected %s, but was %s", "", ret)
	}

	// For Del-7.
	// Insert("aa", "apple")
	// Insert("a", "orange")
	// Insert("ab", "kiwi")
	// Get("aa"), Get("a"), Get("ab")
	fmt.Println("For Del-7: ")
	mpt = p1.NewMPT()
	fmt.Println("Insert(\"aa\", \"apple\")")
	mpt.Insert("aa", "apple")
	fmt.Println("Insert(\"a\", \"orange\")")
	mpt.Insert("a", "orange")
	fmt.Println("Insert(\"ab\", \"kiwi\")")
	mpt.Insert("ab", "kiwi")
	fmt.Println("Get(\"aa\")")
	ret = mpt.Get("aa")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("For Del-7: Insert(\"aa\", \"apple\"), Insert(\"a\", \"orange\"), Insert(\"ab\", \"kiwi\"), Get(\"aa\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}
	fmt.Println("Get(\"a\")")
	ret = mpt.Get("a")
	fmt.Println(ret) // orange
	if ret != "orange" {
		t.Errorf("For Del-7: Insert(\"aa\", \"apple\"), Insert(\"a\", \"orange\"), Insert(\"ab\", \"kiwi\"), Get(\"a\") \n"+
			"Expected %s, but was %s", "orange", ret)
	}
	fmt.Println("Get(\"ab\")")
	ret = mpt.Get("ab")
	fmt.Println(ret) // kiwi
	if ret != "kiwi" {
		t.Errorf("For Del-7: Insert(\"aa\", \"apple\"), Insert(\"a\", \"orange\"), Insert(\"ab\", \"kiwi\"), Get(\"ab\") \n"+
			"Expected %s, but was %s", "kiwi", ret)
	}

	// Del-7.
	fmt.Println("Del-7.")
	fmt.Println("Delete(\"a\")")
	mpt.Delete("a")
	fmt.Println("Get(\"aa\")")
	ret = mpt.Get("aa")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("For Del-7: Insert(\"aa\", \"apple\"), Insert(\"a\", \"orange\"), Insert(\"ab\", \"kiwi\"), Get(\"aa\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}
	fmt.Println("Get(\"a\")")
	ret = mpt.Get("a")
	fmt.Println(ret) // ""
	if ret != "" {
		t.Errorf("For Del-7: Insert(\"aa\", \"apple\"), Insert(\"a\", \"orange\"), Insert(\"ab\", \"kiwi\"), Get(\"a\") \n"+
			"Expected %s, but was %s", "", ret)
	}
	fmt.Println("Get(\"ab\")")
	ret = mpt.Get("ab")
	fmt.Println(ret) // kiwi
	if ret != "kiwi" {
		t.Errorf("For Del-7: Insert(\"aa\", \"apple\"), Insert(\"a\", \"orange\"), Insert(\"ab\", \"kiwi\"), Get(\"ab\") \n"+
			"Expected %s, but was %s", "kiwi", ret)
	}

}
