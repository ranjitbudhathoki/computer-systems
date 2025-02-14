package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("./testdata/150.uint64")

	if err != nil {
		log.Fatal("error reading a file")

	}
	fmt.Println(data)
	num := binary.BigEndian.Uint64(data)
	fmt.Println(num)

	varint := decoder(num)
	fmt.Println("varint", varint)

}

func decoder(num uint64) []byte {
	var varintArr []byte
	for num > 0 {
		bits := num & 0x7F // extracting last 7 bits
		num >>= 7          // removing the extracted 7 bits from a number
		if num > 0 {
			bits = bits | 0x80 // assigning MSB to extracted bit
		}
		varintArr = append(varintArr, uint8(bits))

	}
	return varintArr
}
