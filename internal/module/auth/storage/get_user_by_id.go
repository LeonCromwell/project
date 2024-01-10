package storage

import (
	"context"
	"example/auth-services/model"
)

func (s *sqlStorage) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	var user model.User
	if err := s.db.Where("id = ?", id).First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
