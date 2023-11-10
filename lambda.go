package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

type MyEvent struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type MyResponse struct {
	Message string `json:"message"`
}

func AwsLambdaHandler(event *MyEvent) (*MyResponse, error) {

	if event == nil {
		return nil, nil
	}
	return &MyResponse{Message: fmt.Sprintf("Hello, %s. You are %d years old.", event.Name, event.Age)}, nil
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println(request)
	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello, World",
	}
	return response, nil
}
