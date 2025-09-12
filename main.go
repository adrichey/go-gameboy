package main

import "fmt"

// Special thanks to https://rylev.github.io/DMG-01/public/book/cpu/registers.html
// for his Rust write-up to make this a possibility.

// The Gameboy CPU is composed of 8 different "registers".
// It is an 8-bit CPU meaning each register can hold 8 bits (1 byte) of data.
type Registers struct {
	a uint8
	b uint8
	c uint8
	d uint8
	e uint8
	f uint8
	h uint8
	l uint8
}

// While the CPU only has 8 bit registers, there are instructions that allow
// the game to read and write 16 bits (i.e. 2 bytes) at the same time.
// Therefore, we'll need the ability to read an write these "virtual" 16 bit registers.
// The virtual registers are the following combinations: "af", "bc", "de", "hl"
func (r *Registers) getBC() uint16 {
	// Shift over register "b" to the left by 8 bits
	// Then "append" register "c" to the end using bitwise OR
	return uint16(r.b)<<8 | uint16(r.c)
}

func (r *Registers) setBC(value uint16) {
	// Mask the first byte and shift it right 8 places, preserving the first byte in register "b"
	r.b = uint8((value & 0xFF00) >> 8)

	// Mask the second byte to convert and preserve it to 8bits
	r.c = uint8(value & 0x00FF)
}

func (r *Registers) getDE() uint16 {
	return uint16(r.d)<<8 | uint16(r.e)
}

func (r *Registers) setDE(value uint16) {
	r.d = uint8((value & 0xFF00) >> 8)
	r.e = uint8(value & 0x00FF)
}

func (r *Registers) setHL(value uint16) {
	r.h = uint8((value & 0xFF00) >> 8)
	r.l = uint8(value & 0x00FF)
}

func (r *Registers) getHL() uint16 {
	return uint16(r.h)<<8 | uint16(r.h)
}

// TODO - Need an implementation for this based on flags register implementation
func (r *Registers) getAF() uint16 {
	return uint16(r.a)<<8 | uint16(r.f)
}

func (r *Registers) setAF(value uint16) {
	r.a = uint8((value & 0xFF00) >> 8)
	r.f = uint8(value & 0x00FF)
}

func main() {
	r := Registers{}
	r.b = uint8(12) // 1100
	r.c = uint8(10) // 1010

	fmt.Printf("%8b\n", r.b)
	fmt.Printf("%8b\n", r.c)
	fmt.Printf("%16b\n", r.getBC())

	r.setBC(uint16(64251)) // 1111101011111011

	fmt.Printf("%8b\n", r.b)
	fmt.Printf("%8b\n", r.c)
	fmt.Printf("%16b\n", r.getBC())

	var i uint8 = 1
	fmt.Printf("%8b\n", i)
	fmt.Printf("%8b\n", i<<7)

	fmt.Println("EXITED")
}
