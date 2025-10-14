package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64


	fmt.Print("Enter your revenue:")
	fmt.Scan(&revenue)

	fmt.Print("Enter your expences:")
	fmt.Scan(&expenses)

	fmt.Print("Enter your taxRate:")
	fmt.Scan(&taxRate)

	EBT := revenue - expenses
	profit :=  EBT * (1-taxRate/100)
	ratio := EBT/profit

	fmt.Println(EBT)
	fmt.Println(profit)
	fmt.Println(ratio)
}