package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/makersacademy/go-react-acebook-template/api/src/auth"
	"github.com/makersacademy/go-react-acebook-template/api/src/models"
)

type JSONPost struct {
	ID      uint   `json:"_id"`
	Message string `json:"message"`
	// CreatedAt time.Time `json:"created_at"`
	// add fields that would be needed here, important to comm
	// this to FE
}

func GetAllPosts(ctx *gin.Context) {
	posts, err := models.FetchAllPosts()

	if err != nil {
		SendInternalError(ctx, err)
		return
	}

	val, _ := ctx.Get("userID")
	userID := val.(string)
	token, _ := auth.GenerateToken(userID)

	// Convert posts to JSON Structs
	jsonPosts := make([]JSONPost, 0)
	for _, post := range *posts {
		jsonPosts = append(jsonPosts, JSONPost{
			Message: post.Message,
			ID:      post.ID,
			// CreatedAt: post.CreatedAt,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"posts": jsonPosts, "token": token})
}

type createPostRequestBody struct {
	Message   string
	CreatedAt time.Time
}

func CreatePost(ctx *gin.Context) {
	var requestBody createPostRequestBody
	err := ctx.BindJSON(&requestBody)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	if len(requestBody.Message) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Post message empty"})
		return
	}

	PostTime := time.Now()
	formattedTime := PostTime.Format("2006-01-02 15:04:05")

	newPost := models.Post{
		Message:   requestBody.Message,
		CreatedAt: fmt.Println,
	}

	_, err = newPost.Save()
	if err != nil {
		SendInternalError(ctx, err)
		return
	}

	val, _ := ctx.Get("userID")
	userID := val.(string)
	token, _ := auth.GenerateToken(userID)

	ctx.JSON(http.StatusCreated, gin.H{"message": "Post created", "token": token})
}
