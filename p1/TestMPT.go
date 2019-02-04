package p1

import (
	"fmt"
)

// Print the output of convert_string_to_hex()
func Test_convert_string_to_hex(key string) []uint8 {
	ret := convert_string_to_hex(key)
	fmt.Println(ret)
	return ret
}

// Print the output of compact_encode()
func Test_compact_encode(hex_array []uint8) []uint8 {
	ret := compact_encode(hex_array)
	fmt.Println(ret)
	return ret
}
