package business

import (
	"context"
	"errors"
	"example/auth-services/common"
	"example/auth-services/internal/pkg/auth"
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

func (l *loginBusiness) Login(ctx context.Context, data *model.UserLoginInput) (token *string,refreshToken *string, err error) {
	user, err := l.storage.GetUserByEmail(ctx, data.Email)
	if err != nil {
		return nil,nil, err
	}

	if user == nil {
		return nil,nil, nil
	}

	if !common.CheckPassword(data.Hashpassword, user.Hashpassword) {
		return nil,nil, errors.New("Password is wrong")
	}

	accessTokenString, err1 := auth.SignAccessToken(user.ID)
	refreshTokenString, err2 := auth.SignRefreshToken(user.ID)
	if err1 != nil {
		return nil,nil, err1
	}

	if err2 != nil {
		return nil,nil, err2
	}

	return &accessTokenString, &refreshTokenString, nil
}