package model

import "example/auth-services/common"

type Vertify struct {
	common.Sqlmodel
	Email string `json:"email"`
	Code  string `json:"code"`
}

func (Vertify) TableName() string {
	return "verification_codes"
}

type VertifyInput struct {
	Email string `json:"email"`
	Code string `json:"code"`
}

func (VertifyInput) TableName() string {
	return Vertify{}.TableName()
}
