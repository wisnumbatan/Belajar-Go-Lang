package main

import "fmt"

func main() {
	const (
		width = 25
		height = width * 2
	)

	fmt.Printf("area = %d\n", width*height)
}