package storage

import (
	"context"
	"example/auth-services/model"
)

func (s *sqlStorage) CreateNewCode(ctx context.Context, data *model.VertifyInput) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
