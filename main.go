// Created by: Infinity de Guzman
// Created on: May 2021
//
// This program converts fahrenheit to celsius

package main

import (
	"fmt"

	"github.com/leekchan/accounting"
)

func main() {
	// This function converts fahrenheit to celsius
	var fahrenheit float64

	// input
	accountingFormater := accounting.Accounting{Precision: 2}

	fmt.Println("This program converts fahrenheit to celsius.")
	fmt.Println()
	fmt.Print("Enter the temperature in fahrenheit: ")
	fmt.Scanln(&fahrenheit)

	// process
	var celsius = (fahrenheit - 32) * 5 / 9

	// output
	fmt.Println("The temperature is: ", accountingFormater.FormatMoney(celsius), "°C")
	fmt.Println("\nDone.")
}
