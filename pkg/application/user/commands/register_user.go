package user

import (
	"errors"
	"example/gin-project/pkg/application/common/exceptions"
	"example/gin-project/pkg/domain"

	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type RegisterUser struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
}

func RegisterUserHandle(db *gorm.DB, r RegisterUser) (domain.User, error) {
	u := domain.User{
		Email:     r.Email,
		Password:  r.Password,
		FirstName: r.FirstName,
		LastName:  r.LastName,
	}

	if result := db.Create(&u); result.Error != nil {
		var pgErr *pgconn.PgError

		if errors.As(result.Error, &pgErr) {
			if pgErr.Code == "23505" {
				return domain.User{}, &exceptions.NotUniqueException{Field: "email", Message: "Email already registered"}
			}
		}
		return domain.User{}, result.Error
	}

	return u, nil
}
