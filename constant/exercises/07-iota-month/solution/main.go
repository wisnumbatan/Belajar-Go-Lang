package main

import "fmt"

func main() {
	const(
		November = 11 - iota
		October
		September
	)

	fmt.Println(September, October, November)
}