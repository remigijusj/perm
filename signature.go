package perm

// TODO: caching in Perm?
func (p Perm) Signature() []int {
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

func (p Perm) Sign() int {
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

// TODO: binary reduce? multi-lcm algorithm?
// TODO: control overflow
func (p Perm) Order() int {
	if len(p.elements) < 2 {
		return 1
	}
	sgn := p.Signature()
	ord := 1
	for i, v := range sgn {
		if i >= 2 && v > 0 {
			ord = lcm(ord, i)
		}
	}
	return ord
}

func (p Perm) OrderToCycle(n int) int {
	if n < 2 {
		return -1
	}
	sgn := p.Signature()
	// there must be unique n-cycle
	if sgn[n] != 1 {
		return -1
	}
	pow := 1
	for i, v := range sgn {
		if gcd(i, n) > 1 {
			// no cycles which could reduce to n
			if i != n && v > 0 {
				return -1
			}
		} else {
			// contributes to power
			if i >= 2 && v > 0 {
				pow = lcm(pow, i)
			}
		}
	}
	return pow
}

func lcm(a, b int) int {
	return a * (b / gcd(a, b))
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
