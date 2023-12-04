package utils

import (
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func GenerateID() string {
	return uuid.New().String()
}

type IError struct {
	Field string
	Tag   string
	Value string
}

func GetStructError(err error) []IError {
	var errors []IError

	if err == nil {
		return errors
	}

	// if _, ok := err.(*validator.InvalidValidationError); ok {
	for _, err := range err.(validator.ValidationErrors) {
		errors = append(errors, IError{
			Field: err.Field(),
			Tag:   err.Tag(),
			Value: err.Value().(string),
		})
	}
	// }

	return errors
}

func GenerateSlug(s string) string {
	slug := strings.ReplaceAll(s, " ", "-")
	slug = strings.ToLower(slug)
	return slug
}

func GenerateNumber(date time.Time) string {
	// return date formatted as a string with the following format: YYYYMMDDHHMMSS
	return date.Format("20060102150405")
}
