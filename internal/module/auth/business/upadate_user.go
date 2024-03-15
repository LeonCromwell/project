package business

import "example/auth-services/model"

type UpdateUser interface {
	UpdateUser(id int, user model.UpdateUserInput) (model.UpdateUserInput, error)
}

type updateUser struct {
	storage UpdateUser
}

func UpdateUserBusiness(storage UpdateUser) *updateUser {
	return &updateUser{storage: storage}
}

func (u *updateUser) UpdateUserBusiness(id int, user model.UpdateUserInput) (model.UpdateUserInput, error) {
	return u.storage.UpdateUser(id, user)
}