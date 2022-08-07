# Mathz CLI

A simple CLI for executing the fuctions from the mathz package

## Build

```
go build -o mathz .
```

## Use

```
./mathz
Missing command
Specify one of the following:

isprime (returns true if the given number is prime)
nthprime (returns the N-th prime number)
primesunder (returns a list of all prime numbers less than or equal to the given number)
primefactors (returns a list of all prime factors of the given number)
fibseries (returns a slice of the first n fibonacci numbers)
nthfib (returns the N-th fibonacci number)
```

```
./mathz primesunder 100
2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97
```

```
./mathz nthfib 80
23416728348467685
```
