package storage

import "gorm.io/gorm"

type sqlStorage struct {
	db *gorm.DB
}



func NewStorage(db *gorm.DB) *sqlStorage {
	return &sqlStorage{db: db}
}
