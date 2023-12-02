package main

import (
	"fmt"
	"time"
)

func main() {
	const later = 10

	hours, _ := time.ParseDuration("1h")

	fmt.Printf("%s later...\n",hours*later)
}