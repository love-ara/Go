package main

import "fmt"

var weeklyPay float64
var merchandise float64

func salesCommissionCalculator() {
	fmt.Println("Sales commission calculator")

	weeklyPay = 200.00
	fmt.Println("Your weekly pay is N", weeklyPay)

	fmt.Println("Enter the number items you sold")
	var num int
	_, err2 := fmt.Scanf("%d", &num)
	if err2 != nil {
		fmt.Println("Invalid input")
		return
	}
	for i := 0; i < num; i++ {
		fmt.Println("Enter amount of item ", i+1)
		var amount float64
		_, err := fmt.Scanf("%f", &amount)
		if err != nil {
			fmt.Println("Invalid input")
			return
		}
		merchandise += amount
	}

	var commission float64
	commission = 0.09 * float64(merchandise)

	earnings := weeklyPay + commission

	fmt.Println("Your earning is N", earnings)

}
