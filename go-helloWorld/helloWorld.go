package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(quote.Go())

	for counter := 0; counter < 10; counter++ {
		fmt.Println(counter)
	}

	for secondCounter := 0 ; secondCounter < secondCounter+1 ; secondCounter ++{
		secondCounter -- ;
	}
}
