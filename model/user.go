package model

import (
	"example/auth-services/common"
)

type User struct {
	common.Sqlmodel
	Name         string `json:"name"`
	Email        string `json:"email"`
	Hashpassword string `json:"hashpassword"`
	Authenticated int `json:"authenticated" gorm:"default:0" ` 
	Role         string `json:"role" gorm:"default:'user'"`
	Avartar	  string `json:"avartar"`	
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

type UserOutPut struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Role         string `json:"role" gorm:"default:'user'"`
	Avartar	  string `json:"avartar"`
}

func (UserOutPut) TableName() string {
	return User{}.TableName()
}

type UpdateUserInput struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Avartar	  string `json:"avartar"`
}

func (UpdateUserInput) TableName() string {
	return User{}.TableName()
}