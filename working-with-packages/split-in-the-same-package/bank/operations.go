package main

import "fmt"

func checkBalance(balance *float64) {
	fmt.Printf("Your balance is: %.2f\n", *balance)
}

func deposit(balance *float64) {
	var depositAmount float64

	fmt.Print("Enter the amount to deposit: ")
	fmt.Scan(&depositAmount)

	if depositAmount >= 0 {
		*balance += depositAmount
		writeBalance(balance)
	} else {
		fmt.Println("Invalid deposit amount. Please enter a positive value.")
	}
}

func withdraw(balance *float64) {
	var withdrawAmount float64

	fmt.Print("Enter the amount to withdraw: ")
	fmt.Scan(&withdrawAmount)

	if withdrawAmount > *balance {
		fmt.Println("You don't have enough balance to withdraw that amount.")
	} else {
		*balance -= withdrawAmount
	}
}

func exit(balance *float64) {
	fmt.Println("Thank you for using Go bank. Goodbye!")
}
