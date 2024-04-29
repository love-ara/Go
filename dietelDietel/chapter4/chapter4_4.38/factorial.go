package main

import "math"

func factorial(number int) int {
	if number <= 0 {
		return 1
	}

	return number * factorial(number-1)
}

func ConstantE(number int) float64 {
	if number <= 0 {
		return 1.0
	}
	e := 1.0

	for index := 1; index <= number; index++ {
		e += 1.0 / float64(factorial(index))
	}
	return e
}

func ConstantEx(number int) float64 {
	if number <= 0 {
		return 1.0
	}
	return math.Pow(ConstantE(number), float64(number))
}
