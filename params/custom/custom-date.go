package custom

import (
	"time"

	"github.com/go-playground/validator/v10"
)

func ValidateCustomDate(fl validator.FieldLevel) bool {
	// Extract input
	date := fl.Field().String()

	_, err := time.Parse("2024-01-01", date)
	return err == nil
}
