package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func writeBalance(balance *float64) {
	parsedBalance := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(parsedBalance), 0644)
}

func readBalance() (float64, error) {
	data, err := os.ReadFile("balance.txt")

	if err != nil {
		return 1000, errors.New("Error reading balance file: " + err.Error())
	}

	parseBalanceToText := string(data)
	balance, _ := strconv.ParseFloat(parseBalanceToText, 64)

	if err != nil {
		return 1000, errors.New("Error parsing balance: " + err.Error())
	}

	return balance, nil
}
