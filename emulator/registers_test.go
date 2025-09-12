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

func TestGetBC(t *testing.T) {
	r := Registers{}
	r.B = 200 // 11001000
	r.C = 79  // 01001111

	bc := r.GetBC()

	var want uint16 = 51279 // 1100100001001111

	if bc != want {
		t.Errorf(`GetBC() = %d, want match for %d`, bc, want)
	}
}

func TestSetBC(t *testing.T) {
	r := Registers{}
	var bc uint16 = 51279 // 1100100001001111

	r.SetBC(bc)

	var wantB uint8 = 200 // 11001000
	var wantC uint8 = 79  // 01001111

	if r.B != wantB || r.C != wantC {
		t.Errorf(`Result of SetBC(%d): r.B = %d, r.C = %d, want match for r.B = %d, r.C = %d`, bc, r.B, r.C, wantB, wantC)
	}
}

func TestGetDE(t *testing.T) {
	r := Registers{}
	r.D = 200 // 11001000
	r.E = 79  // 01001111

	de := r.GetDE()

	var want uint16 = 51279 // 1100100001001111

	if de != want {
		t.Errorf(`GetDE() = %d, want match for %d`, de, want)
	}
}

func TestSetDE(t *testing.T) {
	r := Registers{}
	var de uint16 = 51279 // 1100100001001111

	r.SetDE(de)

	var wantD uint8 = 200 // 11001000
	var wantE uint8 = 79  // 01001111

	if r.D != wantD || r.E != wantE {
		t.Errorf(`Result of SetDE(%d): r.D = %d, r.E = %d, want match for r.D = %d, r.E = %d`, de, r.D, r.E, wantD, wantE)
	}
}

func TestGetHL(t *testing.T) {
	r := Registers{}
	r.H = 200 // 11001000
	r.L = 79  // 01001111

	hl := r.GetHL()

	var want uint16 = 51279 // 1100100001001111

	if hl != want {
		t.Errorf(`GetHL() = %d, want match for %d`, hl, want)
	}
}

func TestSetHL(t *testing.T) {
	r := Registers{}
	var hl uint16 = 51279 // 1100100001001111

	r.SetHL(hl)

	var wantH uint8 = 200 // 11001000
	var wantL uint8 = 79  // 01001111

	if r.H != wantH || r.L != wantL {
		t.Errorf(`Result of SetHL(%d): r.H = %d, r.L = %d, want match for r.H = %d, r.L = %d`, hl, r.H, r.L, wantH, wantL)
	}
}

func TestGetAF(t *testing.T) {
	r := Registers{}
	r.A = 200            // 11001000
	r.F = FlagsRegister{ // 01100000
		Zero:      false,
		Subtract:  true,
		HalfCarry: true,
		Carry:     false,
	}

	af := r.GetAF()

	var want uint16 = 51296 // 1100100001100000

	if af != want {
		t.Errorf(`GetAF() = %d, want match for %d`, af, want)
	}
}

func TestSetAF(t *testing.T) {
	r := Registers{}
	var af uint16 = 51311 // 1100100001101111

	r.SetAF(af)

	var wantA uint8 = 200   // 11001000
	wantF := FlagsRegister{ // 01100000
		Zero:      false,
		Subtract:  true,
		HalfCarry: true,
		Carry:     false,
	}

	fRegisterInvalid := r.F.Zero != wantF.Zero || r.F.Subtract != wantF.Subtract || r.F.HalfCarry != wantF.HalfCarry || r.F.Carry != wantF.Carry

	if fRegisterInvalid || r.A != wantA {
		t.Errorf(`Result of SetAF(%d): r.A = %d, r.F = %v, want match for r.A = %d, r.F = %v`, af, r.A, r.F, wantA, wantF)
	}
}
