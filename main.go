package main

import (
	"fmt"
	"errors"
	"os"

)

func main() {
	revenue, err:= getUserInput("Enter your revenue:")

	if err != nil{
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Enter your expences:")

	if err != nil{
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("Enter your taxRate:")

	if err != nil{
		fmt.Println(err)
		return
	}

	EBT, profit, ratio :=calculatevalues(revenue, expenses, taxRate)

	fmt.Println(EBT)
	fmt.Println(profit)
	fmt.Println(ratio)
	storeResults(EBT, profit, ratio)
}
func storeResults(EBT, profit, ratio float64)  {
	results := fmt.Sprintf("EBT: %0.1f\n Profit: %0.1f\n Ratio: %0.3f\n", EBT, profit, ratio)
	os.WriteFile("results.txt", []byte (results), 0644)
	
}
func calculatevalues(revenue, expenses, taxRate float64) (float64 ,float64, float64){
	EBT := revenue - expenses
	profit :=  EBT * (1-taxRate/100)
	ratio := EBT/profit
	return EBT, profit ,ratio
}
func getUserInput(infoText string) (float64, error){
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("value must be a positive number.")
	}
	return userInput, nil

}
