package main

import "fmt"
import "os"

func main() {
	var revenue, expenses, taxRate float64;

	outputTextAndScan("==== ENTER THE REVENUE ==== \n", &revenue);
	if(revenue <= 0) {
		panic("Revenue must be greater than zero. Please enter a valid revenue amount.");
	};

	outputTextAndScan("==== ENTER THE EXPENSES ==== \n", &expenses);
	if(expenses <= 0) {
		panic("Expenses must be greater than zero. Please enter a valid expenses amount.");
	};

	outputTextAndScan("==== ENTER THE TAX RATE ==== \n", &taxRate);
	if(taxRate <= 0) {
		panic("Tax rate must be greater than zero. Please enter a valid tax rate.");
	};


	earningsBeforeTax, profit, ratio := calculateProfit(revenue, expenses, taxRate);

	fmt.Printf("==== EARNINGS BEFORE TAX ==== %.2f\n", earningsBeforeTax);
	fmt.Printf("==== PROFIT ==== %.2f\n", profit);
	fmt.Printf("==== RATIO ==== %.2f\n", ratio);
}

func outputTextAndScan(text string, value *float64) {
	fmt.Print(text);
	fmt.Scan(value);
}

func calculateProfit(revenue, expenses, taxRate float64) (erbft float64, pf float64, rt float64) {
	erbft = revenue - expenses

	// profit = bruto × (1 - impuesto)
	pf = erbft * (1 - taxRate/100)
	rt = erbft / pf

	data := []byte(fmt.Sprintf("Earnings Before Tax: %.2f\nProfit: %.2f\nRatio: %.2f\n", erbft, pf, rt))
	err := os.WriteFile("profit.txt", data, 0644); 
	if err != nil {
		fmt.Println("Error writing profit.txt:", err);
	};

	return erbft, pf, rt
}