package main

import (
	"fmt"
	"sort"
)

func twoLargestNumbers() {
	var count int
	fmt.Println("Enter the number of elements")
	_, err := fmt.Scan(&count)
	if err != nil {
		return
	}

	values := make([]float64, count)

	for index := 0; index < count; index++ {
		var number float64
		fmt.Printf("Enter unique value  %d: ", index+1)
		_, err := fmt.Scanf("%f", &number)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			index--
			continue
		}

		if contains(values[:index], number) {
			fmt.Println("Number already entered. Please choose a unique value.")
			index--
			continue
		}

		values[index] = number
	}

	sort.Slice(values, func(index, indexTwo int) bool {
		return values[index] > values[indexTwo]
	})

	fmt.Printf("The two largest unique values are %.2f and %.2f\n", values[0], values[1])
}

func contains(slice []float64, value float64) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
