package main

import (
	"fmt"
	"math"
)

func main() {

	// si solamente queremos que GO infiera el tipo se podria hacer asi => investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5;
	var investmentAmount float64;
	var years float64;
	var expectedReturnRate float64;
	
	outputTextAndScan("Enter the investment amount: ", investmentAmount);
	outputTextAndScan("Enter the number of years: ", years);
	outputTextAndScan("Enter the expected return rate (in %): ", expectedReturnRate);

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years);

	fmt.Printf("Future Value: %.2f\n", futureValue);
	fmt.Printf("Future Real Value: %.2f\n", futureRealValue);
}

func outputTextAndScan(text string, value float64) {
	fmt.Print(text);
	fmt.Scan(&value);
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	const inflationRate = 2.5;

	fv = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years);
	rfv = fv / math.Pow(1 + inflationRate / 100, years);

	return fv, rfv
}