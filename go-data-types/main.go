package main

import (
	"fmt"
	"sort"
)

func main() {
	var firstArray [5]string
	firstArray[0] = "apple"
	firstArray[1] = "orange"
	firstArray[2] = "potato"
	fmt.Println("first aray value is : ", firstArray ," with the lenght of : " , len(firstArray))

	var secondArray = [5]string{"ali" , "mohammad" , "hossein"}
	fmt.Println("the second array value is as following : " , secondArray)

	firstSlice := []string{"first" , "second" , "third" , "fourth" , "fifth"}
	fmt.Printf("parts of first slice values are : %v \n" , append(firstSlice[1:3]))

	secondSlice := make([]int , 4)
	secondSlice[0] = 10;
	secondSlice[1] = 20;
	secondSlice[2] = 15;
	secondSlice[3] = 40;

	secondSlice = append(secondSlice , 56 , 74 , 12);
	fmt.Printf("secodn slice value is : %v \n" , secondSlice)

	sort.Ints(secondSlice);
	fmt.Println("the sorted form of second slice is : " , secondSlice , "and the sortion status of it is : ", sort.IntsAreSorted(secondSlice))

	var indexOfDeletion int = 2
	fmt.Println("third argument is deleted : " , append(secondSlice[:indexOfDeletion] , secondSlice[indexOfDeletion+1:]...))

	firstMap := make(map[string]string)
	firstMap["first"] = "something"
	firstMap["second"] = "something else"
	firstMap["third"] = "again something else"

	delete(firstMap , "third")

	fmt.Println("the first value of the first map is : " , firstMap["first"])
	for key , value := range(firstMap){
		fmt.Println("for the key of : " , key , "the value equals : " ,value)
	}

	firstStruct := FirstStructType{"mohammad" , 16 , true}
	fmt.Println("simple displation of first struct : " , firstStruct)
	fmt.Printf("more advanced : %+v \n" , firstStruct)
	fmt.Println("destructured form of first struct for name and age of : " , firstStruct.Name , " and ", firstStruct.Age)

	firstStruct.FirstStructFunction()
	fmt.Println("yet the real user age has not been changed : " , firstStruct.Age)
}

type FirstStructType struct{
	Name 	string
	Age 	int
	status 	bool
}

func (f FirstStructType) FirstStructFunction(){
	f.Age = 999 ;
	fmt.Println("manipulated user age is : " , f.Age);
}