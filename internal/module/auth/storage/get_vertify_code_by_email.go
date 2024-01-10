package storage

import (
	"context"
	"example/auth-services/model"
)

func (s *sqlStorage) GetVertifyCodeByEmail(ctx context.Context, email string) (vertifyCode *string, err error) {
    var data model.VertifyInput

    if err := s.db.Where("email = ?", email).First(&data).Error; err != nil {
        return nil, err
    }

    return &data.Code, nil
}
