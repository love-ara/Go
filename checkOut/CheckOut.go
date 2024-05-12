package main

import (
	"fmt"
	"time"
)

type app struct {
	customerName string
	cashierName  string
	subTotal     float64
	discount     float64
	amountPaid   float64
}

func (c *app) getCurrentDateAndTime() string {
	currentDateTime := time.Now()
	dateTimeFormatter := "02-Jan-06 03:04:05 PM"
	currentDateAndTime := currentDateTime.Format(dateTimeFormatter)

	return currentDateAndTime
}

func (c *app) setCustomerName(customerName string) {
	c.customerName = customerName
}

func (c *app) getCustomerName() string {
	return c.customerName
}

func (c *app) setCashierName(cashierName string) {
	c.cashierName = cashierName
}

func (c *app) getCashierName() string {
	return c.cashierName
}

func (c *app) setSubTotal(subTotal float64) {
	c.subTotal = subTotal
}

func (c *app) getSubTotal() float64 {
	return c.subTotal
}

func (c *app) setDiscount(discount float64) {
	if discount < 0 || discount > 100 {
		panic("Discount must be between 0 and 100")
	}
	c.discount = discount
}

func (c *app) getDiscount() float64 {
	return c.discount
}

func (c *app) getDiscountedTotal() float64 {
	return (c.discount / 100.0) * c.subTotal
}

func (c *app) getVAT() float64 {
	return c.subTotal * (7.5 / 100)
}

func (c *app) getBillTotal() float64 {
	return (c.subTotal - c.getDiscountedTotal()) + c.getVAT()
}

func (c *app) setAmountPaid(amountPaid float64) {
	if amountPaid < c.getBillTotal() {
		panic("Amount paid cannot be less than bill total")
	}
	c.amountPaid = amountPaid
}

func (c *app) getAmountPaid() float64 {
	return c.amountPaid
}

func (c *app) getBalance() float64 {
	balance := 0.0
	if c.amountPaid > c.getBillTotal() {
		return c.amountPaid - c.getBillTotal()
	}
	return balance
}

func main() {
	checkout := app{}

	fmt.Print("Enter customer name: ")
	var customerName string
	fmt.Scanln(&customerName)
	checkout.setCustomerName(customerName)

	fmt.Print("Enter cashier name: ")
	var cashierName string
	fmt.Scanln(&cashierName)
	checkout.setCashierName(cashierName)

	fmt.Print("Enter sub-total: ")
	var subTotal float64
	fmt.Scanln(&subTotal)
	checkout.setSubTotal(subTotal)

	fmt.Print("Enter discount (%): ")
	var discount float64
	fmt.Scanln(&discount)
	checkout.setDiscount(discount)

	fmt.Print("Enter amount paid: ")
	var amountPaid float64
	fmt.Scanln(&amountPaid)
	checkout.setAmountPaid(amountPaid)

	fmt.Println("Date and Time:", checkout.getCurrentDateAndTime())
	fmt.Println("Customer Name:", checkout.getCustomerName())
	fmt.Println("Cashier Name:", checkout.getCashierName())
	fmt.Println("Sub-total:", checkout.getSubTotal())
	fmt.Println("Discount:", checkout.getDiscount())
	fmt.Println("Discounted Total:", checkout.getDiscountedTotal())
	fmt.Println("VAT:", checkout.getVAT())
	fmt.Println("Bill Total:", checkout.getBillTotal())
	fmt.Println("Amount Paid:", checkout.getAmountPaid())
	fmt.Println("Balance:", checkout.getBalance())
}
