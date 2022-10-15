package routes

import (
	"final-assignment/configs"
	"final-assignment/controllers"
	"final-assignment/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func LoadRoute() {
	configs.DBInit()

	router := gin.Default()

	//User
	router.POST("/users/register", controllers.Register)
	router.POST("/users/login", controllers.Login)
	router.PUT("/users/:userId", middleware.VerifyBearer(), controllers.UpdateUser)
	router.DELETE("/users", middleware.VerifyBearer(), controllers.DeleteUser)

	//Photo
	router.POST("/photos", middleware.VerifyBearer(), controllers.CreatePhoto)
	router.GET("/photos", middleware.VerifyBearer(), controllers.GetPhoto)
	router.PUT("/photos/:photoId", middleware.VerifyBearer(), controllers.UpdatePhoto)
	router.DELETE("/photos/:photoId", middleware.VerifyBearer(), controllers.DeletePhoto)

	//Comment
	router.POST("/comments", middleware.VerifyBearer(), controllers.CreateComment)
	router.GET("/comments", middleware.VerifyBearer(), controllers.GetComment)
	router.PUT("/comments/:commentId", middleware.VerifyBearer(), controllers.UpdateComment)
	router.DELETE("/comments/:commentId", middleware.VerifyBearer(), controllers.DeleteComment)

	//Social Media
	router.POST("/socialmedias", middleware.VerifyBearer(), controllers.CreateSocialMedia)
	router.GET("/socialmedias", middleware.VerifyBearer(), controllers.GetSocialMedia)
	router.PUT("/socialmedias/:socialMediaId", middleware.VerifyBearer(), controllers.UpdateSocialMedia)
	router.DELETE("/socialmedias/:socialMediaId", middleware.VerifyBearer(), controllers.DeleteSocialMedia)

	router.Run(fmt.Sprintf("localhost:%v", configs.Env.Port))
}
