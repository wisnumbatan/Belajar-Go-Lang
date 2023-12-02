package main

import "fmt"

func main() {
	const (
		hoursInDay = 24
		dayInWeek = 7
		hoursPerWeek = hoursInDay * dayInWeek
	)
	fmt.Printf("There are %d hours in %d weeks.\n",hoursPerWeek*5, 5)
}