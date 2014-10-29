package perm

import (
	. "gopkg.in/check.v1"
)

func (s *MySuite) TestSign0(c *C) {
	p, e := NewPerm([]int{})
	c.Assert(e, IsNil)
	c.Check(p.Signature(), DeepEquals, []int{0})
	c.Check(p.Sign(), Equals, 1)
	c.Check(p.Order(), Equals, 1)
}

func (s *MySuite) TestSign1(c *C) {
	p, e := NewPerm([]int{0})
	c.Assert(e, IsNil)
	c.Check(p.Signature(), DeepEquals, []int{0, 1})
	c.Check(p.Sign(), Equals, 1)
	c.Check(p.Order(), Equals, 1)
}

func (s *MySuite) TestSign2(c *C) {
	p, e := NewPerm([]int{1, 0})
	c.Assert(e, IsNil)
	c.Check(p.Signature(), DeepEquals, []int{0, 0, 1})
	c.Check(p.Sign(), Equals, -1)
	c.Check(p.Order(), Equals, 2)
}

func (s *MySuite) TestSign5a(c *C) {
	p, e := NewPerm([]int{1, 0, 3, 2, 4})
	c.Assert(e, IsNil)
	c.Check(p.Signature(), DeepEquals, []int{0, 1, 2, 0, 0, 0})
	c.Check(p.Sign(), Equals, 1)
	c.Check(p.Order(), Equals, 2)
}

func (s *MySuite) TestSign5b(c *C) {
	p, e := NewPerm([]int{0, 1, 3, 4, 2})
	c.Assert(e, IsNil)
	c.Check(p.Signature(), DeepEquals, []int{0, 2, 0, 1, 0, 0})
	c.Check(p.Sign(), Equals, 1)
	c.Check(p.Order(), Equals, 3)
}

func (s *MySuite) TestSign6a(c *C) {
	p, e := NewPerm([]int{1, 2, 3, 4, 5, 0})
	c.Assert(e, IsNil)
	c.Check(p.Signature(), DeepEquals, []int{0, 0, 0, 0, 0, 0, 1})
	c.Check(p.Sign(), Equals, -1)
	c.Check(p.Order(), Equals, 6)
}

func (s *MySuite) TestSign6b(c *C) {
	p, e := NewPerm([]int{0, 2, 1, 4, 5, 3})
	c.Assert(e, IsNil)
	c.Check(p.Signature(), DeepEquals, []int{0, 1, 1, 1, 0, 0, 0})
	c.Check(p.Sign(), Equals, -1)
	c.Check(p.Order(), Equals, 6)
}
