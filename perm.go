package perm

import (
	"errors"
	"fmt"
	"sort"
)

type Perm struct {
	elements []int
}

func NewPerm(from []int) (*Perm, error) {
	if !validSlice(from) {
		return nil, errors.New("invalid constructing list")
	}
	return &Perm{copySlice(from)}, nil
}

func Identity(size int) *Perm {
	elements := make([]int, size)
	for i := 0; i < size; i++ {
		elements[i] = i
	}
	return &Perm{elements}
}

func (p *Perm) String() string {
	return fmt.Sprintf("%v", p.elements)
}

func (p *Perm) Size() int {
	return len(p.elements)
}

func (p *Perm) On(i int) int {
	if i >= 0 && i < len(p.elements) {
		return p.elements[i]
	} else {
		return i
	}
}

func (p *Perm) Inverse() *Perm {
	elements := make([]int, len(p.elements))
	for i := 0; i < len(elements); i++ {
		elements[p.elements[i]] = i
	}
	return &Perm{elements}
}

func (p *Perm) Compose(o *Perm) *Perm {
	size := len(p.elements)
	osize := len(o.elements)
	if osize > size {
		size = osize
	}
	elements := make([]int, size)
	for i := 0; i < len(elements); i++ {
		elements[i] = o.On(p.On(i))
	}
	return &Perm{elements}
}

func (p *Perm) Power(n int) *Perm {
	if n == 0 {
		return Identity(len(p.elements))
	}
	if n < 0 {
		return p.Inverse().Power(-n)
	}
	elements := make([]int, len(p.elements))
	for i := 0; i < len(elements); i++ {
		j := i
		for k := 0; k < n; k++ {
			j = p.elements[j]
		}
		elements[i] = j
	}
	return &Perm{elements}
}

func (p *Perm) Signature() []int {
	size := len(p.elements)
	sign := make([]int, size+1)

	marks := make([]bool, size)
	for {
		// find first unmarked
		m := -1
		for i := 0; i < size; i++ {
			if !marks[i] {
				m = i
				break
			}
		}
		if m == -1 {
			break
		}
		// trace a cycle
		cnt := 0
		for j := m; !marks[j]; j = p.elements[j] {
			marks[j] = true
			cnt++
		}
		sign[cnt]++
	}
	return sign
}

// general helpers

func validSlice(from []int) bool {
	check := copySlice(from)
	sort.Ints(check)
	for i := 0; i < len(check); i++ {
		if check[i] != i {
			return false
		}
	}
	return true
}

func copySlice(from []int) []int {
	to := make([]int, len(from))
	copy(to, from)
	return to
}
