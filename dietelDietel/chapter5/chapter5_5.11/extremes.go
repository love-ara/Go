package main

import "fmt"

func extremes() {
	fmt.Println()
	var count int
	var number int
	var firstLargest int
	var secondLargest int

	for count = 0; count < 10; count++ {

		fmt.Println("Enter first number: ")
		_, err := fmt.Scanf("%d", &number)
		if err != nil {
			fmt.Println("Invalid input")
			return
		}

		if number > firstLargest {
			firstLargest = number
		}
		if number < secondLargest {
			secondLargest = number
		}
	}
	fmt.Println(firstLargest, secondLargest)

}
