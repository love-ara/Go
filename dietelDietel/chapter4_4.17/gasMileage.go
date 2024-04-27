package main

import "fmt"

func gasMileage() {
	var totalMiles int64
	var totalGallons int64
	var tripCount int64

	fmt.Println("Welcome to the Mileage Calculator!")

	for {
		fmt.Print("Enter miles driven (or -1 to stop): ")
		var miles int64
		_, err := fmt.Scanf("%d", &miles)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			continue
		}
		if miles == -1 {
			break
		}

		fmt.Print("Enter gallons used: ")
		var gallons int64
		_, err = fmt.Scanf("%d", &gallons)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			continue
		}

		tripCount++
		totalMiles += miles
		totalGallons += gallons

		mpg := float64(miles) / float64(gallons)
		fmt.Printf("Miles per gallon on this trip = %.2f\n", mpg)
	}

	if tripCount > 0 {
		overallMPG := float64(totalMiles) / float64(totalGallons)
		fmt.Printf("Combined miles per gallon for all trips = %.2f\n", overallMPG)
	} else {
		fmt.Println("Not valid.")
	}
}
