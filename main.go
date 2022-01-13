package main

import (
	"log"
	"net/http"
	"strings"
	"surekapi/auth"
	"surekapi/handler"
	"surekapi/helper"
	"surekapi/kategoripenerima"
	"surekapi/naskahdinas"
	"surekapi/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/surekapi_db?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	kategoriPenerimaRepository := kategoripenerima.NewRepository(db)
	naskahDinasRepository := naskahdinas.NewRepository(db)

	userService := user.NewService(userRepository)
	kategoriPenerimaService := kategoripenerima.NewService(kategoriPenerimaRepository)
	naskahDinasService := naskahdinas.NewService(naskahDinasRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandle(userService, authService)
	kategoriPenerimaHandler := handler.NewKategoriPenerimaHandler(kategoriPenerimaService)
	naskahDinasHandler := handler.NewNaskahDinasHandler(naskahDinasService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/username_checker", userHandler.CheckUsernameAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)
	api.GET("/kategori_penerima", authMiddleware(authService, userService), kategoriPenerimaHandler.FindKategoriPenerima)
	api.GET("/naskah_dinas", authMiddleware(authService, userService), naskahDinasHandler.FindNaskahDinas)
	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))
		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
