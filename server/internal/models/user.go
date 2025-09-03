package models

import "time"

type User struct {
	ID uint `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"not null"`
	Email string `gorm:"not null;unique"`
	Password string `gorm:"not null" json:"-"`
	CreatedAt time.Time

	Habits []Habit
}

