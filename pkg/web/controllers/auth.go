package controllers

import (
	"errors"
	"net/http"

	"example/gin-project/pkg/application/common/exceptions"
	user "example/gin-project/pkg/application/user/commands"

	"github.com/gin-gonic/gin"
)

func (b base) Register(c *gin.Context) {
	var input user.RegisterUser

	if err := c.ShouldBindJSON(&input); err != nil {
		out := buildValidationError(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		return
	}

	token, err := user.RegisterUserHandle(b.DB, input)
	if err != nil {
		var nue *exceptions.NotUniqueException
		if errors.As(err, &nue) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": []ErrorMsg{{nue.Field, nue.Message}}})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": token})

}

func (b base) Login(c *gin.Context) {
	var body user.LoginCredentials

	if err := c.BindJSON(&body); err != nil {
		out := buildValidationError(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		return
	}

	token, err := user.LoginHandle(b.DB, body)
	if err != nil {
		var nue *exceptions.InvalidCredentialsException
		if errors.As(err, &nue) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, nue.Error())
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
