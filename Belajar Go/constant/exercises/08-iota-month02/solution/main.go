package main

import "fmt"

func main() {
	const(
		_ = iota
		january
		february
		march
	)
	fmt.Println(january, february, march)
}