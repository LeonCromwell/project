package business

import (
	"context"
	"errors"
	"example/auth-services/common"
	"example/auth-services/model"
)

type Login interface {
	GetUserByEmail(ctx context.Context, email string) (user *model.User, err error)
	
}

type loginBusiness struct {
	storage Login
}

func LoginBusiness(storage Login) *loginBusiness {
	return &loginBusiness{storage: storage}
}

func (l *loginBusiness) Login(ctx context.Context, data *model.UserLoginInput) (token *string, err error) {
	user, err := l.storage.GetUserByEmail(ctx, data.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	if !common.CheckPassword(data.Hashpassword, user.Hashpassword) {
		return nil, errors.New("Password is wrong")
	}

	
	return nil, nil
}