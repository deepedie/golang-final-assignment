package router

import (
	"assignment-4/controllers"
	"assignment-4/database"
	"assignment-4/middlewares"
	repositories "assignment-4/repository"
	"assignment-4/services"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	db := database.GetDB()

	// Set up user repository and service
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	userRouter := r.Group("/users")

	{
		userRouter.POST("/register", userController.Register)
		userRouter.POST("/login", userController.Login)
	}

	// Set up photo repository and service
	photoRepo := repositories.NewPhotoRepository(db)
	photoService := services.NewPhotoService(photoRepo)
	photoController := controllers.NewPhotoController(photoService)

	photoRoutes := r.Group("/photos")
	photoRoutes.Use(middlewares.Authentication())
	{
		photoRoutes.GET("/", photoController.GetAllPhotos)
		photoRoutes.GET("/:id", photoController.GetPhotoByID)
		photoRoutes.POST("/", photoController.CreatePhoto)
		photoRoutes.PUT("/:id", photoController.UpdatePhoto)
		photoRoutes.DELETE("/:id", photoController.DeletePhoto)
	}
	// Set up comment repository and service
	commentRepo := repositories.NewCommentRepository(db)
	commentService := services.NewCommentService(commentRepo)
	commentController := controllers.NewCommentController(commentService)

	commentRoutes := r.Group("/comments")
	commentRoutes.Use(middlewares.Authentication())
	{
		commentRoutes.GET("/", commentController.GetAllComments)
		commentRoutes.GET("/:id", commentController.GetCommentByID)
		commentRoutes.POST("/", commentController.CreateComment)
		commentRoutes.PUT("/:id", commentController.UpdateComment)
		commentRoutes.DELETE("/:id", commentController.DeleteComment)
	}

	// Set up social media repository and service
	socialMediaRepo := repositories.NewSocialMediaRepository(db)
	socialMediaService := services.NewSocialMediaService(socialMediaRepo)
	socialMediaController := controllers.NewSocialMediaController(socialMediaService)

	socialMediaRoutes := r.Group("/socialmedia")
	socialMediaRoutes.Use(middlewares.Authentication())
	{
		socialMediaRoutes.GET("/", socialMediaController.GetAllSocialMedias)
		socialMediaRoutes.GET("/:id", socialMediaController.GetSocialMediaByID)
		socialMediaRoutes.POST("/", socialMediaController.CreateSocialMedia)
		socialMediaRoutes.PUT("/:id", socialMediaController.UpdateSocialMedia)
		socialMediaRoutes.DELETE("/:id", socialMediaController.DeleteSocialMedia)
	}

	return r
}
