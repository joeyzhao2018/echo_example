package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	ddlambda "github.com/DataDog/datadog-lambda-go"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


var echoLambda *echoadapter.EchoLambda

func init() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	fmt.Print("Starting the server")
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"msg": "hello world", "time": time.Now().Format(time.UnixDate)})
	})

	echoLambda = echoadapter.New(e)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Print("Handling the request")
	res, err:= echoLambda.ProxyWithContext(ctx, req)
	if err != nil {
		fmt.Print("Error handling the request")
	}
	return res, err
}

func main() {
    lambda.Start(ddlambda.WrapHandler(Handler,nil))
}
