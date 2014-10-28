package perm

import (
	. "gopkg.in/check.v1"
)

func (s *MySuite) TestParseCycles(c *C) {
	var e error

	_, e = ParseCycles("()")
	c.Assert(e, IsNil)

	_, e = ParseCycles("( )( ( )(")
	c.Assert(e, IsNil)

	_, e = ParseCycles("(1)")
	c.Assert(e, IsNil)

	_, e = ParseCycles("(1,2)")
	c.Assert(e, IsNil)

	_, e = ParseCycles("(1, 2) (3, 4) ")
	c.Assert(e, IsNil)

	_, e = ParseCycles("(1 2)(3 12)(7 16)")
	c.Assert(e, IsNil)

	_, e = ParseCycles("(1 2)(3, 8)(7 4)()")
	c.Assert(e, IsNil)

	_, e = ParseCycles("(1 2 3 0)")
	c.Assert(e, NotNil)

	_, e = ParseCycles("(1 2)(2, 3)")
	c.Assert(e, NotNil)
}

func (s *MySuite) TestPrintCycles(c *C) {
	p, e := NewPerm([]int{})
	c.Assert(e, IsNil)
	c.Check(p.PrintCycles(), Equals, "()")

	p, e = NewPerm([]int{0})
	c.Assert(e, IsNil)
	c.Check(p.PrintCycles(), Equals, "(1)")

	p, e = NewPerm([]int{1, 2, 3, 4, 5, 0})
	c.Assert(e, IsNil)
	c.Check(p.PrintCycles(), Equals, "(1, 2, 3, 4, 5, 6)")

	p, e = NewPerm([]int{1, 2, 0, 4, 5, 3})
	c.Assert(e, IsNil)
	c.Check(p.PrintCycles(), Equals, "(1, 2, 3)(4, 5, 6)")

	p, e = NewPerm([]int{5, 4, 3, 2, 1, 0})
	c.Assert(e, IsNil)
	c.Check(p.PrintCycles(), Equals, "(1, 6)(2, 5)(3, 4)")
}
