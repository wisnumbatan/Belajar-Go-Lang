package main

import "fmt"

func main() {
	const meter int = 100
	
	cm := 100
	m := cm / meter
	fmt.Printf("%dcm is %dm\n",cm,m)

	cm = 200
	m = cm / meter
	fmt.Printf("%dcm is %dm\n",cm,m)
}