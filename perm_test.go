package perm

import (
	. "gopkg.in/check.v1"
	"testing"
)

// hook up gocheck into the "go test" runner
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestNewPerm(c *C) {
	var e error

	_, e = NewPerm([]int{1, 2, 3})
	c.Check(e, NotNil)

	_, e = NewPerm([]int{0, 0, 1})
	c.Check(e, NotNil)

	_, e = NewPerm([]int{-2, 0, 1})
	c.Check(e, NotNil)

	_, e = NewPerm([]int{0, 2, 3})
	c.Check(e, NotNil)

	_, e = NewPerm([]int{3, 2, 1, 0})
	c.Check(e, IsNil)
}

func (s *MySuite) TestIdentity(c *C) {
	p, e := Identity(-1)
	c.Check(e, NotNil)

	p, e = Identity(1<<16 + 1)
	c.Check(e, NotNil)

	p, e = Identity(3)
	c.Check(e, IsNil)
	c.Check(p.Size(), Equals, 3)
	c.Check(p.String(), Equals, "[0 1 2]")
}

func (s *MySuite) TestString(c *C) {
	p, e := NewPerm([]int{1, 0, 2})
	c.Assert(e, IsNil)
	c.Check(p.String(), Equals, "[1 0 2]")
}

func (s *MySuite) TestSize(c *C) {
	p, e := NewPerm([]int{1, 3, 2, 0})
	c.Assert(e, IsNil)
	c.Check(p.Size(), Equals, 4)
}

func (s *MySuite) TestOn(c *C) {
	p, e := NewPerm([]int{1, 4, 2, 0, 3})
	c.Assert(e, IsNil)
	c.Check(p.On(0), Equals, 1)
	c.Check(p.On(1), Equals, 4)
	c.Check(p.On(2), Equals, 2)
	c.Check(p.On(3), Equals, 0)
	c.Check(p.On(4), Equals, 3)
	c.Check(p.On(5), Equals, 5)
}

func (s *MySuite) TestInverse(c *C) {
	p, e := NewPerm([]int{1, 2, 3, 4, 0})
	c.Assert(e, IsNil)
	c.Check(p.Inverse().String(), Equals, "[4 0 1 2 3]")
}

func (s *MySuite) TestCompose(c *C) {
	p, e1 := NewPerm([]int{1, 2, 0})
	r, e2 := NewPerm([]int{0, 3, 4, 1, 2})
	c.Assert(e1, IsNil)
	c.Assert(e2, IsNil)
	c.Check(p.Compose(r).String(), Equals, "[3 4 0 1 2]")
}

func (s *MySuite) TestPower(c *C) {
	p, e := NewPerm([]int{1, 2, 3, 4, 5, 0})
	c.Assert(e, IsNil)
	c.Check(p.Power(2).String(), Equals, "[2 3 4 5 0 1]")
}

func (s *MySuite) TestSignature(c *C) {
	p, e := NewPerm([]int{})
	c.Assert(e, IsNil)
	c.Check(p.Signature(), DeepEquals, []int{0})

	p, e = NewPerm([]int{0})
	c.Assert(e, IsNil)
	c.Check(p.Signature(), DeepEquals, []int{0, 1})

	p, e = NewPerm([]int{1, 0})
	c.Assert(e, IsNil)
	c.Check(p.Signature(), DeepEquals, []int{0, 0, 1})

	p, e = NewPerm([]int{1, 0, 3, 2, 4})
	c.Assert(e, IsNil)
	c.Check(p.Signature(), DeepEquals, []int{0, 1, 2, 0, 0, 0})

	p, e = NewPerm([]int{1, 0, 3, 2, 4})
	c.Assert(e, IsNil)
	c.Check(p.Signature(), DeepEquals, []int{0, 1, 2, 0, 0, 0})

	p, e = NewPerm([]int{1, 2, 3, 4, 5, 0})
	c.Assert(e, IsNil)
	c.Check(p.Signature(), DeepEquals, []int{0, 0, 0, 0, 0, 0, 1})
}
