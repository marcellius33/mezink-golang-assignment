package controllers

import (
	"mezink-golang-assignment/helpers"
	"mezink-golang-assignment/params"
	"mezink-golang-assignment/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RecordController struct {
	service services.RecordService
}

func NewRecordController(service services.RecordService) *RecordController {
	return &RecordController{
		service: service,
	}
}

func (r *RecordController) GetRecords(c *gin.Context) {
	recordRequest := params.GetRecordRequest{}

	// Bind request body to struct
	if err := c.ShouldBindJSON(&recordRequest); err != nil {
		helpers.WriteJsonResponse(c, helpers.BadRequestResponse(err))
		return
	}

	// Validate input struct
	validate := validator.New()
	if err := validate.Struct(recordRequest); err != nil {
		helpers.WriteJsonResponse(c, helpers.BadRequestResponse(err))
		return
	}

	records, err := r.service.GetRecords(recordRequest)
	if err != nil {
		helpers.WriteJsonResponse(c, helpers.InternalServerError(err))
		return
	}

	helpers.WriteJsonResponse(c, helpers.SuccessResponse(records))
}
