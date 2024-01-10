package business

import (
	"context"
	"errors"
	"example/auth-services/model"
)

type Vertify interface {
	GetVertifyCodeByEmail(ctx context.Context, email string) (vertifyCode *string, err error)
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

	if data.Code != *vertifyCode {
		return errors.New("Wrong code")
	}

	return nil
}
