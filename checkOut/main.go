package main

import (
	"errors"
	"fmt"
	"time"
)

type CheckOutApp struct {
	customerName string
	cashierName  string
	subTotal     float64
	discount     float64
	amountPaid   float64
}

func (c *CheckOutApp) getCurrentDateAndTime() string {
	currentDateTime := time.Now()
	dateTimeFormatter := "02-Jan-06 03:04:05 PM"
	currentDateAndTime := currentDateTime.Format(dateTimeFormatter)

	return currentDateAndTime
}

func (c *CheckOutApp) setCustomerName(customerName string) {
	c.customerName = customerName
}

func (c *CheckOutApp) getCustomerName() string {
	return c.customerName
}

func (c *CheckOutApp) setCashierName(cashierName string) {
	c.cashierName = cashierName
}

func (c *CheckOutApp) getCashierName() string {
	return c.cashierName
}

func (c *CheckOutApp) setSubTotal(subTotal float64) {
	c.subTotal = subTotal
}

func (c *CheckOutApp) getSubTotal() float64 {
	return c.subTotal
}

func (c *CheckOutApp) setDiscount(discount float64) {
	if discount < 0 || discount > 100 {
		panic("Discount must be between 0 and 100")
	}
	c.discount = discount
}

func (c *CheckOutApp) getDiscount() float64 {
	return c.discount
}

func (c *CheckOutApp) getDiscountedTotal() float64 {
	return (c.discount / 100.0) * c.subTotal
}

func (c *CheckOutApp) getVAT() float64 {
	return c.subTotal * (7.5 / 100)
}

