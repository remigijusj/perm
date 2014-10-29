package perm

// TODO: uint16? caching in Perm?
func (p *Perm) Signature() []int {
	size := len(p.elements)
	sign := make([]int, size+1)

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
		// trace a cycle
		cnt := 0
		for j := dot(m); !marks[j]; j = p.elements[j] {
			marks[j] = true
			cnt++
		}
		sign[cnt]++
	}
	return sign
}

func (p *Perm) Sign() int {
	sgn := p.Signature()
	sum := 0
	for i := 2; i < len(sgn); i += 2 {
		sum += sgn[i]
	}
	if sum%2 == 0 {
		return 1
	} else {
		return -1
	}
}

// TODO: binary reduce?
func (p *Perm) Order() int {
	if p.Size() < 2 {
		return 1
	}
	sgn := p.Signature()
	ord := 1
	for i, v := range sgn {
		if v > 0 && i >= 2 {
			ord = lcm(ord, i)
		}
	}
	return ord
}

// TODO: int64?
func lcm(a, b int) int {
	return a * (b / gcd(a, b))
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
