package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-self-learning/docs"
	"github.com/go-self-learning/handlers"
	"github.com/go-self-learning/types"
	"gorm.io/driver/postgres"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// @title Blog API
// @version 1.0
// @description This is the Swagger Testing API.
// @host localhost:5678
// @BasePath /v1
func main() {

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	dsn := "host=localhost user=gorm_user password=gorm_password dbname=gorm_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&types.Users{}, &types.Posts{})
	/*--------------------------CRUD Operation using gorms USERS-------------------------------------*/
	userController := handlers.NewUserController(db)
	postController := handlers.NewPostController(db)

	router.GET("/api/v1/users", userController.GetUsers)
	router.POST("/api/v1/users", userController.CreateUser)
	router.GET("/api/v1/users/:id", userController.GetUser)
	router.PUT("/api/v1/users/:id", userController.UpdateUser)
	router.DELETE("/api/v1/users/:id", userController.DeleteUser)

	/*--------------------------CRUD Operation using gorms POSTS-------------------------------------*/
	router.GET("/api/v1/posts", postController.GetPosts)
	router.POST("/api/v1/posts", postController.CreatePost)
	router.GET("/api/v1/posts/:id", postController.GetPost)
	router.PUT("/api/v1/posts/:id", postController.UpdatePost)
	router.DELETE("/api/v1/posts/:id", postController.DeletePost)
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":5678")
}