func (c *CheckOutApp) getBillTotal() float64 {
	return (c.subTotal - c.getDiscountedTotal()) + c.getVAT()
}
func (c *CheckOutApp) setAmountPaid(amountPaid float64) {
	if amountPaid < c.getBillTotal() {
		err := errors.New("amount paid cannot be less than bill total")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	c.amountPaid = amountPaid
}

func (c *CheckOutApp) getAmountPaid() float64 {
	return c.amountPaid
}

func (c *CheckOutApp) getBalance() float64 {
	balance := 0.0
	if c.amountPaid > c.getBillTotal() {
		return c.amountPaid - c.getBillTotal()
	}
	return balance
}

func main() {
	checkout := &CheckOutApp{}

	fmt.Print("What is the customer's Name? ")
	var customerName string
	_, err := fmt.Scanln(&customerName)
	if err != nil {
		return
	}
	checkout.setCustomerName(customerName)

	for {
		fmt.Print("What the customer buy? ")
		var itemInfo string
		_, err := fmt.Scanln(&itemInfo)
		if err != nil {
			return
		}

		fmt.Print("How many pieces: ")
		var quantity int
		_, err = fmt.Scanln(&quantity)
		if err != nil {
			return
		}

		fmt.Print("How much per unit: ")
		var price float64
		_, err = fmt.Scanln(&price)
		if err != nil {
			return
		}

		fmt.Print("\nAdd more items (yes/no)? ")
		var sentinelValue string
		_, err = fmt.Scanln(&sentinelValue)
		if err != nil {
			return
		}
		if sentinelValue == "no" {
			break
		}

		itemBought = append(itemBought, itemInfo)
		quantityPurchased = append(quantityPurchased, quantity)
		priceOfItem = append(priceOfItem, price)
	}

	fmt.Print("Enter cashier's name  ")
	var cashierInfo string
	_, err = fmt.Scanln(&cashierInfo)
	if err != nil {
		return
	}
	checkout.setCashierName(cashierInfo)

	fmt.Print("How much discount will the customer get? ")
	var discount float64
	_, err = fmt.Scanln(&discount)
	if err != nil {
		return
	}
	checkout.setDiscount(discount)

	sum := 0.0
	for i := range itemBought {
		sum += priceOfItem[i] * float64(quantityPurchased[i])
	}
	checkout.setSubTotal(sum)

	displayInvoice(checkout)
	setPayment(checkout)
	displayReceipt(checkout)

	fmt.Println("\nThank you for shopping with us!")
}

func displayReceipt(checkout *CheckOutApp) {
	fmt.Printf("\n\nSEMICOLON STORES\nMAIN BRANCH\nLOCATION: 312 HERBERT MACAULAY WAY, SABO YABA, LAGOS.\nTEL: 03293828343\nDate: %s\nCashier: %s\nCustomer Name: %s\n",
		checkout.getCurrentDateAndTime(), checkout.getCashierName(), checkout.getCustomerName())

	fmt.Printf("\n=======================================================\n")
	fmt.Printf("         ITEM           QTY     PRICE        TOTAL(NGN)\n")
	fmt.Printf("-------------------------------------------------------\n")

	for i := range itemBought {
		fmt.Printf("   %10s           %3d     %10.2f       %10.2f\n", itemBought[i], quantityPurchased[i], priceOfItem[i], priceOfItem[i]*float64(quantityPurchased[i]))
	}

	fmt.Printf("-------------------------------------------------------\n")
	fmt.Printf("\t\t\t%13s\t  %10.2f\n\t\t\t%13s\t  %10.2f\n\t\t\t%13s\t  %10.2f\n", "Sub Total:", checkout.getSubTotal(), "Discount:", checkout.getDiscountedTotal(), "VAT @ 7.50%:", checkout.getVAT())
	fmt.Printf("=======================================================\n")
	fmt.Printf("\t\t\t%13s\t  %10.2f\n\t\t\t%13s\t  %10.2f\n\t\t\t%13s\t  %10.2f\n", "Bill Total:", checkout.getBillTotal(), "Amount Paid:", checkout.getAmountPaid(), "Balance:", checkout.getBalance())
	fmt.Printf("=======================================================\n\t  THANK YOU FOR YOUR PATRONAGE\n=======================================================\n")
}

func displayInvoice(checkout *CheckOutApp) {
	fmt.Printf("\nSEMICOLON STORES\nMAIN BRANCH\nLOCATION: 312 HERBERT MACAULAY WAY, SABO YABA, LAGOS.\nTEL: 03293828343\nDate: %s\nCashier: %s\nCustomer Name: %s\n",
		checkout.getCurrentDateAndTime(), checkout.getCashierName(), checkout.getCustomerName())

	fmt.Printf("\n=======================================================\n")
	fmt.Printf("         ITEM           QTY     PRICE        TOTAL(NGN)\n")
	fmt.Printf("-------------------------------------------------------\n")

	for i := range itemBought {
		fmt.Printf("   %10s           %3d     %10.2f       %10.2f\n", itemBought[i], quantityPurchased[i], priceOfItem[i], priceOfItem[i]*float64(quantityPurchased[i]))
	}

	fmt.Printf("-------------------------------------------------------\n")
	fmt.Printf("\t\t\t%13s\t  %10.2f\n\t\t\t%13s\t  %10.2f\n\t\t\t%13s\t  %10.2f\n", "Sub Total:", checkout.getSubTotal(), "Discount:", checkout.getDiscountedTotal(), "VAT @ 7.50%:", checkout.getVAT())
	fmt.Printf("=======================================================\n\t\t\t%13s\t  %10.2f\n   THIS IS NOT A RECEIPT, KINDLY PAY %.2f\n=======================================================\n",
		"Bill Total:", checkout.getBillTotal(), checkout.getBillTotal())
}

func setPayment(checkout *CheckOutApp) {
	fmt.Print("\nHow much did the customer give to you? ")
	var amountPaid float64
	_, err := fmt.Scanln(&amountPaid)
	if err != nil {
		return
	}
	checkout.setAmountPaid(amountPaid)
}

var (
	itemBought        []string
	quantityPurchased []int
	priceOfItem       []float64
)
