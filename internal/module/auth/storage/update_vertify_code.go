package storage

import (
	"context"
	"example/auth-services/model"
)

func (s *sqlStorage) UpdateVertifyCode(ctx context.Context, data *model.VertifyInput) error {
	var result model.Vertify
	if err := s.db.Where("email = ?", data.Email).First(&result).Error; err != nil {
		return err
	}

	result.Code = data.Code
	if err := s.db.Save(&result).Error; err != nil {
		return err
	}
	return nil
}
