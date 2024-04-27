package main

import "fmt"

func creditLimitCalculator() {
	fmt.Println("Check your credit limit")
	fmt.Println("Please enter the questions below correctly")

	fmt.Println("Enter your account number: ")
	var accountNumber float64
	_, err := fmt.Scanf("%f", &accountNumber)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid integer.")
		return
	}

	fmt.Println("How much was your balance at the beginning of the month?")
	var balance float64
	_, err = fmt.Scanf("%f", &balance)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid integer.")
		return
	}

	fmt.Println("How much have you spend in total this month?")
	var totalCharges float64
	_, err = fmt.Scanf("%f", &totalCharges)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid integer.")
		return
	}

	fmt.Println("How much is the total credit applied to you this month?")
	var monthlyCredits float64
	_, err = fmt.Scanf("%f", &monthlyCredits)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid integer.")
		return
	}

	fmt.Println("How much is your allowed credit limit?")
	var allowedCredits float64
	_, err = fmt.Scanf("%f", &allowedCredits)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid integer.")
		return

	}

	var newBalance float64
	newBalance = balance + totalCharges - monthlyCredits

	fmt.Println("Your balance is ", newBalance)

	if newBalance > allowedCredits {
		fmt.Println("Credit limit exceeded")
		fmt.Println("You've exceeded your credit limit by: ", newBalance-allowedCredits)
	} else {
		fmt.Println("Your balance is ", allowedCredits-newBalance)
	}
}
