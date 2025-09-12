package main

import (
	"fmt"

	"github.com/adrichey/go-gameboy/emulator"
)

// Special thanks to https://rylev.github.io/DMG-01/public/book/cpu/registers.html
// for his Rust write-up to make this a possibility.

func main() {
	r := emulator.Registers{}
	r.B = uint8(12) // 1100
	r.C = uint8(10) // 1010

	fmt.Printf("%8b\n", r.B)
	fmt.Printf("%8b\n", r.C)
	fmt.Printf("%16b\n", r.GetBC())

	r.SetBC(uint16(64251)) // 1111101011111011

	fmt.Printf("%8b\n", r.B)
	fmt.Printf("%8b\n", r.C)
	fmt.Printf("%16b\n", r.GetBC())

	var i uint8 = 1
	fmt.Printf("%8b\n", i)
	fmt.Printf("%8b\n", i<<7)

	fmt.Println("EXITED")
}
