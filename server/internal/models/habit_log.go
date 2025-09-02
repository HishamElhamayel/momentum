package models

import "time"

type HabitLog struct {
	ID uint
	HabitID uint `gorm:"not null"`
	Date time.Time `gorm:"not null"`
	Completed bool `gorm:"default:false"`
	CreatedAt time.Time
}