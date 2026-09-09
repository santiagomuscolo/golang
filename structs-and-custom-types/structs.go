package main

import (
	"fmt"
	"structs-and-custom-types/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ");
	lastName := getUserData("Please enter your last name: ");
	birthDate := getUserData("Please enter your birth date (YYYY-MM-DD): ");

	appUser, err := user.New(firstName, lastName, birthDate);

	if err != nil {
		fmt.Println(err);
		return;
	}

	appUser.OutputUserData();
	appUser.CleanUserName();
	appUser.OutputUserData();
}

func getUserData (promptText string) string {
	fmt.Printf(promptText);
	var value string;
	fmt.Scan(&value);
	return value;
};