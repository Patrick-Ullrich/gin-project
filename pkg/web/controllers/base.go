package controllers

import (
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

type base struct {
	DB *gorm.DB
}

func New(db *gorm.DB) base {
	return base{db}
}
