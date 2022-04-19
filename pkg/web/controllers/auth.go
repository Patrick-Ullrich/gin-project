package controllers

import (
	"errors"
	"log"
	"net/http"

	"example/gin-project/pkg/application/common/exceptions"
	user "example/gin-project/pkg/application/user/commands"
	"example/gin-project/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (b base) Register(c *gin.Context) {
	var input user.RegisterUser
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("%v", err)
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))

			for i, fe := range ve {
				out[i] = ErrorMsg{utils.FirstToLower(fe.Field()), getErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	user, err := user.RegisterUserHandle(b.DB, input)
	if err != nil {
		var nue *exceptions.NotUniqueException
		if errors.As(err, &nue) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": []ErrorMsg{{nue.Field, nue.Message}}})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, user)

}

func (b base) Login(c *gin.Context) {
	// var input AuthInput

	// if err := c.BindJSON(&input); err != nil {
	// 	fmt.Println(err)
	// 	c.IndentedJSON(http.StatusBadRequest, gin.H{"Invalid input": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"message": "Login success"})
}
