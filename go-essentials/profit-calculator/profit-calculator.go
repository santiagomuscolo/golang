package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64;

	outputTextAndScan("==== ENTER THE REVENUE ==== \n", revenue);
	outputTextAndScan("==== ENTER THE EXPENSES ==== \n", expenses);
	outputTextAndScan("==== ENTER THE TAX RATE ==== \n", taxRate);

	earningsBeforeTax, profit, ratio := calculateProfit(revenue, expenses, taxRate);

	fmt.Printf("==== EARNINGS BEFORE TAX ==== %.2f\n", earningsBeforeTax);
	fmt.Printf("==== PROFIT ==== %.2f\n", profit);
	fmt.Printf("==== RATIO ==== %.2f\n", ratio);
}

func outputTextAndScan(text string, value float64) {
	fmt.Print(text);
	fmt.Scan(&value);
}

func calculateProfit (revenue, expenses, taxRate float64) (erbft float64, pf float64, rt float64) {
    erbft = revenue - expenses;

	// profit = bruto × (1 - impuesto)
	pf = erbft * (1 - taxRate / 100);
	rt = erbft / pf;

	return erbft, pf, rt;
}