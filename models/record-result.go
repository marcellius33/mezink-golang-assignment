package models

import "time"

type RecordResult struct {
	ID         uint
	CreatedAt  time.Time
	TotalMarks uint
}
