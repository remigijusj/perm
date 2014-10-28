package perm

import (
	"testing"
)

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
