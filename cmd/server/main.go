package main

import (
	"context"
	"fmt"
	_ "go-endpoints/internal/infrastructure/db"
	"go-endpoints/internal/infrastructure/server"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
)

var echoLambda *echoadapter.EchoLambda

func init() {
	echoLambda = echoadapter.New(server.Server)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("NEW EVENT!!!")
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
