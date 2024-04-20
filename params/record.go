package params

import (
	"mezink-golang-assignment/models"
	"time"
)

type GetRecordRequest struct {
	StartDate string `json:"startDate" validate:"required"`
	EndDate   string `json:"endDate" validate:"required"`
	MinCount  uint   `json:"minCount" validate:"required,number"`
	MaxCount  uint   `json:"maxCount" validate:"required,number,gtfield=MinCount"`
}

type GetRecordResponse struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalMarks uint      `json:"totalMarks"`
}

func ParseToGetRecordResponse(records []models.RecordResult) []GetRecordResponse {
	var results []GetRecordResponse
	for _, record := range records {
		results = append(results, GetRecordResponse{
			ID:         record.ID,
			CreatedAt:  record.CreatedAt,
			TotalMarks: record.TotalMarks,
		})
	}

	return results
}
