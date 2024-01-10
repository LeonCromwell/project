package storage

import (
	"context"
	"example/auth-services/model"
)

func (s *sqlStorage) CreateNewUser(ctx context.Context, data *model.UserInput) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
