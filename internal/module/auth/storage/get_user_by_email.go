package storage

import (
	"context"
	"example/auth-services/model"
)

func (s *sqlStorage) GetUserByEmail(ctx context.Context, email string) (user *model.User, err error) {
	var user1 model.User

	// Sửa truy vấn SQL để bao quanh giá trị email bằng dấu nháy đơn ('')
	if err := s.db.Where("email = ?", email).First(&user1).Error; err != nil {
		return nil, err
	}

	// Trả về ID của người dùng nếu tìm thấy
	return &user1, nil
}
