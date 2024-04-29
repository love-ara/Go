package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter your number: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		return
	}
	result := ConstantEx(number)
	fmt.Printf("factorial = %f", result)
}
