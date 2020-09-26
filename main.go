package main

import (
	"flag"
	"fmt"
	"rl-algo/rl"
)

// Usage:
// ./rl-algo -test="simple-bandit"
func main() {
	test := flag.String("test", "", "The test to launch (eg. simple-bandit)")

	flag.Parse()

	// Print special information for library users
	fmt.Println("")
	fmt.Println("COPYRIGHT NOTICE")
	fmt.Println("================")
	fmt.Println("This library contains my implementations of the 'Reinforcement Learning - An Introduction' algorithms. It's available under a MIT license.")
	fmt.Println("")
	fmt.Println("© 2020 Cyril Dever. All rights reserved.")
	fmt.Println("")

	switch *test {
	case "simple-bandit":
		rl.TestSimpleBandit()
	}
}
