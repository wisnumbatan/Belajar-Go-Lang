package main

import "fmt"

func main() {
	const min = 42

	var i = min

	var f float64 = min
	var b byte = min
	var j int32 = min
	var r rune = min

	b = byte(min)

	fmt.Println(i, f, b, j, r)
}