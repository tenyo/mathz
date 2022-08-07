# mathz

A Go package with some math-related functions, currently related to prime and fibonacci numbers

## Benchmarks

```
go test ./... -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/tenyo/mathz
cpu: VirtualApple @ 2.50GHz
BenchmarkFibSeries-8      	 6299250	       189.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkNthFib-8         	37383807	        34.40 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsPrime-8        	 8014004	       153.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkNthPrime-8       	   13479	     88863 ns/op	       0 B/op	       0 allocs/op
BenchmarkPrimesUnder-8    	    9570	    118072 ns/op	   10240 B/op	       1 allocs/op
BenchmarkPrimeFactors-8   	    3952	    300962 ns/op	   19080 B/op	       2 allocs/op
PASS
ok  	github.com/tenyo/mathz	8.810s
?   	github.com/tenyo/mathz/cmd	[no test files]
```

## CLI

See `cmd/` folder for a simple CLI that calls the functions in this package
