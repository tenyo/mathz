package mathz

// FibSeries returns a slice of the first n fibonacci numbers
// uint64 will overflow for n>94
func FibSeries(n uint64) []uint64 {
	switch n {
	case 0:
		return []uint64{}
	case 1:
		return []uint64{0}
	}

	fibs := make([]uint64, 2, 94)
	copy(fibs, []uint64{0, 1})

	for i := uint64(2); i < n; i++ {
		f := fibs[i-2] + fibs[i-1]
		fibs = append(fibs, f)
	}

	return fibs
}

// NthFib returns the Nth fibonacci number
// uint64 will overflow for n>93
func NthFib(n uint64) uint64 {
	if n < 2 {
		return n
	}

	var f2, f1, f uint64 = 0, 1, 1

	for i := uint64(2); i <= n; i++ {
		f = f2 + f1
		f2, f1 = f1, f
	}

	return f
}

// 𝑓𝑛=(((5‾√+1)/2)𝑛−((1−5‾√)/2)𝑛)/5‾√
