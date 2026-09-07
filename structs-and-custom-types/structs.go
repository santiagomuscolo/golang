package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string;
	lastName string;
	birthDate string;
	createdAt time.Time;
}

func (u *user) outPutUserData() {
	fmt.Println(u.firstName, u.lastName, u.birthDate);
}

func (u *user) cleanUserName() {
	u.firstName = "";
	u.lastName = "";
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

	// outputUserData(&userData)
	userData.outPutUserData();
	userData.cleanUserName();
	userData.outPutUserData();
}

// func outputUserData(u *user) {
	// Go shortcut for dereferencing a pointer is to use the dot operator directly on the pointer.
// 	fmt.Println(u.firstName, u.lastName, u.birthDate);
// }

func getUserData (promptText string) string {
	fmt.Printf(promptText);
	var value string;
	fmt.Scan(&value);
	return value;
};