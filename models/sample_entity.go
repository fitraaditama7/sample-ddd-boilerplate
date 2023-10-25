package models

import "time"

type Sample struct {
	ID        int64      `gorm:"column:id;primaryKey"`
	Name      string     `gorm:"column:name"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}
