package main

import (
	"fmt"
	"math"
)

func main() {
	const min int = 1 + 1
	const pi = math.Pi
	const version string = "2.1.0" + "beta"
	const debug = !true
	const A rune = 'A' + 1

	fmt.Println(min, pi, version, debug, A)
}
