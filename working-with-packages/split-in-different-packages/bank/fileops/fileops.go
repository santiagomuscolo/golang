package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFile(file string, value *float64) error {
	parsedValue := fmt.Sprint(value)
	err := os.WriteFile(file, []byte(parsedValue), 0644)

	if err != nil {
		return errors.New("Error writing file: " + err.Error())
	}

	return nil
}

func ReadFile(file string) (float64, error) {
	data, err := os.ReadFile(file)

	if err != nil {
		return 1000, errors.New("Error reading file: " + err.Error())
	}

	parseValueToText := string(data)
	balance, _ := strconv.ParseFloat(parseValueToText, 64)

	if err != nil {
		return 1000, errors.New("Error parsing file: " + err.Error())
	}

	return balance, nil
}
