package user

import (
	"errors"
	"example/gin-project/pkg/application/common/exceptions"
	"example/gin-project/pkg/domain"
	"example/gin-project/pkg/infrastructure/services/jwt_service"

	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type RegisterUser struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
}

func RegisterUserHandle(db *gorm.DB, r RegisterUser) (string, error) {
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
				return "", &exceptions.NotUniqueException{Field: "email", Message: "Email already registered"}
			}
		}
		return "", result.Error
	}

	token, err := jwt_service.GenerateJwt(u.ID, u.Email, u.FirstName, u.LastName)
	if err != nil {
		return "", err
	}

	return token, nil
}
