package main

import (
	"fmt"
	"time"
)

func main() {
	type Duration int64

	const (
		Nanosecond Duration = 1
		Microsecond Duration = 1000 * Nanosecond
		Millisecond Duration = 1000 * Microsecond
		Second      Duration = 1000 * Millisecond
		Minute      Duration = 60 * Second
		Hour        Duration = 60 * Minute
	)

	fmt.Printf("Nanosecond is %T\n", Nanosecond)
	fmt.Printf("Microsecond is %T\n", Microsecond) 
	fmt.Printf("Millisecond is %T\n", Millisecond)
	fmt.Printf("Second is %T\n", Second)
	fmt.Printf("Minute is %T\n", Minute)
	fmt.Printf("Hour is %T\n", Hour)

	fmt.Printf("Nanosecond is %d\n", time.Nanosecond)
	fmt.Printf("Microsecond is %d\n", time.Microsecond)
	fmt.Printf("Millisecond is %d\n", time.Millisecond)
	fmt.Printf("Second is %d\n", time.Second)
	fmt.Printf("Minute is %d\n", time.Minute)
	fmt.Printf("Hour is %d\n", time.Hour)
}