package model

import "example/auth-services/common"

type Token struct {
	common.Sqlmodel
	UserID       int    `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiryTime   string `json:"expiry_time"`
}

func (Token) TableName() string {
	return "tokens"
}
