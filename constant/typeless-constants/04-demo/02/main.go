package main

import (
	"fmt"
	"time"
)

func main(){
	t := time.Second * 10
	fmt.Println(t)

	i := 10
	
	t = time.Second * time.Duration(i)

	fmt.Println(t)

	const c = 10 

	t = time.Second * c

	fmt.Println(t)

}