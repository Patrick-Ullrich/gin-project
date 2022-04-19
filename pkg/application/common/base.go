package application

import "gorm.io/gorm"

type base struct {
	DB *gorm.DB
}

// Using this to inject a shared DB connection to all handlers
func New(db *gorm.DB) base {
	return base{db}
}
