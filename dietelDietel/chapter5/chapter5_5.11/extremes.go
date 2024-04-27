package main

import "fmt"

func extremes() {
	var count int
	var maximum int
	var minimum int

	fmt.Print("Enter the number of values: ")
	_, err := fmt.Scanf("%d", &count)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid integer.")
		return
	}

	for index := 1; index <= count; index++ {
		var number int
		fmt.Printf("Enter value  %d: ", index)
		_, err := fmt.Scanf("%d", &number)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			continue
		}

		if index == 1 {
			maximum, minimum = number, number
		} else {
			if number > maximum {
				maximum = number
			}
			if number < minimum {
				minimum = number
			}
		}
	}

	fmt.Println("Maximum value:", maximum)
	fmt.Println("Minimum value:", minimum)
}
