package controllers

import (
	"errors"
	"example/gin-project/pkg/utils"
	"fmt"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fe.Field())
	}
	return "Unknown error"
}

func buildValidationError(err error) []ErrorMsg {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ErrorMsg, len(ve))

		for i, fe := range ve {
			out[i] = ErrorMsg{utils.FirstToLower(fe.Field()), getErrorMsg(fe)}
		}
		return out
	}
	return nil
}

type base struct {
	DB *gorm.DB
}

func New(db *gorm.DB) base {
	return base{db}
}
