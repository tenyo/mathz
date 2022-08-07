package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/tenyo/mathz"
)

type command struct {
	Description string
	Exec        commandFunc
}

type commandFunc func(uint64) string

var cliCommands map[string]command

func cli() {
	if len(os.Args) <= 1 {
		fmt.Println("Missing command")
		help()
	}

	cmdName := os.Args[1]
	cmd, ok := cliCommands[cmdName]
	if !ok {
		fmt.Printf("Unknown command: %s\n", cmdName)
		help()
	}

	if len(os.Args) <= 2 {
		fmt.Printf("Missing parameter to %q command, please specify a number\n\n", cmdName)
		os.Exit(1)
	}

	cmdArg, err := strconv.ParseUint(os.Args[2], 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(cmd.Exec(cmdArg))
}

func init() {
	cliCommands = map[string]command{
		"isprime": {
			Description: "returns true if the given number is prime",
			Exec: func(arg uint64) string {
				return fmt.Sprintf("%t", mathz.IsPrime(arg))
			},
		},
		"nthprime": {
			Description: "returns the N-th prime number",
			Exec: func(arg uint64) string {
				return fmt.Sprintf("%d", mathz.NthPrime(arg))
			},
		},
		"primesunder": {
			Description: "returns a list of all prime numbers less than or equal to the given number",
			Exec: func(arg uint64) string {
				return strings.Trim(fmt.Sprint(mathz.PrimesUnder(arg)), "[]")
			},
		},
		"primefactors": {
			Description: "returns a list of all prime factors of the given number",
			Exec: func(arg uint64) string {
				return fmt.Sprintf("%v", mathz.PrimeFactors(arg))
			},
		},
		"fibseries": {
			Description: "returns a slice of the first n fibonacci numbers",
			Exec: func(arg uint64) string {
				return strings.Trim(fmt.Sprint(mathz.FibSeries(arg)), "[]")
			},
		},
		"nthfib": {
			Description: "returns the N-th fibonacci number",
			Exec: func(arg uint64) string {
				return fmt.Sprintf("%d", mathz.NthFib(arg))
			},
		},
	}
}

func help() {
	list := func(cc map[string]command) string {
		l := []string{}
		for name, c := range cc {
			l = append(l, fmt.Sprintf("%s (%s)", name, c.Description))
		}
		return strings.Join(l, "\n")
	}

	fmt.Printf("Specify one of the following:\n\n%s\n\n", list(cliCommands))
	os.Exit(0)
}
