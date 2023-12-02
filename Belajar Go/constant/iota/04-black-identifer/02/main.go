package main

import "fmt"

func main() {
	const (
		EST = -(5 + iota)
		MST
		PST
	)
	fmt.Println(EST, MST, PST)
}