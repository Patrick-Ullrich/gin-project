package user

import (
	"errors"
	"example/gin-project/pkg/application/common/exceptions"
	"example/gin-project/pkg/domain"
	"example/gin-project/pkg/infrastructure/services/jwt_service"
	"log"

	"github.com/jackc/pgconn"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginCredentials struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandle(db *gorm.DB, lc LoginCredentials) (string, error) {

	u := domain.User{}

	err := db.Model(domain.User{}).Where("email = ?", lc.Email).Take(&u).Error
	if err != nil {
		var pgErr *pgconn.PgError
		log.Println("NOTFOUND: ", err)
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return "", &exceptions.InvalidCredentialsException{}
			}
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(lc.Password))

	if err != nil {
		return "", &exceptions.InvalidCredentialsException{}
	}

	token, err := jwt_service.GenerateJwt(u.ID, u.Email, u.FirstName, u.LastName)

	if err != nil {
		return "", err
	}

	return token, nil
}
