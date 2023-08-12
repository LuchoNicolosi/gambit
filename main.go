package main

import (
	"context"
	"os"

	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/luchonicolosi/gambit/awsgo"
	"github.com/luchonicolosi/gambit/handlers"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	awsgo.InicializoAws()

	var res *events.APIGatewayProxyResponse

	path := strings.Replace(request.RawPath, os.Getenv("UrlPrefix"), "", -1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	header := request.Headers

	status, message := handlers.Handlers(path, method, body, header, request)

	headersRes := map[string]string{
		"Content-Type": "application/json",
	}

	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(message),
		Headers:    headersRes,
	}

	return res, nil
}
