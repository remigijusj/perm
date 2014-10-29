package perm

import (
	. "gopkg.in/check.v1"
)

func (s *MySuite) TestParseCyclesInvalid(c *C) {
	var e error

	_, e = ParseCycles("(1 2 3 0)")
	c.Check(e, NotNil)

	_, e = ParseCycles("(1 2)(2, 3)")
	c.Check(e, NotNil)

	_, e = ParseCycles("(1, 2, 65537)")
	c.Check(e, NotNil)
}

func (s *MySuite) TestParseCyclesValid(c *C) {
	var p *Perm
	var e error

	p, e = ParseCycles("()")
	c.Assert(e, IsNil)
	c.Check(p.String(), Equals, "[]")

	p, e = ParseCycles("( )( ( )(")
	c.Assert(e, IsNil)
	c.Check(p.String(), Equals, "[]")

	p, e = ParseCycles("(1)")
	c.Assert(e, IsNil)
	c.Check(p.String(), Equals, "[0]")

	p, e = ParseCycles("(1,2)")
	c.Assert(e, IsNil)
	c.Check(p.String(), Equals, "[1 0]")

	p, e = ParseCycles("(3,5)")
	c.Assert(e, IsNil)
	c.Check(p.String(), Equals, "[0 1 4 3 2]")

	p, e = ParseCycles("(1, 2) (3, 4) ")
	c.Assert(e, IsNil)
	c.Check(p.String(), Equals, "[1 0 3 2]")

	p, e = ParseCycles("(1 2)(3 12)(7 16)")
	c.Assert(e, IsNil)
	c.Check(p.String(), Equals, "[1 0 11 3 4 5 15 7 8 9 10 2 12 13 14 6]")

	p, e = ParseCycles("(1 2)(3, 8)(7 4)()")
	c.Check(e, IsNil)
	c.Check(p.String(), Equals, "[1 0 7 6 4 5 3 2]")
}

func (s *MySuite) TestPrintCycles(c *C) {
	p, e := NewPerm([]int{})
	c.Assert(e, IsNil)
	c.Check(p.PrintCycles(), Equals, "()")

	p, e = NewPerm([]int{0})
	c.Assert(e, IsNil)
	c.Check(p.PrintCycles(), Equals, "()")

	p, e = NewPerm([]int{1, 2, 3, 4, 5, 0})
	c.Assert(e, IsNil)
	c.Check(p.PrintCycles(), Equals, "(1, 2, 3, 4, 5, 6)")

	p, e = NewPerm([]int{1, 2, 0, 4, 5, 3})
	c.Assert(e, IsNil)
	c.Check(p.PrintCycles(), Equals, "(1, 2, 3)(4, 5, 6)")

	p, e = NewPerm([]int{5, 4, 3, 2, 1, 0})
	c.Assert(e, IsNil)
	c.Check(p.PrintCycles(), Equals, "(1, 6)(2, 5)(3, 4)")

	p, e = NewPerm([]int{0, 1, 4, 3, 2, 5})
	c.Assert(e, IsNil)
	c.Check(p.PrintCycles(), Equals, "(3, 5)")
}
