package models

import (
	"gorm.io/gorm"
	"time"
)

type MODEL struct {
	ID        int64 `json:"id" gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
