package perm

import (
	. "gopkg.in/check.v1"
	"testing"
)

// hook up gocheck into the "go test" runner
func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})

func (s *S) TestNewPerm(c *C) {
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

func (s *S) TestIdentityInvalid(c *C) {
	_, e := Identity(-1)
	c.Check(e, NotNil)

	_, e = Identity(1<<16 + 1)
	c.Check(e, NotNil)
}

func (s *S) TestIdentityValid(c *C) {
	p, e := Identity(3)
	c.Assert(e, IsNil)
	c.Check(p.Size(), Equals, 3)
	c.Check(p.String(), Equals, "[0 1 2]")
}

func (s *S) TestRandom(c *C) {
	p, e := Random(128)
	c.Assert(e, IsNil)
	c.Check(p.Size(), Equals, 128)
}

func (s *S) TestString(c *C) {
	p, e := NewPerm([]int{1, 0, 2})
	c.Assert(e, IsNil)
	c.Check(p.String(), Equals, "[1 0 2]")
}

func (s *S) TestSize(c *C) {
	p, e := NewPerm([]int{1, 3, 2, 0})
	c.Assert(e, IsNil)
	c.Check(p.Size(), Equals, 4)
}

func (s *S) TestOn(c *C) {
	p, e := NewPerm([]int{1, 4, 2, 0, 3})
	c.Assert(e, IsNil)
	c.Check(p.On(0), Equals, 1)
	c.Check(p.On(1), Equals, 4)
	c.Check(p.On(2), Equals, 2)
	c.Check(p.On(3), Equals, 0)
	c.Check(p.On(4), Equals, 3)
	c.Check(p.On(5), Equals, 5)
}

func (s *S) TestInverse(c *C) {
	p, e := NewPerm([]int{1, 2, 3, 4, 0})
	c.Assert(e, IsNil)
	c.Check(p.Inverse().String(), Equals, "[4 0 1 2 3]")
}

func (s *S) TestCompose(c *C) {
	p, e1 := NewPerm([]int{1, 2, 0})
	q, e2 := NewPerm([]int{0, 3, 4, 1, 2})
	c.Assert(e1, IsNil)
	c.Assert(e2, IsNil)
	c.Check(p.Compose(q).String(), Equals, "[3 4 0 1 2]")
}

func (s *S) TestPower(c *C) {
	p, e := NewPerm([]int{1, 2, 3, 4, 5, 0})
	c.Assert(e, IsNil)
	c.Check(p.Power(2).String(), Equals, "[2 3 4 5 0 1]")
}

func (s *S) TestConjugate0(c *C) {
	var p, q *Perm
	var e error

	p, _ = Identity(6)
	q, e = Random(12)
	c.Assert(e, IsNil)
	c.Check(p.Conjugate(q).IsIdentity(), Equals, true)

	p, e = Random(6)
	q, _ = Identity(8)
	c.Assert(e, IsNil)
	c.Check(p.Conjugate(q).IsEqual(p), Equals, true)
}

func (s *S) TestConjugate1(c *C) {
	p, e1 := NewPerm([]int{1, 2, 0})
	q, e2 := NewPerm([]int{0, 3, 4, 1, 2})
	c.Assert(e1, IsNil)
	c.Assert(e2, IsNil)
	c.Check(p.Conjugate(q).String(), Equals, "[3 1 2 4 0]")
}

func (s *S) TestConjugate2(c *C) {
	p, e1 := NewPerm([]int{4, 2, 0, 1, 3})
	q, e2 := NewPerm([]int{1, 2, 0})
	c.Assert(e1, IsNil)
	c.Assert(e2, IsNil)
	c.Check(p.Conjugate(q).String(), Equals, "[1 4 0 2 3]")
}

func (s *S) TestIsIdentity(c *C) {
	var p *Perm
	var e error

	p, e = NewPerm([]int{})
	c.Assert(e, IsNil)
	c.Check(p.IsIdentity(), Equals, true)

	p, e = NewPerm([]int{0, 1})
	c.Assert(e, IsNil)
	c.Check(p.IsIdentity(), Equals, true)

	p, e = NewPerm([]int{0, 1, 2})
	c.Assert(e, IsNil)
	c.Check(p.IsIdentity(), Equals, true)

	p, e = NewPerm([]int{1, 0})
	c.Assert(e, IsNil)
	c.Check(p.IsIdentity(), Equals, false)

	p, e = NewPerm([]int{0, 1, 3, 2})
	c.Assert(e, IsNil)
	c.Check(p.IsIdentity(), Equals, false)
}

func (s *S) TestIsEqual(c *C) {
	var p, q *Perm
	var e, r error

	p, e = NewPerm([]int{})
	c.Assert(e, IsNil)
	c.Check(p.IsEqual(p), Equals, true)

	p, e = NewPerm([]int{0})
	c.Assert(e, IsNil)
	c.Check(p.IsEqual(p), Equals, true)

	p, e = NewPerm([]int{0, 1})
	q, r = NewPerm([]int{0, 1, 2})
	c.Assert(e, IsNil)
	c.Assert(r, IsNil)
	c.Check(p.IsEqual(q), Equals, true)
	c.Check(q.IsEqual(p), Equals, true)
}
