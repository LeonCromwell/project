package storage

import "example/auth-services/model"

func (s *sqlStorage) UpdateUser(id int, user model.UpdateUserInput) (model.UpdateUserInput, error) {
	var u model.UpdateUserInput
	if err := s.db.Model(&u).Where("id = ?", id).Updates(user).Error; err != nil {
		return u, err
	}
	return u, nil
}