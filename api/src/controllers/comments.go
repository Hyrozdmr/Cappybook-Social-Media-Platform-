package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/makersacademy/go-react-acebook-template/api/src/auth"
	"github.com/makersacademy/go-react-acebook-template/api/src/models"
	"net/http"
	"strconv"
	"time"
)

type JSONComment struct {
	ID      uint   `json:"_id"`
	Message string `json:"message"`
	Likes   int    `json:"likes"`
	PostId  int    `json:"post_id"`
	UserID  string `json:"user_id"`
}

type createCommentRequestBody struct {
	Message string
}

func CreateComment(ctx *gin.Context) {
	var requestBody createCommentRequestBody
	err := ctx.BindJSON(&requestBody)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	if len(requestBody.Message) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Post message empty"})
		return
	}

	postID := ctx.Param("id")
	postIdInt, err := strconv.Atoi(postID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	PostTime := time.Now()

	userIDToken, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"ERROR": "USER ID NOT FOUND IN CONTEXT"})
		return
	}
	// in this function a token is not generated, therefore not returned as part of the JSON
	userIDString := userIDToken.(string)

	newComment := models.Comment{
		UserID:    strconv.Itoa(int([]byte(userIDString)[0])),
		Message:   requestBody.Message,
		CreatedAt: PostTime,
		Likes:     0,
		PostId:    postIdInt,
	}

	_, err = newComment.Save()
	if err != nil {
		SendInternalError(ctx, err)
		return
	}

	comments, err := models.FetchAllCommentsByPostId(postIdInt)

	if err != nil {
		SendInternalError(ctx, err)
		return
	}

	val, _ := ctx.Get("userID")
	userID := val.(string)
	token, _ := auth.GenerateToken(userID)

	// Convert comments to JSON Structs
	jsonComments := make([]JSONComment, 0)
	for _, comment := range *comments {
		jsonComments = append(jsonComments, JSONComment{
			Message: comment.Message,
			ID:      comment.ID,
			Likes:   comment.Likes,
			PostId:  comment.PostId,
			UserID:  comment.UserID,
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Comment created", "comments": jsonComments, "token": token})
}

func GetAllCommentsByPostId(ctx *gin.Context) {
	// retrieving a parameter from the request URL here below
	// needing route to be structured like this: /posts/:post_id
	postID := ctx.Param("id")

	postIdInt, err := strconv.Atoi(postID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	comments, err := models.FetchAllCommentsByPostId(postIdInt)

	if err != nil {
		SendInternalError(ctx, err)
		return
	}

	val, _ := ctx.Get("userID")
	userID := val.(string)
	token, _ := auth.GenerateToken(userID)

	// Convert comments to JSON Structs
	jsonComments := make([]JSONComment, 0)
	for _, comment := range *comments {
		jsonComments = append(jsonComments, JSONComment{
			Message: comment.Message,
			ID:      comment.ID,
			Likes:   comment.Likes,
			PostId:  comment.PostId,
			UserID:  comment.UserID,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"comments": jsonComments, "token": token})
}

func GetSpecificComment(ctx *gin.Context) {
	postID := ctx.Param("id")

	postIdInt, err := strconv.Atoi(postID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	commentID := ctx.Param("comment_id")

	commentIdInt, err := strconv.Atoi(commentID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	comment, err := models.FetchSpecificComment(postIdInt, commentIdInt)
	if err != nil {
		SendInternalError(ctx, err)
		return
	}

	jsonComment := JSONComment{
		Message: comment.Message,
		ID:      comment.ID,
		Likes:   comment.Likes,
		PostId:  comment.PostId,
		UserID:  comment.UserID,
	}

	ctx.JSON(http.StatusOK, gin.H{"comment": jsonComment})
}
