package repositories

import (
	"mezink-golang-assignment/models"

	"gorm.io/gorm"
)

type RecordRepository interface {
	CreateRecord(record *models.Record) (*models.Record, error)
	GetRecords(records *[]models.RecordResult, startDate string, endDate string, minCount int, maxCount int) (*[]models.RecordResult, error)
}

type recordRepository struct {
	db *gorm.DB
}

func NewRecordRepository(db *gorm.DB) RecordRepository {
	return &recordRepository{
		db: db,
	}
}

func (r recordRepository) CreateRecord(record *models.Record) (*models.Record, error) {
	return record, r.db.Create(record).Error
}

func (r recordRepository) GetRecords(records *[]models.RecordResult, startDate string, endDate string, minCount int, maxCount int) (*[]models.RecordResult, error) {
	err := r.db.Raw(`
		SELECT
			id,
			created_at,
			sum(cast(unnested_marks as integer)) as total_marks
		FROM
			records,
			unnest(string_to_array(regexp_replace(marks, '^\[|\]$', '', 'g'), ',')) AS unnested_marks
		WHERE
			created_at BETWEEN ? AND ?
		GROUP BY
			id
		HAVING
			SUM(cast(unnested_marks as integer)) BETWEEN ? AND ?
		ORDER BY 
			id 
	`, startDate, endDate, minCount, maxCount).Scan(&records).Error

	return records, err
}
