package main

import "fmt"

func largestNumber() int {
	var count int
	var number int
	var largest int
	for count = 0; count <= 10; count++ {
		fmt.Println("Enter number: ")
		_, err := fmt.Scan(&number)
		if err != nil {
			return 0
		}
		if number > largest {
			largest = number
		}
	}
	fmt.Println("The largest number is: ", largest)
	return largest
}
