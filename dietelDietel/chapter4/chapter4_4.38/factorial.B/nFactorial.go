package main

import (
	"fmt"
)

func nFactorial() {

	var exponential float64
	exponential = 1.0
	fmt.Println("Enter the number of terms you want to calculate for")
	var term int
	_, err := fmt.Scan(&term)
	if err != nil {
		return
	}

	for index := 1; index <= term; index++ {
		factorial := 1.0
		for indexTwo := 1; indexTwo <= index; indexTwo++ {
			factorial *= float64(indexTwo)
		}
		exponential += 1 / factorial
	}
	fmt.Println("e = ", exponential)
}
