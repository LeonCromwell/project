package business

import (
	"context"
	"example/auth-services/model"
)

type Getuser interface {
	GetUserByID(ctx context.Context, id int) (*model.UserOutPut, error)
}

type getuserbussiness struct {
	store Getuser
}

func GetuserBusiness(store Getuser) *getuserbussiness {
	return &getuserbussiness{store: store}
}

func (g *getuserbussiness) GetUser(ctx context.Context, id int) (*model.UserOutPut, error) {
	user, err := g.store.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}