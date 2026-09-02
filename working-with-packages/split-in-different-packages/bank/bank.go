package main

import "fmt"
import "bank/fileops"

const balanceFile = "balance.txt"

func main() {
	var option int
	var balance, err = fileops.ReadFile(balanceFile)

	if err != nil {
		fmt.Println("Error reading balance:", err)
		panic("Can't continue, sorry!.")
	}

	fmt.Println("Welcome to Go bank!")

	actions := map[int]func(*float64){
		1: checkBalance,
		2: deposit,
		3: withdraw,
		4: exit,
	}

	for option != 4 {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		fmt.Print("Enter your option: ")
		fmt.Scan(&option)

		action, exists := actions[option]

		if exists {
			action(&balance)
		} else {
			fmt.Println("Invalid option. Please select a valid option.")
		}
	}

	fmt.Println("You selected option: ", option)
}
