package controllers

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/makersacademy/go-react-acebook-template/api/src/models"
)

func CreateUser(ctx *gin.Context) {
	var newUser models.User
	// err := ctx.ShouldBindJSON(&newUser)

	// ERROR HANDLING for ShouldBindJSON below

	// if err != nil {
	// 	if jsonErr, ok := err.(*json.UnmarshalTypeError); ok {
	// 		fieldName := jsonErr.Field
	// 		errorMsg := fmt.Sprintf("Invalid value for field '%s': %v", fieldName, jsonErr.Error())
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": errorMsg})
	// 	} else {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	}
	// 	return
	// }

	// The below block reads the image data from the request where
	// the content-type is set to multipart/form-data (in Headers)

	file, header, err := ctx.Request.FormFile("image")
	// image is the key in the Postman form-data POST request
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()
	// Read file data
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newUser = models.User{
		// Update user fields with file information
		Email:    ctx.PostForm("email"),
		Password: ctx.PostForm("password"),
		Username: ctx.PostForm("username"),
		Filename: header.Filename,
		FileSize: header.Size,
		FileType: header.Header.Get("Content-Type"),
		FileData: fileBytes,
	}

	if newUser.Email == "" || newUser.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Must supply username and password"})
		return
	}

	_, err = newUser.Save()
	if err != nil {
		SendInternalError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "OK"})
}
