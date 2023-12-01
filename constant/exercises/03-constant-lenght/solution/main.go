package main

import "fmt"

func main() {
	const (
		home = " Milky Way Galaxy"
		lenght = len(home)
	)

	fmt.Printf("There are %d characters inside %q\n", lenght, home)
}