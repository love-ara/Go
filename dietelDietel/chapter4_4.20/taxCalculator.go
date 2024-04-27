package main

import "fmt"

func taxCalculator() {
	var taxRate = 0.15
	var extraTax = 0.2
	var earnings float64
	var totalTax float64

	fmt.Println("Enter your name: ")
	var name string
	_, err := fmt.Scanf("%s", &name)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	fmt.Println("How much do you earn?")
	_, err = fmt.Scanf("%f", &earnings)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	if earnings <= 30000 {
		totalTax = taxRate * earnings
		earningAfterTax := earnings - totalTax
		fmt.Println("Your total tax:", totalTax)
		fmt.Println("Your earning after tax:", earningAfterTax)
	} else if earnings > 30000 {
		extraIncome := earnings - 30000
		var tax = extraTax * extraIncome
		var this = taxRate * 30000
		totalTax := tax + this
		var earningAfterTax = earnings - totalTax
		fmt.Println("Your total tax:", totalTax)
		fmt.Println("Your earning after tax:", earningAfterTax)
	}
}
