package models

import (
	"gorm.io/gorm"
	"time"
)

type MODEL struct {
	ID       int64 `json:"id" gorm:"primarykey"`
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt gorm.DeletedAt `gorm:"index"`
}
