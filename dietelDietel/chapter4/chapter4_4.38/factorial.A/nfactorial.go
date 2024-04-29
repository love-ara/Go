package main

import "fmt"

func nFactorial() {

	fmt.Println("Enter a non-negative number")
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Not a valid input")
		return
	}
	if number < 0 {
		fmt.Println("Not a valid number")
		return
	}
	if number == 0 {
		fmt.Println("1")
		return
	}
	result := 1
	for index := number; index > 0; index-- {
		result *= index

	}

	fmt.Println("n! = ", result)

}
