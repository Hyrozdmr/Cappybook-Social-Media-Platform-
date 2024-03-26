package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/makersacademy/go-react-acebook-template/api/src/auth"
	"github.com/makersacademy/go-react-acebook-template/api/src/models"
)

type JSONPost struct {
	ID      uint   `json:"_id"` // Relates to the primary key ID for the posts table
	Message string `json:"message"`
}

func GetAllPosts(ctx *gin.Context) {
	// FetchAllPosts() returns all posts from the database in a slice
	posts, err := models.FetchAllPosts()

	if err != nil {
		SendInternalError(ctx, err)
		return
	}

	//Following three lines retrieve the userID in order to generate a new auth token
	val, _ := ctx.Get("userID")
	userID := val.(string)
	token, _ := auth.GenerateToken(userID)

	// Convert posts to JSON Structs
	jsonPosts := make([]JSONPost, 0)
	for _, post := range *posts {
		jsonPosts = append(jsonPosts, JSONPost{
			Message: post.Message,
			ID:      post.ID,
		})
	}

	// Sends a JSON response with a status code of 200 (http.StatusOK) containing the posts and generated token
	ctx.JSON(http.StatusOK, gin.H{"posts": jsonPosts, "token": token})
}

type createPostRequestBody struct {
	Message string
	// ** UNSURE HOW THIS MATCHES TO LOWERCASE "message" JSON FIELD **
}

func CreatePost(ctx *gin.Context) {
	var requestBody createPostRequestBody
	err := ctx.BindJSON(&requestBody)
	// ctx.BindJSON reads the JSON payload from the request body (frontend/src/services/posts.js)
	// it parses the JSON payload and attempts to match the JSON fields with the fields in the requestBody struct
	// if the JSON payload has a field named "message" it assigns the corresponding value to the Message field of the requestBody

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	if len(requestBody.Message) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Post message empty"})
		return
	}

	newPost := models.Post{
		Message: requestBody.Message,
	}

	_, err = newPost.Save()
	if err != nil {
		SendInternalError(ctx, err)
		return
	}

	val, _ := ctx.Get("userID")
	userID := val.(string)
	token, _ := auth.GenerateToken(userID)

	ctx.JSON(http.StatusCreated, gin.H{"message": "Post created", "token": token}) //sends confirmation message back if successfully saved
}
