package storage

import (
	"context"
	"example/auth-services/model"
)

func (s *sqlStorage) CreateToken(ctx context.Context, data *model.Token) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}