package model

import "fmt"

type User struct {
	Username string
	Password string
	Email string
	Phone string
}

func (user *User) Init(name string, psw string, email string, phone string) {
	user.Username = name
	user.Password = psw
	user.Email = email
	user.Phone = phone
}

func (user User) GetUsername() string {
	return user.Username
}

func (user *User) SetUsername(username string) {
	user.Username = username
}

func (user User) GetPassword() string {
	return user.Password
}

func (user *User) SetPassword(password string) {
	user.Password = password
}

func (user User) GetEmail() string {
	return user.Email
}

func (user *User) SetEmail(email string) {
	user.Email = email
}

func (user User) GetPhone() string {
	return user.Phone
}

func (user *User) SetPhone(phone string) {
	user.Phone = phone
}

func (user User) String() {
	fmt.Println(user.GetUsername() + ", " + user.GetEmail() + ", " + user.GetPhone())
}