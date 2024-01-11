package storage

import "context"

func (s *sqlStorage) SetStateAuth(ctx context.Context,email string, state int) error {
	if err := s.db.Exec("UPDATE user SET authenticated = ? Where email = ?", state, email).Error; err != nil {
		return err
	}
	return nil
}
