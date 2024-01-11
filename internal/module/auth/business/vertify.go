package business

import (
	"context"
	"errors"
	"example/auth-services/internal/pkg/auth"
	"example/auth-services/model"
)

type Vertify interface {
	GetVertifyCodeByEmail(ctx context.Context, email string) (vertifyCode *string, err error)
	SetStateAuth(ctx context.Context, email string, state int) error
}

type vertifyBusiness struct {
	storage Vertify
}

func VertifyBusiness(storage Vertify) *vertifyBusiness {
	return &vertifyBusiness{storage: storage}
}
func (v *vertifyBusiness) VertifyBusiness(ctx context.Context, data *model.VertifyInput) error {
	vertifyCode, err := v.storage.GetVertifyCodeByEmail(ctx, data.Email)
	if err != nil {
		return err
	}

	if !auth.VerifyCode(data.Code, *vertifyCode) {
		return errors.New("Wrong code")
	}

	if err := v.storage.SetStateAuth(ctx, data.Email, 1); err != nil {
		return err
	}
	return nil
}
