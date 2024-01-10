package common

type Sqlmodel struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	CreatedAt string `json:"created_at"`
}
