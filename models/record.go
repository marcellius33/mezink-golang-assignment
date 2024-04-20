package models

import "time"

type Record struct {
	ID        uint `gorm:"primaryKey;type:integer"`
	Name      string
	Marks     []uint `gorm:"json"`
	CreatedAt time.Time
}
