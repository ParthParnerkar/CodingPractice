package main

import "fmt"

func main() {
	var number int
	fmt.Println("Enter Your number: ")
	fmt.Scanln(&number)

	if number%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}
}
