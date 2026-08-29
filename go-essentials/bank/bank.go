package main

import "fmt"
import "os"
import "strconv"
import "errors"

func writeBalance(balance *float64) {
	parsedBalance := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(parsedBalance), 0644);
}

func readBalance() (float64, error) {
	data, err := os.ReadFile("balance.txt");

	if err != nil {
		return 1000, errors.New("Error reading balance file: " + err.Error())
	};

	parseBalanceToText := string(data);
	balance, _ := strconv.ParseFloat(parseBalanceToText, 64);

	if err != nil {
		return 1000, errors.New("Error parsing balance: " + err.Error())
	};

	return balance, nil;
}

func checkBalance(balance *float64){
	fmt.Printf("Your balance is: %.2f\n", *balance)
}

func deposit(balance *float64){
	var depositAmount float64;

	fmt.Print("Enter the amount to deposit: ");
	fmt.Scan(&depositAmount);

	if(depositAmount >= 0) {
		*balance += depositAmount;
		writeBalance(balance);
	} else {
		fmt.Println("Invalid deposit amount. Please enter a positive value.");
	}
}

func withdraw(balance *float64){
	var withdrawAmount float64;

	fmt.Print("Enter the amount to withdraw: ")
	fmt.Scan(&withdrawAmount);

	if(withdrawAmount > *balance) {
		fmt.Println("You don't have enough balance to withdraw that amount.");
	} else {
		*balance -= withdrawAmount;
	}
}

func exit (balance *float64){
	fmt.Println("Thank you for using Go bank. Goodbye!");
}

func main() {
	var option int;
	var balance, err = readBalance();

	if err != nil {
		fmt.Println("Error reading balance:", err)
		panic("Can't continue, sorry!.")
	}

	fmt.Println("Welcome to Go bank!");

	actions := map[int]func(*float64){
		1: checkBalance,
		2: deposit,
		3: withdraw,
		4: exit,
	}
	
	for option != 4 {
		fmt.Println("What do you want to do?");
		fmt.Println("1. Check balance");
		fmt.Println("2. Deposit money");
		fmt.Println("3. Withdraw money");
		fmt.Println("4. Exit");

		fmt.Print("Enter your option: ");
		fmt.Scan(&option);

		action, exists := actions[option];

		if exists {
			action(&balance);
		}else{
			fmt.Println("Invalid option. Please select a valid option.");
		}
	}


	fmt.Println("You selected option: ", option);
}