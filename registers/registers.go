package registers

const ZERO_FLAG_BYTE_POSITION uint8 = 7
const SUBTRACT_FLAG_BYTE_POSITION uint8 = 6
const HALF_CARRY_FLAG_BYTE_POSITION uint8 = 5
const CARRY_FLAG_BYTE_POSITION uint8 = 4

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
