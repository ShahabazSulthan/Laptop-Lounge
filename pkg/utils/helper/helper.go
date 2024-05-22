package helper

import (
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
	resCustomError "Laptop_Lounge/pkg/models/responseModel/custom_error"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func Pagination(page string, limit string) (int, int, error) {

	fmt.Println("jj", page, limit)

	pageNO, err := strconv.Atoi(page)
	if err != nil {
		return 0, 0, resCustomError.ErrConversionOFPage
	}

	if pageNO < 1 {
		return 0, 0, resCustomError.ErrConversionOFPage
	}

	limits, err := strconv.Atoi(limit)
	if err != nil {
		return 0, 0, resCustomError.ErrConversionOfLimit
	}

	if limits <= 0 {
		return 0, 0, resCustomError.ErrPageLimit
	}

	offSet := (pageNO * limits) - limits
	limits = pageNO * limits

	return offSet, limits, nil
}

func Validation(data interface{}) ([]responsemodel.Errors, error) {
	var validationErrors []responsemodel.Errors
	validate := validator.New()

	// Register custom date validation function
	if err := validate.RegisterValidation("date", validateDate); err != nil {
		return validationErrors, err
	}

	err := validate.Struct(data)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			errorMessages := map[string]string{
				"required":         "%s is required",
				"min":              "%s should be at least %s",
				"max":              "%s should be at most %s",
				"email":            "%s should be in email format",
				"eqfield":          "%s should be equal to %s",
				"len":              "%s should have length %s",
				"alpha":            "%s should be alphabetic",
				"number":           "%s should be numeric",
				"numeric":          "%s should be numeric",
				"uppercase":        "%s should be uppercase",
				"gtcsfield":        "%s should be greater than %s",
				"date.required":    "%s is required",
				"date.date":        "%s should be a valid date (YYYY-MM-DD)",
				"date.layout":      "Invalid date format provided",
				"date.lesserequal": "%s should be lesser or equal to ExpireDate",
			}

			for _, e := range ve {
				if msg, exists := errorMessages[e.Tag()]; exists {
					errMsg := fmt.Sprintf(msg, e.Field(), e.Param())
					validationErrors = append(validationErrors, responsemodel.Errors{Err: errMsg})
				}
			}
		}
		return validationErrors, errors.New("validation failed")
	}

	return validationErrors, nil
}

// ValidateDate checks if the provided date is in the correct format and is not in the past
func validateDate(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()
	layout := "2006-01-02"

	// Check if date is in the correct format
	if _, err := time.Parse(layout, dateStr); err != nil {
		return false
	}

	// Check if date is not in the past
	date, _ := time.Parse(layout, dateStr)
	return date.After(time.Now().AddDate(0, 0, -1))
}

func GenerateUUID() string {
	newUUID := uuid.New()

	uuidString := newUUID.String()
	return uuidString
}

func StringToUintConvertion(value string) (uint, error) {

	result, err := strconv.Atoi(value)
	if err != nil {
		return 0, errors.New("convertion lead to error")
	}
	return uint(result), nil
}

func FindDiscount(originalPrice, percentageOffer float64) uint {

	discountAmount := (percentageOffer / 100) * originalPrice
	discountPrice := originalPrice - discountAmount
	return uint(discountPrice)

}

func GenerateReferalCode() (string, error) {
	bytes := make([]byte, 10)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", nil
	}

	return base64.URLEncoding.EncodeToString(bytes)[:5], nil
}
