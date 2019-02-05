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
	mpt.Insert("a", "apple")
	ret := mpt.Get("a")
	fmt.Println(ret)

}
