package perm

import (
	"errors"
	"fmt"
	"sort"
)

const TOP_LEN = 1 << 16

type dot uint16

type Perm struct {
	elements []dot
}

func NewPerm(from []int) (*Perm, error) {
	if len(from) > TOP_LEN {
		return nil, errors.New("constructing list too long")
	}
	if !validSlice(from) {
		return nil, errors.New("invalid constructing list")
	}
	elements := convertSlice(from)
	return &Perm{elements}, nil
}

func Identity(size int) (*Perm, error) {
	if size < 0 || size > TOP_LEN {
		return nil, errors.New("invalid identity size")
	}
	elements := make([]dot, size)
	for i := 0; i < size; i++ {
		elements[i] = dot(i)
	}
	return &Perm{elements}, nil
}

func (p *Perm) String() string {
	return fmt.Sprintf("%v", p.elements)
}

func (p *Perm) Size() int {
	return len(p.elements)
}

func (p *Perm) On(i int) int {
	if i >= 0 && i < len(p.elements) {
		return int(p.elements[i])
	} else {
		return i
	}
}

func (p *Perm) Inverse() *Perm {
	elements := make([]dot, len(p.elements))
	for i := 0; i < len(elements); i++ {
		elements[p.elements[i]] = dot(i)
	}
	return &Perm{elements}
}

func (p *Perm) Compose(q *Perm) *Perm {
	var elements []dot
	psize := dot(len(p.elements))
	qsize := dot(len(q.elements))
	if psize > qsize {
		elements = make([]dot, psize)
	} else {
		elements = make([]dot, qsize)
	}
	for i := 0; i < len(elements); i++ {
		k := dot(i)
		if k < psize {
			k = p.elements[k]
		}
		if k < qsize {
			k = q.elements[k]
		}
		elements[i] = k
	}
	return &Perm{elements}
}

func (p *Perm) Power(n int) *Perm {
	if n == 0 {
		o, _ := Identity(len(p.elements))
		return o
	}
	if n < 0 {
		return p.Inverse().Power(-n)
	}
	elements := make([]dot, len(p.elements))
	for i := 0; i < len(elements); i++ {
		j := dot(i)
		for k := 0; k < n; k++ {
			j = p.elements[j]
		}
		elements[i] = j
	}
	return &Perm{elements}
}

// helpers

func validSlice(from []int) bool {
	check := make([]int, len(from))
	copy(check, from)
	sort.Ints(check)
	for i := 0; i < len(check); i++ {
		if check[i] != i {
			return false
		}
	}
	return true
}

func convertSlice(from []int) []dot {
	to := make([]dot, len(from))
	for i := 0; i < len(from); i++ {
		to[i] = dot(from[i])
	}
	return to
}