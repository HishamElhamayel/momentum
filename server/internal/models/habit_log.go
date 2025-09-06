package models

import "time"

type HabitLog struct {
	ID uint
	HabitID uint `gorm:"not null"`
	Date time.Time `gorm:"not null"`
	CreatedAt time.Time
}