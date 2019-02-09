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

	// // Test prefixLen()
	// a := []uint8{2, 2, 3, 4}
	// b := []uint8{2, 2, 3, 5}
	// p1.Test_prefixLen(a, b) // 3
	// b = []uint8{0, 2, 3, 5}
	// p1.Test_prefixLen(a, b) // 0
	// b = []uint8{2, 2, 3, 4}
	// p1.Test_prefixLen(a, b) // 4

	// Test Insert(), Get() start.
	fmt.Println("Test Insert(), Get(): ")

	mpt := p1.NewMPT()
	var ret string

	// Insert("a", "apple")
	// Get("a")
	fmt.Println("Insert(\"a\", \"apple\")")
	mpt.Insert("a", "apple")
	fmt.Println("Get(\"a\")")
	ret = mpt.Get("a")
	fmt.Println(ret) // apple
	if ret != "apple" {
		t.Errorf("Expected %s, but was %s", "apple", ret)
	}

	// Case A (Exact match):
	// Insert("a", "apple")
	// Insert("a", "orange")
	// Get("a")
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

	// Case B-2 (Partial match):
	// Insert("a", "apple")
	// Insert("b", "orange")
	// Get("a")
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

	// Case B-2 (Partial match):
	// Insert("a", "apple")
	// Insert("b", "orange")
	// Get("b")
	fmt.Println("Get(\"b\")")
	ret = mpt.Get("b")
	fmt.Println(ret) // orange
	if ret != "orange" {
		t.Errorf("Case B-2 (Partial match): Insert(\"a\", \"apple\"), Insert(\"b\", \"orange\"), Get(\"b\") \n"+
			"Expected %s, but was %s", "orange", ret)
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

	// Case C (Mismatch).
	// Insert("a", "apple")
	// Insert("p", "orange")
	// Get("a")
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

	// Case C (Mismatch).
	// Insert("a", "apple")
	// Insert("p", "orange")
	// Get("p")
	fmt.Println("Get(\"p\")")
	ret = mpt.Get("p")
	fmt.Println(ret) // orange
	if ret != "orange" {
		t.Errorf("Case C (Mismatch): Insert(\"a\", \"apple\"), Insert(\"p\", \"orange\"), Get(\"p\") \n"+
			"Expected %s, but was %s", "orange", ret)
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

	// Instructor's test cases. (Case D)
	// Insert("a", "apple")
	// Insert("p", "banana")
	// Insert("abc", "new")
	// Get("a"), Get("p"), Get("new")
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

	// Instructor's test cases.
	// Insert("p", "apple")
	// Insert("aa", "banana")
	// Insert("ap", "orange")
	// Get("p"), Get("aa"), Get("ap")
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

	// Instructor's test cases.
	// Insert("p", "apple")
	// Insert("aaaaa", "banana")
	// Insert("aaaap", "orange")
	// Insert("aa", "new")
	// Get("p"), Get("aaaaa"), Get("aaaap"), Get("aa")
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
		t.Errorf("Instructor's test case. Insert(\"p\", \"apple\"), Insert(\"aaaaa\", \"banana\"), Insert(\"aaaap\", \"orange\"), Insert(\"aa\", \"new\"), Get(\"p\") \n"+
			"Expected %s, but was %s", "apple", ret)
	}
	fmt.Println("Get(\"aaaaa\")")
	ret = mpt.Get("aaaaa")
	fmt.Println(ret) // ??
	if ret != "banana" {
		t.Errorf("Instructor's test case. Insert(\"p\", \"apple\"), Insert(\"aaaaa\", \"banana\"), Insert(\"aaaap\", \"orange\"), Insert(\"aa\", \"new\"), Get(\"aaaaa\") \n"+
			"Expected %s, but was %s", "banana", ret)
	}
	// fmt.Println("Get(\"aaaap\")")
	// ret = mpt.Get("aaaap")
	// fmt.Println(ret) // ??
	// if ret != "orange" {
	// 	t.Errorf("Instructor's test case. Insert(\"p\", \"apple\"), Insert(\"aaaaa\", \"banana\"), Insert(\"aaaap\", \"orange\"), Insert(\"aa\", \"new\"), Get(\"aaaap\") \n"+
	// 		"Expected %s, but was %s", "orange", ret)
	// }
	// fmt.Println("Get(\"aa\")")
	// ret = mpt.Get("aa")
	// fmt.Println(ret) // new
	// if ret != "new" {
	// 	t.Errorf("Instructor's test case. Insert(\"p\", \"apple\"), Insert(\"aaaaa\", \"banana\"), Insert(\"aaaap\", \"orange\"), Insert(\"aa\", \"new\"), Get(\"aa\") \n"+
	// 		"Expected %s, but was %s", "new", ret)
	// }

}
