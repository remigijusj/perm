package perm

import (
	"testing"
)

func TestNewPerm(t *testing.T) {
	var e error
	_, e = NewPerm([]int{1, 2, 3})
	if e == nil {
		t.Fail()
	}
	_, e = NewPerm([]int{0, 0, 1})
	if e == nil {
		t.Fail()
	}
	_, e = NewPerm([]int{0, 2, 3})
	if e == nil {
		t.Fail()
	}
	_, e = NewPerm([]int{3, 2, 1, 0})
	if e != nil {
		t.Fail()
	}
}

func TestIdentity(t *testing.T) {
	p := Identity(3)
	if p.Size() != 3 || p.String() != "[0 1 2]" {
		t.Fail()
	}
}

func TestString(t *testing.T) {
	p, e := NewPerm([]int{1, 0, 2})
	if e != nil || p.String() != "[1 0 2]" {
		t.Fail()
	}
}

func TestSize(t *testing.T) {
	p, e := NewPerm([]int{1, 3, 2, 0})
	if e != nil || p.Size() != 4 {
		t.Fail()
	}
}

func TestOn(t *testing.T) {
	p, e := NewPerm([]int{1, 4, 2, 0, 3})
	if e != nil || p.On(0) != 1 || p.On(1) != 4 || p.On(2) != 2 || p.On(3) != 0 || p.On(4) != 3 {
		t.Fail()
	}
}

func TestInverse(t *testing.T) {
	p, e := NewPerm([]int{1, 2, 3, 4, 0})
	if e != nil || p.Inverse().String() != "[4 0 1 2 3]" {
		t.Fail()
	}
}

func TestCompose(t *testing.T) {
	p, e1 := NewPerm([]int{1, 2, 0})
	r, e2 := NewPerm([]int{0, 3, 4, 1, 2})
	if e1 != nil || e2 != nil || p.Compose(r).String() != "[3 4 0 1 2]" {
		t.Fail()
	}
}

func TestPower(t *testing.T) {
	p, e := NewPerm([]int{1, 2, 3, 4, 5, 0})
	if e != nil || p.Power(2).String() != "[2 3 4 5 0 1]" {
		t.Fail()
	}
}

func TestParseCycles(t *testing.T) {
	var p *Perm
	var e error
	p, e = ParseCycles("()")
	if e != nil {
		t.Fail()
	}
	p, e = ParseCycles("( )( ( )(")
	if e != nil {
		t.Fail()
	}
	p, e = ParseCycles("(1)")
	if e != nil {
		t.Fail()
	}
	p, e = ParseCycles("(1,2)")
	if e != nil {
		t.Fail()
	}
	p, e = ParseCycles("(1, 2) (3, 4) ")
	if e != nil {
		t.Fail()
	}
	p, e = ParseCycles("(1 2)(3 12)(7 16)")
	if e != nil {
		t.Fail()
	}
	p, e = ParseCycles("(1 2)(3, 8)(7 4)()")
	if e != nil {
		t.Fail()
	}

	p, e = ParseCycles("(1 2 3 0)")
	if e == nil {
		t.Fail()
	}
	p, e = ParseCycles("(1 2)(2, 3)")
	if e == nil {
		t.Fail()
	}
	_ = p
}

func TestPrintCycles(t *testing.T) {
	p, e := NewPerm([]int{})
	if e != nil || p.PrintCycles() != "()" {
		t.Fail()
	}
	p, e = NewPerm([]int{0})
	if e != nil || p.PrintCycles() != "(1)" {
		t.Fail()
	}
	p, e = NewPerm([]int{1, 2, 3, 4, 5, 0})
	if e != nil || p.PrintCycles() != "(1, 2, 3, 4, 5, 6)" {
		t.Fail()
	}
	p, e = NewPerm([]int{1, 2, 0, 4, 5, 3})
	if e != nil || p.PrintCycles() != "(1, 2, 3)(4, 5, 6)" {
		t.Fail()
	}
	p, e = NewPerm([]int{5, 4, 3, 2, 1, 0})
	if e != nil || p.PrintCycles() != "(1, 6)(2, 5)(3, 4)" {
		t.Fail()
	}
}
