package perm

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
)

// cycle representation

func ParseCycles(from string) (*Perm, error) {
	parts, max, err := scanCycleRep(from)
	if err != nil {
		return nil, err
	}
	perm, err := buildPermFromCycleRep(parts, max)
	return perm, err
}

// scan integers, liberally
// ex: (1 2)(3, 8)(7 4)() -> []int{-1, 0, 1, -1, 2, 7, -1, 6, 3, -1}
func scanCycleRep(from string) ([]int, int, error) {
	rx := regexp.MustCompile(`\d+|[()]+`)
	items := rx.FindAllString(from, -1)
	parts := []int{}
	max := -1
	for _, item := range items {
		part := -1
		fmt.Sscanf(item, "%d", &part)
		if part > 0 {
			part--
		} else if part == 0 {
			return nil, 0, errors.New("integers can't be zero")
		}
		parts = append(parts, part)
		if part > max {
			max = part
		}
	}
	return parts, max, nil
}

// build permutation
// ex: []int{-1, 0, 1, -1, 2, 7, -1, 6, 3, -1} -> []dot{1, 0, 7, 6, 4, 5, 3, 2}
func buildPermFromCycleRep(parts []int, max int) (*Perm, error) {
	perm, err := Identity(max + 1)
	if err != nil {
		return nil, err
	}
	first, point := -1, -1
	for _, part := range parts {
		if part == -1 {
			if first >= 0 && point >= 0 {
				if int(perm.elements[point]) != point {
					return nil, errors.New("integers must be unique")
				}
				perm.elements[point] = dot(first)
			}
			first, point = -1, -1
		} else {
			if point == -1 {
				first, point = part, part
			} else {
				if int(perm.elements[point]) != point {
					return nil, errors.New("integers must be unique")
				}
				perm.elements[point] = dot(part)
				point = part
			}
		}
	}
	// fmt.Printf("%s -> %#v, %#v\n", from, parts, perm.elements)
	return perm, nil
}

func (p *Perm) PrintCycles() string {
	cycles := p.getCycles()
	var buf bytes.Buffer
	for _, cycle := range cycles {
		buf.WriteString("(")
		for idx, i := range cycle {
			if idx > 0 {
				buf.WriteString(", ")
			}
			fmt.Fprintf(&buf, "%d", i+1)
		}
		buf.WriteString(")")
	}
	return buf.String()
}

// TODO: publish?
func (p *Perm) getCycles() [][]dot {
	size := len(p.elements)
	cycles := make([][]dot, 0, 1)
	marks := make([]bool, size)
	m := 0
	for {
		// find next unmarked
		for m < size && marks[m] {
			m++
		}
		if m == size {
			break
		}
		// construct a cycle
		cycle := make([]dot, 0, 1)
		for j := dot(m); !marks[j]; j = p.elements[j] {
			marks[j] = true
			cycle = append(cycle, j)
		}
		cycles = append(cycles, cycle)
	}
	// exceptional case: empty
	if len(cycles) == 0 {
		cycles = append(cycles, []dot{})
	}
	return cycles
}
