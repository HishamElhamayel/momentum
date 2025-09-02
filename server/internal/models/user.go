package models

import "time"

type User struct {
	ID uint 
	Name string `grom:"size:100;not null"`
	Email string `gorm:"size:100; not null; uniqIndex"`
	Password string `gorm:"not null" json:"-"`
	CreatedAt time.Time

	Habits []Habit
}

