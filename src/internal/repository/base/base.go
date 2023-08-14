package repository_base

import "time"

type BaseRepository struct {
	ID        int64      `json:"id"         gorm:"primaryKey;autoIncrement"`
	CreatedAt *time.Time `json:"created_at" gorm:"not null;autoCreateTime:nano"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"not null;autoUpdateTime:nano"`
}
