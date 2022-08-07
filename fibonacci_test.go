package mathz

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFibUnder(t *testing.T) {
	cases := []struct {
		number uint64
		want   []uint64
	}{
		{0, []uint64{}},
		{1, []uint64{0}},
		{2, []uint64{0, 1}},
		{3, []uint64{0, 1, 1}},
		{4, []uint64{0, 1, 1, 2}},
		{7, []uint64{0, 1, 1, 2, 3, 5, 8}},
		{10, []uint64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}},
		{30, []uint64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811, 514229}},
	}

	for _, test := range cases {
		got := FibSeries(test.number)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: want %+v, got %+v", test.number, test.want, got)
		}
	}
}

func ExampleFibSeries() {
	fmt.Println(FibSeries(30))
	// Output: [0 1 1 2 3 5 8 13 21 34 55 89 144 233 377 610 987 1597 2584 4181 6765 10946 17711 28657 46368 75025 121393 196418 317811 514229]
}

func BenchmarkFibSeries(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibSeries(93)
	}
}

func TestNthFib(t *testing.T) {
	cases := []struct {
		number uint64
		want   uint64
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{10, 55},
		{25, 75025},
		{50, 12586269025},
		{75, 2111485077978050},
		{93, 12200160415121876738},
	}

	for _, test := range cases {
		got := NthFib(test.number)
		if got != test.want {
			t.Errorf("%d: want %d, got %d", test.number, test.want, got)
		}
	}
}

func ExampleNthFib() {
	fmt.Println(NthFib(10))
	// Output:
	// 55
}

func BenchmarkNthFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NthFib(93)
	}
}
