package main

import "fmt"

func main() {
	score, valid := 5, true

	if score > 3 && valid {
		fmt.Println("Good")
	} else if score == 3 {
		fmt.Println("On the edge")
	} else {
		fmt.Println("Low")
	}
}