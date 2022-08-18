package main

import "fmt"

func main() {
	var firstPointer *int
	fmt.Println("first pinter value is : "  ,firstPointer)

	var firstVariable int = 20 
	secondPointer := &firstVariable
	fmt.Println("second pinter address : " , secondPointer)
	fmt.Println("second pinter value : " , *secondPointer)

	*secondPointer = *secondPointer +1 
	fmt.Println("the edited value of first variable : " , firstVariable)
}