package seed

import (
	"mezink-golang-assignment/models"
	"mezink-golang-assignment/repositories"
	"time"
)

type RecordSeeder struct {
	repository repositories.RecordRepository
}

func NewRecordSeeder(repository repositories.RecordRepository) *RecordSeeder {
	return &RecordSeeder{
		repository: repository,
	}
}

func (r *RecordSeeder) Execute() {
	createdAts := []time.Time{
		time.Now().Add(time.Hour * -24),
		time.Now().Add(time.Hour * 5),
		time.Now().Add(time.Hour * -48),
		time.Now().Add(time.Hour * -47),
		time.Now().Add(time.Hour * 120),
		time.Now().Add(time.Hour * 300),
	}

	records := []models.Record{
		{
			Name:      "Curry",
			Marks:     []uint{100, 20, 30, 50},
			CreatedAt: createdAts[0],
		},
		{
			Name:      "Stephen",
			Marks:     []uint{30, 50, 90, 54},
			CreatedAt: createdAts[1],
		},
		{
			Name:      "James",
			Marks:     []uint{80, 70, 30, 20},
			CreatedAt: createdAts[2],
		},
		{
			Name:      "Jimmy",
			Marks:     []uint{10, 20, 30, 40},
			CreatedAt: createdAts[3],
		},
		{
			Name:      "Kelly",
			Marks:     []uint{30, 30, 30, 80},
			CreatedAt: createdAts[4],
		},
		{
			Name:      "Coby",
			Marks:     []uint{30, 50, 100, 40},
			CreatedAt: createdAts[5],
		},
	}

	for _, record := range records {
		r.repository.CreateRecord(&record)
	}
}
