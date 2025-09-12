package emulator

import (
	"testing"
)

func TestGetByteFromFlags(t *testing.T) {
	// Test for 11110000
	fr := FlagsRegister{
		Zero:      true,
		Subtract:  true,
		HalfCarry: true,
		Carry:     true,
	}

	b := fr.GetByteFromFlags()
	var wantByte uint8 = 240

	if b != wantByte {
		t.Errorf(`GetByteFromFlags() = %v, want match for %v`, b, wantByte)
	}

	// Test for 00000000
	fr = FlagsRegister{
		Zero:      false,
		Subtract:  false,
		HalfCarry: false,
		Carry:     false,
	}

	b = fr.GetByteFromFlags()
	wantByte = 0

	if b != wantByte {
		t.Errorf(`GetByteFromFlags() = %v, want match for %v`, b, wantByte)
	}

	// Test for 01110000
	fr = FlagsRegister{
		Zero:      false,
		Subtract:  true,
		HalfCarry: true,
		Carry:     true,
	}

	b = fr.GetByteFromFlags()
	wantByte = 112

	if b != wantByte {
		t.Errorf(`GetByteFromFlags() = %v, want match for %v`, b, wantByte)
	}

	// Test for 00110000
	fr = FlagsRegister{
		Zero:      false,
		Subtract:  false,
		HalfCarry: true,
		Carry:     true,
	}

	b = fr.GetByteFromFlags()
	wantByte = 48

	if b != wantByte {
		t.Errorf(`GetByteFromFlags() = %v, want match for %v`, b, wantByte)
	}

	// Test for 00010000
	fr = FlagsRegister{
		Zero:      false,
		Subtract:  false,
		HalfCarry: false,
		Carry:     true,
	}

	b = fr.GetByteFromFlags()
	wantByte = 16

	if b != wantByte {
		t.Errorf(`GetByteFromFlags() = %v, want match for %v`, b, wantByte)
	}

	// Test for 10010000
	fr = FlagsRegister{
		Zero:      true,
		Subtract:  false,
		HalfCarry: false,
		Carry:     true,
	}

	b = fr.GetByteFromFlags()
	wantByte = 144

	if b != wantByte {
		t.Errorf(`GetByteFromFlags() = %v, want match for %v`, b, wantByte)
	}
}

func TestSetFlagsFromByte(t *testing.T) {
	// Test for 11110000
	fr := FlagsRegister{}
	var b uint8 = 240

	fr.SetFlagsFromByte(b)

	want := FlagsRegister{
		Zero:      true,
		Subtract:  true,
		HalfCarry: true,
		Carry:     true,
	}

	if fr.Zero != want.Zero || fr.Subtract != want.Subtract || fr.HalfCarry != want.HalfCarry || fr.Carry != want.Carry {
		t.Errorf(`SetFlagsFromByte(%d) = %v, want match for %v`, b, fr, want)
	}

	// Test for 00000000
	fr = FlagsRegister{
		Zero:      true,
		Subtract:  true,
		HalfCarry: true,
		Carry:     true,
	}
	b = 0

	fr.SetFlagsFromByte(b)

	want = FlagsRegister{
		Zero:      false,
		Subtract:  false,
		HalfCarry: false,
		Carry:     false,
	}

	if fr.Zero != want.Zero || fr.Subtract != want.Subtract || fr.HalfCarry != want.HalfCarry || fr.Carry != want.Carry {
		t.Errorf(`SetFlagsFromByte(%d) = %v, want match for %v`, b, fr, want)
	}

	// Test for 01110000
	fr = FlagsRegister{}
	b = 112

	fr.SetFlagsFromByte(b)

	want = FlagsRegister{
		Zero:      false,
		Subtract:  true,
		HalfCarry: true,
		Carry:     true,
	}

	if fr.Zero != want.Zero || fr.Subtract != want.Subtract || fr.HalfCarry != want.HalfCarry || fr.Carry != want.Carry {
		t.Errorf(`SetFlagsFromByte(%d) = %v, want match for %v`, b, fr, want)
	}

	// Test for 00110000
	fr = FlagsRegister{}
	b = 48

	fr.SetFlagsFromByte(b)

	want = FlagsRegister{
		Zero:      false,
		Subtract:  false,
		HalfCarry: true,
		Carry:     true,
	}

	if fr.Zero != want.Zero || fr.Subtract != want.Subtract || fr.HalfCarry != want.HalfCarry || fr.Carry != want.Carry {
		t.Errorf(`SetFlagsFromByte(%d) = %v, want match for %v`, b, fr, want)
	}

	// Test for 00010000
	fr = FlagsRegister{}
	b = 16

	fr.SetFlagsFromByte(b)

	want = FlagsRegister{
		Zero:      false,
		Subtract:  false,
		HalfCarry: false,
		Carry:     true,
	}

	if fr.Zero != want.Zero || fr.Subtract != want.Subtract || fr.HalfCarry != want.HalfCarry || fr.Carry != want.Carry {
		t.Errorf(`SetFlagsFromByte(%d) = %v, want match for %v`, b, fr, want)
	}

	// Test for 10010000
	fr = FlagsRegister{}
	b = 144

	fr.SetFlagsFromByte(b)

	want = FlagsRegister{
		Zero:      true,
		Subtract:  false,
		HalfCarry: false,
		Carry:     true,
	}

	if fr.Zero != want.Zero || fr.Subtract != want.Subtract || fr.HalfCarry != want.HalfCarry || fr.Carry != want.Carry {
		t.Errorf(`SetFlagsFromByte(%d) = %v, want match for %v`, b, fr, want)
	}
}
