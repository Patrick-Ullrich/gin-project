package domain

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	Email     string    `json:"email" gorm:"type:varchar(100);not null;unique;"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null;"`
	FirstName string    `json:"first_name" gorm:"type:varchar(100);"`
	LastName  string    `json:"last_name" gorm:"type:varchar(100);"`
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	userId, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	u.ID = userId
	u.Password = string(hashedPassword)

	return nil
}
