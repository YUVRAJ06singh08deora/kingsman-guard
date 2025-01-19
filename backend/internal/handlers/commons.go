package handlers

import (
	"encoding/json" // Adjust the import path as needed
	"log"

	"github.com/aws/aws-lambda-go/events"
)

// ApiResponse is the common structure for the Lambda function responses
type ApiResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// CreateJsonResponse is a helper function to create consistent JSON responses
func CreateJsonResponse(statusCode int, resp ApiResponse) events.APIGatewayProxyResponse {
	body, _ := json.Marshal(resp)
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(body),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}

// CheckHttpMethod validates the HTTP method
func CheckHttpMethod(method string, expectedMethod string) bool {
	return method == expectedMethod
}

// LogInfo logs the request and context information
func LogInfo(requestContext events.APIGatewayProxyRequestContext, t string, msg string) {
	log.Printf("-> %s -> %s -> %s -> %s\n", t, requestContext.HTTPMethod, requestContext.ResourcePath, msg)
}
