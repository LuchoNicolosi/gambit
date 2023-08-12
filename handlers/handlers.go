package handlers

import (
	"fmt"
	// "strconv"

	"github.com/aws/aws-lambda-go/events"
)

func Handlers(path string, method string, body string, headers map[string]string, req events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Voy a procesar " + path + " > " + method)

	// id := req.PathParameters["id"]
	//	idn, _ := strconv.Atoi(id)

	return 400, "Method invalid"
}
