package business

import (
	"context"
	"errors"

	"example/auth-services/common"
	"example/auth-services/model"

	"gorm.io/gorm"
)

type Register interface {
	GetUserByEmail(ctx context.Context, email string) (user *model.User, err error)
	CreateNewUser(ctx context.Context, data *model.UserInput) error
}

type registerBusiness struct {
	storage Register
}

func RegisterBusiness(storage Register) *registerBusiness {
	return &registerBusiness{storage: storage}
}

func (r *registerBusiness) RegisterBusiness(ctx context.Context, data *model.UserInput) error {
	user1, err := r.storage.GetUserByEmail(ctx, data.Email)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	if user1 != nil {
		return errors.New("Email already exist")
	}

	password := data.Hashpassword
	var newHashPassword = common.HashPassword(password)
	data.Hashpassword = newHashPassword
	if err := r.storage.CreateNewUser(ctx, data); err != nil {
		return err
	}

	return nil
}
