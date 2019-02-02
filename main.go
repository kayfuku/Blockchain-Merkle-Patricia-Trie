package main

import "cs686/cs686_blockchain_P1_Go_skeleton/p1"

func main() {
	mpt := p1.NewMPT()
	mpt.Insert("", "")

	hex_array := []uint8{1, 6, 1}
	p1.Test_compact_encode(hex_array) // [17 97]
	key := "do"
	p1.Test_convert_string_to_hex(key) // [6 4 6 15 16]

}
