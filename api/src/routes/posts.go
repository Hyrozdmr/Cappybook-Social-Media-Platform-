package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/makersacademy/go-react-acebook-template/api/src/controllers"
	"github.com/makersacademy/go-react-acebook-template/api/src/middleware"
)

func setupPostRoutes(baseRouter *gin.RouterGroup) {
	posts := baseRouter.Group("/posts")

	posts.POST("", middleware.AuthenticationMiddleware, controllers.CreatePost)
	posts.GET("", middleware.AuthenticationMiddleware, controllers.GetAllPosts)
}


func setupLikePostRoute(baseRouter *gin.RouterGroup) {
	post := baseRouter.Group("/posts/:id/likes")
	post.PUT("", middleware.AuthenticationMiddleware, controllers.UpdatePostLikes)
}

