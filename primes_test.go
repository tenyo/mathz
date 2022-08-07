package mathz

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsPrime(t *testing.T) {
	cases := []struct {
		number uint64
		want   bool
	}{
		{0, false},
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{29, true},
		{91, false},
		{97, true},
		{911, true},
		{1000, false},
	}

	for _, test := range cases {
		got := IsPrime(test.number)
		if got != test.want {
			t.Errorf("%d: want %t, got %t", test.number, test.want, got)
		}
	}
}

func ExampleIsPrime() {
	fmt.Println(IsPrime(13))
	fmt.Println(IsPrime(100))
	// Output:
	// true
	// false
}

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(99991)
		IsPrime(100000)
		IsPrime(987654321)
	}
}

func TestNthPrime(t *testing.T) {
	cases := []struct {
		number uint64
		want   uint64
	}{
		{0, 0},
		{1, 2},
		{2, 3},
		{3, 5},
		{4, 7},
		{5, 11},
		{10, 29},
		{25, 97},
		{50, 229},
		{100, 541},
		{1000, 7919},
		{10000, 104729},
	}

	for _, test := range cases {
		got := NthPrime(test.number)
		if got != test.want {
			t.Errorf("%d: want %d, got %d", test.number, test.want, got)
		}
	}
}

func ExampleNthPrime() {
	fmt.Println(NthPrime(10))
	// Output:
	// 29
}

func BenchmarkNthPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NthPrime(1000)
	}
}

func TestPrimesUnder(t *testing.T) {
	cases := []struct {
		number uint64
		want   []uint64
	}{
		{0, []uint64{}},
		{1, []uint64{}},
		{2, []uint64{2}},
		{3, []uint64{2, 3}},
		{4, []uint64{2, 3}},
		{7, []uint64{2, 3, 5, 7}},
		{10, []uint64{2, 3, 5, 7}},
		{100, []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}},
	}

	for _, test := range cases {
		got := PrimesUnder(test.number)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: want %+v, got %+v", test.number, test.want, got)
		}
	}
}

func ExamplePrimesUnder() {
	fmt.Println(PrimesUnder(50))
	// Output: [2 3 5 7 11 13 17 19 23 29 31 37 41 43 47]
}

func BenchmarkPrimesUnder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimesUnder(10000)
	}
}

func TestPrimeFactors(t *testing.T) {
	cases := []struct {
		number uint64
		want   []uint64
	}{
		{0, []uint64{0}},
		{1, []uint64{1}},
		{2, []uint64{2}},
		{3, []uint64{3}},
		{4, []uint64{2, 2}},
		{5, []uint64{5}},
		{8, []uint64{2, 2, 2}},
		{9, []uint64{3, 3}},
		{10, []uint64{2, 5}},
		{11, []uint64{11}},
		{100, []uint64{2, 2, 5, 5}},
		{911, []uint64{911}},
		{17711, []uint64{89, 199}},
		{832040, []uint64{2, 2, 2, 5, 11, 31, 61}},
		{1234567, []uint64{127, 9721}},
		{12345678, []uint64{2, 3, 3, 47, 14593}},
		{1234567891, []uint64{1234567891}},
	}

	for _, test := range cases {
		got := PrimeFactors(test.number)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: want %+v, got %+v", test.number, test.want, got)
		}
	}
}

func ExamplePrimeFactors() {
	fmt.Println(PrimeFactors(100))
	fmt.Println(PrimeFactors(911))
	// Output:
	// [2 2 5 5]
	// [911]
}

func BenchmarkPrimeFactors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimeFactors(433494437)
	}
}
