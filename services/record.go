package services

import (
	"mezink-golang-assignment/models"
	"mezink-golang-assignment/params"
	"mezink-golang-assignment/repositories"
)

type RecordService interface {
	GetRecords(getRecordRequest params.GetRecordRequest) (*[]params.GetRecordResponse, error)
}

type recordService struct {
	repository repositories.RecordRepository
}

func NewRecordService(repository repositories.RecordRepository) RecordService {
	return &recordService{
		repository: repository,
	}
}

func (r *recordService) GetRecords(getRecordRequest params.GetRecordRequest) (*[]params.GetRecordResponse, error) {
	var records []models.RecordResult
	_, err := r.repository.GetRecords(&records, getRecordRequest.StartDate, getRecordRequest.EndDate, int(getRecordRequest.MinCount), int(getRecordRequest.MaxCount))

	if err != nil {
		return &[]params.GetRecordResponse{}, err
	}

	resp := params.ParseToGetRecordResponse(records)

	return &resp, nil
}
