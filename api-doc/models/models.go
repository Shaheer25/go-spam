package models

import "database/sql"

type User struct {
	UserName string `json:"username"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	IsActive bool   `json:"isactive"`
}

var users =[]User{
	{
	UserName: "User1",
	FullName:"User Name",
	Email: "Email@email.com",
	IsActive : true,
},

{
	UserName: "User12",
	FullName:"User Name2",
	Email: "Email@gmail.com",
	IsActive : false,
},
	}



type Country struct {
	Name string `json:"name"`
}



var countries =[]Country{
	{Name : "India"},
	{Name : "Canada"},
	{Name : "Vatican City"},
	{Name : "Rwanda"},
	{Name : "Germany"},
	{Name : "Swiss"},
}


func ListUser() []User {
	return users
}

func GetUser(username string) (*User, error) {
	for _, user := range users {
		if user.UserName == username {
			return &user, nil
		}
	}
	return nil, sql.ErrNoRows
}

func CreateUser(user User) error{
	users=append(users, user)
	return nil
}


func ListCounties ()[]Country{
	return countries
}