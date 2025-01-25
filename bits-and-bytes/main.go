package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("150.uint64")
	num := binary.BigEndian.Uint64(data)

	fmt.Println(encode(num))
}

func encode(number uint64) []byte {
	var out []byte
	for number > 0 {
		/*
			// we have to break a number in 7 bits chunk and the MSB acts like a continuation bit
			// bitwise & with bit mask of 0b01111111 (gives us the last 7 bits of a number)
			// we can also modulo it by 128 as it gives the remainder which is lowest  bit (but bitwise is performant than airthmetic operations)
		*/
		bits := number & 0x7f
		/*
			// after extracting lowest 7 bit, we now have to change the number to the it subtract the extracted 7 bit from the number
			// we can shift the number to left by 7 bits to fill the position of the extracted bit
		*/
		number >>= 7

		/*
			// now we have to add MSB to the extractd 7 bits, for that we have to first check is there are any bits to follow
			// if there is still bits to follow add 1 to the MSB of the extracted 7 bits
			// if we do bitwise or with the MSB 1 it sets the MSB to 1
		*/

		if number > 0 {
			bits |= 0b10000000
		}
		out = append(out, byte(bits))

	}
	return out

}
