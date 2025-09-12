package registers

// The "f" register is a special register called the "flags" register.
// The lower four bits of the register are always 0s and the CPU
// automatically writes to the upper four bits when certain things happen.
// In other words, the CPU "flags" certain states.
//    ┌-> Carry
//  ┌-+> Subtraction
//  | |
// 1111 0000
// | |
// └-+> Zero
//   └-> Half Carry
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
		zeroByte = 1 << 7
	} else {
		zeroByte = 0
	}

	if f.Subtract {
		subtractByte = 1 << 6
	} else {
		subtractByte = 0
	}

	if f.HalfCarry {
		halfCarryByte = 1 << 5
	} else {
		halfCarryByte = 0
	}

	if f.Carry {
		carryByte = 1 << 4
	} else {
		carryByte = 0
	}

	return zeroByte | subtractByte | halfCarryByte | carryByte
}

func (f *FlagsRegister) SetFlagsFromByte(b uint8) {
	// TODO
}
