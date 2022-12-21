package main

import (
	"fmt"
	"math"
)

/*
description: calculate per month emi amount of any loan
param: principalAmount, numberOfYears, interestRate float64
*/
func calculateEMI(principalAmount, numberOfYears, interestRate float64) {

	monthlyInterestRate := interestRate / (12 * 100)
	totalPayments := numberOfYears * 12
	emi := (principalAmount * monthlyInterestRate) / (1 - (math.Pow((1 + monthlyInterestRate), -totalPayments)))

	fmt.Println("Per Month EMI: ", emi, " Rs.")
}

func main() {

	var principalAmount, numberOfYears, interestRate float64

	fmt.Println("To Compute Per Month EMI Of A Loan")
	fmt.Print("Enter Principle Amount: ")
	fmt.Scanln(&principalAmount)
	fmt.Print("Enter Number of Years: ")
	fmt.Scanln(&numberOfYears)
	fmt.Print("Enter Interest Rate: ")
	fmt.Scanln(&interestRate)

	calculateEMI(principalAmount, numberOfYears, interestRate)

}
