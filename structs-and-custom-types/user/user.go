package user;

import "time";
import "fmt";
import "errors";

type User struct {
	FirstName string;
	LastName string;
	BirthDate string;
	CreatedAt time.Time;
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if(firstName == "" || lastName == "" || birthDate == "") {
		return nil, errors.New("Error: All fields are required to create a new user.");
	};

	return &User{
		FirstName: firstName,
		LastName: lastName,
		BirthDate: birthDate,
		CreatedAt: time.Now(),
	}, nil;
};

func (u *User) OutputUserData() {
	fmt.Println(u.FirstName, u.LastName, u.BirthDate);
}

func (u *User) CleanUserName() {
	u.FirstName = "";
	u.LastName = "";
}

