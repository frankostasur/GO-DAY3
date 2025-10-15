package main

import (
	"fmt"

)

func main() {
	revenue := getUserInput("Enter your revenue:")
	expenses := getUserInput("Enter your expences:")
	taxRate := getUserInput("Enter your taxRate:")


	EBT, profit, ratio :=calculatevalues(revenue, expenses, taxRate)

	fmt.Println(EBT)
	fmt.Println(profit)
	fmt.Println(ratio)
}
func calculatevalues(revenue, expenses, taxRate float64) (float64 ,float64, float64){
	EBT := revenue - expenses
	profit :=  EBT * (1-taxRate/100)
	ratio := EBT/profit
	return EBT, profit ,ratio
}
func getUserInput(infoText string) float64{
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput

}
