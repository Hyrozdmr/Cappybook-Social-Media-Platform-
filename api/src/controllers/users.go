package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/makersacademy/go-react-acebook-template/api/src/models"
)

func CreateUser(ctx *gin.Context) { // ctx refers to the context of the incoming HTTP request
	var newUser models.User       // Creates a variable called newUser with the User struct type User{gorm.Model(id,...), email, password}
	err := ctx.BindJSON(&newUser) // Parses the JSON from the request and attempts to match the fields to the newUser fields

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if newUser.Email == "" || newUser.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Must supply username and password"}) // Returns error if email and password are blank
		return
	}

	_, err = newUser.Save() // Adds newUser to database
	if err != nil {
		SendInternalError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "OK"}) //sends confirmation message back if successfully saved
}
