package storage

import (
	"context"
	"example/auth-services/model"
)

func (s *sqlStorage) GetUserByEmail(ctx context.Context, email string) (id *int, err error) {
	var user model.User

	// Sửa truy vấn SQL để bao quanh giá trị email bằng dấu nháy đơn ('')
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	// Trả về ID của người dùng nếu tìm thấy
	return &user.ID, nil
}
