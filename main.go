package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/go-self-learning/docs"
	"github.com/go-self-learning/handlers"
	"github.com/go-self-learning/types"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title Blog API
// @version 1.0
// @description This is the Swagger Testing API.
// @host localhost:5678
// @BasePath /v1

var ginLambda *ginadapter.GinLambda

func init() {

	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {
		c.Redirect(302, "/swagger/index.html")
	})

	dsn := "host=ep-cool-hill-178722.eu-central-1.aws.neon.tech user=postgres password=5xIfpXrVEW2s dbname=aws-todo  sslmode=require"
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
	ginLambda = ginadapter.New(router)

}

func HandleLambdaEvent(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(HandleLambdaEvent)
}