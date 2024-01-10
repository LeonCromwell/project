package business

import (
	"context"

	"example/auth-services/model"
)

type Register interface {
	CreateNewUser(ctx context.Context, data *model.UserInput) error
}

type registerBusiness struct {
	register Register
}

func RegisterBusiness(register Register) *registerBusiness {
	return &registerBusiness{register: register}
}

func (r *registerBusiness) RegisterBusiness(ctx context.Context, data)