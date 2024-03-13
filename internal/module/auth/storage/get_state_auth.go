package storage

import "context"


func (s *sqlStorage) GetStateAuth(ctx context.Context,email string) (int, error) {
	var state int
	if err := s.db.Raw("SELECT authenticated FROM user WHERE email = ?", email).Scan(&state).Error; err != nil {
		return 0, err
	}
	return state, nil
}