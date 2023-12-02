package main

import "fmt"

func main() {
	const(
		spring = (iota + 1) * 3
		summer
		Fall
		winter
	)
	fmt.Println(winter, spring, summer, Fall)
}