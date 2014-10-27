package perm

import (
  "errors"
  "fmt"
  _ "regexp"
  "sort"
)

type Perm struct {
  elements []int
}

func NewPerm(from []int) (*Perm, error) {
  if !validSlice(from) {
    return nil, errors.New("invalid constructing list")
  }
  return &Perm{ elements: copySlice(from) }, nil
}

func ParseCycles(from string) (*Perm, error) {
  // TBD
  return nil, nil
}

func Identity(size int) *Perm {
  elements := make([]int, size)
  for i:=0; i < size; i++ {
    elements[i] = i
  }
  return &Perm{ elements: elements }
}

func (p *Perm) String() string {
  return fmt.Sprintf("%v", p.elements)
}

func (p *Perm) PrintCycles() string {
  // cycles := getCycles()
  // TBD
  return ""
}

func (p *Perm) Size() int {
  return len(p.elements)
}

func (p *Perm) On(i int) int {
  if i >= 0 && i < p.Size() {
    return p.elements[i]
  } else {
    return i
  }
}

func (p *Perm) Inverse() *Perm {
  elements := make([]int, p.Size())
  for i:=0; i<len(elements); i++ {
    elements[p.elements[i]] = i
  }
  return &Perm{ elements: elements }
}

func (p *Perm) Compose(o *Perm) *Perm {
  size := p.Size()
  if o.Size() > size { size = o.Size() }
  elements := make([]int, size)
  for i:=0; i<len(elements); i++ {
    elements[i] = o.On(p.On(i))
  }
  return &Perm{ elements: elements }
}

func (p *Perm) Power(n int) *Perm {
  if n == 0 {
    return Identity(p.Size())
  }
  if n < 0 {
    return p.Inverse().Power(-n)
  }
  elements := make([]int, p.Size())
  for i:=0; i<len(elements); i++ {
    j := i
    for k:=0; k<n; k++ {
      j = p.elements[j]
    }
    elements[i] = j
  }
  return &Perm{ elements: elements }
}

// helpers

func validSlice(from []int) bool {
  check := copySlice(from)
  sort.Ints(check)
  for i:=0; i<len(check); i++ {
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

func (p *Perm) getCycles() [][]int {
  // TBD
  return nil
}
