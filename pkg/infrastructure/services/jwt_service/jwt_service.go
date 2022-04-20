package jwt_service

import (
	"log"
	"os"
	"strconv"
	"time"

	jwt "github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func GenerateJwt(uId uuid.UUID, email string, firstName string, lastName string) (string, error) {

	token_lifespan_string := os.Getenv("TOKEN_LIFESPAN_HOURS")
	jwt_secret := os.Getenv("JWT_SECRET")

	if token_lifespan_string == "" {
		log.Fatalln("Missing ENV variable <TOKEN_LIFESPAN_HOURS>")
	}

	if jwt_secret == "" {
		log.Fatalln("Missing ENV variable <JWT_SECRET>")
	}

	token_lifespan, err := strconv.Atoi(token_lifespan_string)

	if err != nil {
		log.Fatalln("Error reading <TOKEN_LIFESPAN_HOURS> env variable", err.Error())
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        uId,
		"email":     email,
		"firstName": firstName,
		"lastName":  lastName,
		"exp":       time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix(),
	})

	return token.SignedString([]byte(jwt_secret))
}
