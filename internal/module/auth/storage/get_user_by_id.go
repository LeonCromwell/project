package storage

import (
	"context"
	"example/auth-services/model"
)

func (s *sqlStorage) GetUserByID(ctx context.Context, id int) (*model.UserOutPut, error) {
	var user model.UserOutPut
	if err := s.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
