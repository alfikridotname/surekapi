package router

import (
	"surekapi/auth"
	"surekapi/config"
	"surekapi/handler"
	"surekapi/middleware"
	"surekapi/recipient_category"
	"surekapi/script"
	"surekapi/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupConnection()
)

func GetRouter() {
	userRepository := user.NewRepository(db)
	recipientCategoryRepository := recipient_category.NewRepository(db)
	scriptRepository := script.NewRepository(db)

	userService := user.NewService(userRepository)
	recipientCategoryService := recipient_category.NewService(recipientCategoryRepository)
	scriptService := script.NewService(scriptRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandle(userService, authService)
	recipientCategoryHandler := handler.NewRecipientCategoryHandler(recipientCategoryService)
	scriptHandler := handler.NewScriptHandler(scriptService)

	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/username_checker", userHandler.CheckUsernameAvailability)
	api.POST("/avatars", middleware.AuthMiddleware(authService, userService), userHandler.UploadAvatar)
	api.GET("/recipient_category", middleware.AuthMiddleware(authService, userService), recipientCategoryHandler.GetRecipientCategory)
	api.GET("/scripts", middleware.AuthMiddleware(authService, userService), scriptHandler.GetScripts)
	router.Run(":8080")
}
