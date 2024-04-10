package helper

import (
	responsemodel "Laptop_Lounge/pkg/models/responseModel"
	resCustomError "Laptop_Lounge/pkg/models/responseModel/custom_error"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func Pagination(page string, limit string) (int, int, error) {

	pageNO, err := strconv.Atoi(page)
	if err != nil {
		return 0, 0, resCustomError.ErrConversionOfPage
	}

	if pageNO < 1 {
		return 0, 0, resCustomError.ErrConversionOfPage
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
	var afterErrorCorrection []responsemodel.Errors
	validate := validator.New()

	err := validate.Struct(data)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			errorMessages := map[string]string{
				"required":  "%s is required",
				"min":       "%s should be at least %s",
				"max":       "%s should be at most %s",
				"email":     "%s should be in email format",
				"eqfield":   "%s should be equal to %s",
				"len":       "%s should have length %s",
				"alpha":     "%s should be alphabetic",
				"number":    "%s should be numeric",
				"numeric":   "%s should be numeric",
				"uppercase": "%s should be uppercase",
				"gtcsfield": "%s should be greater than %s",
			}

			for _, e := range ve {
				if msg, exists := errorMessages[e.Tag()]; exists {
					errMsg := fmt.Sprintf(msg, e.Field(), e.Param())
					afterErrorCorrection = append(afterErrorCorrection, responsemodel.Errors{Err: errMsg})
				}
			}
		}
		return afterErrorCorrection, errors.New("doesn't fulfill the requirements")
	}
	return afterErrorCorrection, nil
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
