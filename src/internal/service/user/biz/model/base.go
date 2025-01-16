package model

import "time"

type Base struct {
	ID        uint32    `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
