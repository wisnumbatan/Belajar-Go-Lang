package main 

import "fmt"

func main() {
	cm:=100
	m:=cm/100
	fmt.Printf("%dcm is %dm\n",cm,m)

	cm = 200
	m =cm /100 
	fmt.Printf("%dcm is %dm\n",cm,m)
}