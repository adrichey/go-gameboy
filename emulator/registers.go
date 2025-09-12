package emulator

const ZERO_FLAG_BYTE_POSITION uint8 = 7
const SUBTRACT_FLAG_BYTE_POSITION uint8 = 6
const HALF_CARRY_FLAG_BYTE_POSITION uint8 = 5
const CARRY_FLAG_BYTE_POSITION uint8 = 4

// The Gameboy CPU is composed of 8 different "registers".
// It is an 8-bit CPU meaning each register can hold 8 bits (1 byte) of data.
type Registers struct {
	A uint8
	B uint8
	C uint8
	D uint8
	E uint8
	F FlagsRegister
	H uint8
	L uint8
}

// While the CPU only has 8 bit registers, there are instructions that allow
// the game to read and write 16 bits (i.e. 2 bytes) at the same time.
// Therefore, we'll need the ability to read an write these "virtual" 16 bit registers.
// The virtual registers are the following combinations: "af", "bc", "de", "hl"
func (r *Registers) GetBC() uint16 {
	// Shift over register "b" to the left by 8 bits
	// Then "append" register "c" to the end using bitwise OR
	return uint16(r.B)<<8 | uint16(r.C)
}

func (r *Registers) SetBC(value uint16) {
	// Mask the first byte and shift it right 8 places, preserving the first byte in register "b"
	r.B = uint8((value & 0xFF00) >> 8)

	// Mask the second byte to convert and preserve it to 8bits
	r.C = uint8(value & 0x00FF)
}

func (r *Registers) GetDE() uint16 {
	return uint16(r.D)<<8 | uint16(r.E)
}

func (r *Registers) SetDE(value uint16) {
	r.D = uint8((value & 0xFF00) >> 8)
	r.E = uint8(value & 0x00FF)
}

func (r *Registers) SetHL(value uint16) {
	r.H = uint8((value & 0xFF00) >> 8)
	r.L = uint8(value & 0x00FF)
}

func (r *Registers) GetHL() uint16 {
	return uint16(r.H)<<8 | uint16(r.L)
}

func (r *Registers) GetAF() uint16 {
	return uint16(r.A)<<8 | uint16(r.F.GetByteFromFlags())
}

func (r *Registers) SetAF(value uint16) {
	r.A = uint8((value & 0xFF00) >> 8)
	r.F.SetFlagsFromByte(uint8(value & 0x00FF))
}

// The "f" register is a special register called the "flags" register.
// The lower four bits of the register are always 0s and the CPU
// automatically writes to the upper four bits when certain things happen.
// In other words, the CPU "flags" certain states.
/*
   ┌-> Carry
 ┌-+> Subtraction
 | |
1111 0000
| |
└-+> Zero
  └-> Half Carry
*/
type FlagsRegister struct {
	Zero      bool
	Subtract  bool
	HalfCarry bool
	Carry     bool
}

func (f *FlagsRegister) GetByteFromFlags() (b uint8) {
	var zeroByte uint8
	var subtractByte uint8
	var halfCarryByte uint8
	var carryByte uint8

	if f.Zero {
		zeroByte = 1 << ZERO_FLAG_BYTE_POSITION
	} else {
		zeroByte = 0
	}

	if f.Subtract {
		subtractByte = 1 << SUBTRACT_FLAG_BYTE_POSITION
	} else {
		subtractByte = 0
	}

	if f.HalfCarry {
		halfCarryByte = 1 << HALF_CARRY_FLAG_BYTE_POSITION
	} else {
		halfCarryByte = 0
	}

	if f.Carry {
		carryByte = 1 << CARRY_FLAG_BYTE_POSITION
	} else {
		carryByte = 0
	}

	return zeroByte | subtractByte | halfCarryByte | carryByte
}

func (f *FlagsRegister) SetFlagsFromByte(b uint8) {
	if ((b >> ZERO_FLAG_BYTE_POSITION) & 0b1) == 1 {
		f.Zero = true
	} else {
		f.Zero = false
	}

	if ((b >> SUBTRACT_FLAG_BYTE_POSITION) & 0b1) == 1 {
		f.Subtract = true
	} else {
		f.Subtract = false
	}

	if ((b >> HALF_CARRY_FLAG_BYTE_POSITION) & 0b1) == 1 {
		f.HalfCarry = true
	} else {
		f.HalfCarry = false
	}

	if ((b >> CARRY_FLAG_BYTE_POSITION) & 0b1) == 1 {
		f.Carry = true
	} else {
		f.Carry = false
	}
}
