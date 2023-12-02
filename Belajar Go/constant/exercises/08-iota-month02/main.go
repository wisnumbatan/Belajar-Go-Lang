package main

import "fmt"

func main() {
	const _ = iota
	const (
		january = 1 + iota
		february
		march
	)
	fmt.Println(january, february, march)

}