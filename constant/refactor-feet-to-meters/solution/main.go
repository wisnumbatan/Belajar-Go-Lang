package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	const (
		feetInMeters float64 = 0.3048
		feetInYards          = feetInMeters / 0.9144
	)

	if len(os.Args) < 2 {
		fmt.Println("Please provide a number of feet.")
		return
	}

	arg := os.Args[1]

	feet, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Println("Invalid input. Please provide a valid number.")
		return
	}

	meters := feet * feetInMeters
	yards := feet / feetInYards

	fmt.Printf("%g feet is %g meters.\n", feet, meters)
	fmt.Printf("%g feet is %g yards.\n", feet, math.Round(yards))
}
