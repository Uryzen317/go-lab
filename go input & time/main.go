package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("hello dear user")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the demanded input : ")

	//read user input till you reach a new line
	//notice the singel comma
	userInput , _ := reader.ReadString('\n')
	fmt.Printf("you've entered the value of %v" , userInput)

	//conveting the input strign into an number
	addedUserNumber , err := strconv.ParseFloat(strings.TrimSpace(userInput) , 64);
	if err != nil {
		log.Fatal(err)
	}else{
		fmt.Printf("added one to your input : %v" , addedUserNumber + 1)
	}


	/////tiem section
	currentTime := time.Now();
	fmt.Println("\n" , currentTime)

	//formatting time
	fmt.Println("sorted time of right now is : " , currentTime.Format("01-02-2006 Monday 15:04:05"))

	//getting a custome time
	fmt.Println(time.Date(2020 ,time.April ,12 , 23 , 23 , 0 , 0 , time.UTC).Format("01-02-2006 Monday"))
}