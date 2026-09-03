package main

import "fmt"

func main() {
	age := 32;
	var agePointer *int = &age;

	fmt.Println("Age:", *agePointer);

    adultYears := getAdultYears(*agePointer);
	fmt.Println("Adult years:", adultYears);
}

func getAdultYears(age int) int {
	return age - 18;
}