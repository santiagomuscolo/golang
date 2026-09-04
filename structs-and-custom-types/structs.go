package main

import "fmt"
import "time"

type user struct {
	firstName string;
	lastName string;
	birthDate string;
	createdAt time.Time;
}

func main() {
	firstName := getUserData("Please enter your first name: ");
	lastName := getUserData("Please enter your last name: ");
	birthDate := getUserData("Please enter your birth date (YYYY-MM-DD): ");

	userData := user{
		firstName: firstName,
		lastName: lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}
	outputUserData(userData)
}

func outputUserData(u user) {
	fmt.Println(u.firstName, u.lastName, u.birthDate);
}

func getUserData (promptText string) string {
	fmt.Println(promptText);
	var value string;
	fmt.Scan(&value);
	return value;
};