package perm

import (
	. "gopkg.in/check.v1"
)

func (s *S) TestSignature0(c *C) {
	p, e := NewPerm([]int{})
	c.Assert(e, IsNil)
	c.Check(p.Signature(), DeepEquals, []int{0})
	c.Check(p.Sign(), Equals, 1)
	c.Check(p.Order(), Equals, 1)
}

func (s *S) TestSignature1(c *C) {
	p, e := NewPerm([]int{0})
	c.Assert(e, IsNil)
	c.Check(p.Signature(), DeepEquals, []int{0, 1})
	c.Check(p.Sign(), Equals, 1)
	c.Check(p.Order(), Equals, 1)
}

func (s *S) TestSignature2(c *C) {
	p, e := NewPerm([]int{1, 0})
	c.Assert(e, IsNil)
	c.Check(p.Signature(), DeepEquals, []int{0, 0, 1})
	c.Check(p.Sign(), Equals, -1)
	c.Check(p.Order(), Equals, 2)
	c.Check(p.OrderToCycle(2), Equals, 1)
}

func (s *S) TestSignature4(c *C) {
	p, e := NewPerm([]int{1, 0, 3, 2})
	c.Assert(e, IsNil)
	c.Check(p.Signature(), DeepEquals, []int{0, 0, 2, 0, 0})
	c.Check(p.Sign(), Equals, 1)
	c.Check(p.Order(), Equals, 2)
	c.Check(p.OrderToCycle(2), Equals, -1)
}

func (s *S) TestSignature5a(c *C) {
	p, e := NewPerm([]int{1, 0, 3, 2, 4})
	c.Assert(e, IsNil)
	c.Check(p.Signature(), DeepEquals, []int{0, 1, 2, 0, 0, 0})
	c.Check(p.Sign(), Equals, 1)
	c.Check(p.Order(), Equals, 2)
}

func (s *S) TestSignature5c(c *C) {
	p, e := NewPerm([]int{1, 0, 3, 4, 2})
	c.Assert(e, IsNil)
	c.Check(p.Signature(), DeepEquals, []int{0, 0, 1, 1, 0, 0})
	c.Check(p.Sign(), Equals, -1)
	c.Check(p.Order(), Equals, 6)
	c.Check(p.OrderToCycle(2), Equals, 3)
	c.Check(p.OrderToCycle(3), Equals, 2)
}

func (s *S) TestSignature5d(c *C) {
	p, e := NewPerm([]int{0, 1, 3, 4, 2})
	c.Assert(e, IsNil)
	c.Check(p.Signature(), DeepEquals, []int{0, 2, 0, 1, 0, 0})
	c.Check(p.Sign(), Equals, 1)
	c.Check(p.Order(), Equals, 3)
	c.Check(p.OrderToCycle(3), Equals, 1)
}

func (s *S) TestSignature6a(c *C) {
	p, e := NewPerm([]int{1, 2, 3, 4, 5, 0})
	c.Assert(e, IsNil)
	c.Check(p.Signature(), DeepEquals, []int{0, 0, 0, 0, 0, 0, 1})
	c.Check(p.Sign(), Equals, -1)
	c.Check(p.Order(), Equals, 6)
}

func (s *S) TestSignature6b(c *C) {
	p, e := NewPerm([]int{0, 2, 1, 4, 5, 3})
	c.Assert(e, IsNil)
	c.Check(p.Signature(), DeepEquals, []int{0, 1, 1, 1, 0, 0, 0})
	c.Check(p.Sign(), Equals, -1)
	c.Check(p.Order(), Equals, 6)
}

func (s *S) TestSignature6c(c *C) {
	p, e := NewPerm([]int{5, 4, 1, 2, 3, 0})
	c.Assert(e, IsNil)
	c.Check(p.Signature(), DeepEquals, []int{0, 0, 1, 0, 1, 0, 0})
	c.Check(p.Sign(), Equals, 1)
	c.Check(p.Order(), Equals, 4)
	c.Check(p.OrderToCycle(4), Equals, -1)
}
