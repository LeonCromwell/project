package model

import "example/auth-services/common"

type User struct {
	common.Sqlmodel
	Name         string `json:"name"`
	Email        string `json:"email"`
	Hashpassword string `json:"hashpassword"`
}

func (User) TableName() string {
	return "user"
}

type UserInput struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Hashpassword string `json:"hashpassword"`
}

func (UserInput) TableName() string {
	return User{}.TableName()
}

type UserLoginInput struct {
	Email        string `json:"email"`
	Hashpassword string `json:"hashpassword"`
}

func (UserLoginInput) TableName() string {
	return User{}.TableName()
}