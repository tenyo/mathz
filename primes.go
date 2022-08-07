package mathz

import (
	"math"
)

// IsPrime returns true if the given number is prime, false otherwise
func IsPrime(num uint64) bool {
	switch {
	case num < 2:
		return false
	case num == 2:
		return true
	case num%2 == 0:
		return false
	}

	for i := uint64(3); i <= uint64(math.Sqrt(float64(num))); i += 2 {
		if num%i == 0 {
			return false
		}
	}

	return true
}

// NthPrime returns the Nth prime number
func NthPrime(n uint64) uint64 {
	switch n {
	case 0:
		return 0
	case 1:
		return 2
	}

	var prime uint64

	for ctr, p := uint64(2), uint64(3); ctr <= n; p += 2 {
		if IsPrime(p) {
			prime = p
			ctr++
		}
	}

	return prime
}

// PrimesUnder returns a slice of all prime numbers less than or equal to num
func PrimesUnder(num uint64) []uint64 {
	if num < 2 {
		return []uint64{}
	}

	if num == 2 {
		return []uint64{2}
	}

	// using Legendre's formula for estimating the number of primes
	// π(x) ~ x/(ln(x) − 1.08366)
	pnt := func(x float64) int {
		return int(math.Ceil(x / (math.Log(x) - 1.08366)))
	}

	// this lets us preallocate the slice and avoid costly appends
	primes := make([]uint64, 1, pnt(float64(num)))
	primes[0] = 2

outerloop:
	for i := uint64(3); i <= num; i += 2 {
		for j := uint64(3); j <= uint64(math.Sqrt(float64(i))); j += 2 {
			if i%j == 0 {
				continue outerloop
			}
		}
		primes = append(primes, i)
	}

	return primes
}

// PrimeFactors returns a slice of all prime factors of a given number
func PrimeFactors(num uint64) []uint64 {
	var factors []uint64

	primes := PrimesUnder(uint64(math.Sqrt(float64(num))))

	for _, p := range primes {
		for num%p == 0 && num > p {
			factors = append(factors, p)
			num /= p
		}
	}

	return append(factors, num)
}
