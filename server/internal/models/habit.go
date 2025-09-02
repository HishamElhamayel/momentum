package models

import "time"


type Habit struct {
	ID uint
	UserID uint `gorm:"not null"`
	Title string `gorm:"size:100; not null"`
	Description string `gorm:"size:255; not null"`
	CreatedAt time.Time

	Logs []HabitLog
}
