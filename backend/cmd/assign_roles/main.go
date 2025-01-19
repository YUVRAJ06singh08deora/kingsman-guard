package main

import (
	"context"
	"kingsmenguard/internal/handlers"
	"kingsmenguard/internal/services"
	"kingsmenguard/pkg/db"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func assignRolesHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Initialize the service for assigning roles
	roleService := services.NewGameService()
	log.Printf("Received request: Path=%s, HTTPMethod=%s, Body=%s", request.Path, request.HTTPMethod, request.Body)
	var err error

	// Check if the DynamoDB connection is established
	if db.DynamoInstance.Client == nil {
		if err = db.Connect(ctx); err != nil {
			log.Printf("Error connecting to DynamoDB: %v", err)
			return events.APIGatewayProxyResponse{
				StatusCode: 500,
				Body:       "Internal server error: unable to connect to database",
				Headers:    map[string]string{"Access-Control-Allow-Origin": "*"},
			}, nil
		}
	}

	var response events.APIGatewayProxyResponse

	// Handle the POST request for assigning roles
	if request.HTTPMethod == "POST" {
		// Call the handler to assign roles to players
		response, err = handlers.AssignRoles(ctx, request, roleService)

		if err != nil {
			log.Printf("Error assigning roles: %v", err)
			return events.APIGatewayProxyResponse{
				StatusCode: 500,
				Body:       "Internal server error: failed to assign roles",
				Headers:    map[string]string{"Access-Control-Allow-Origin": "*"},
			}, nil
		}
	} else {
		// Return 405 if the method is not supported
		log.Printf("Unsupported HTTP method: %s", request.HTTPMethod)
		return events.APIGatewayProxyResponse{
			StatusCode: 405,
			Body:       "Method Not Allowed",
			Headers:    map[string]string{"Access-Control-Allow-Origin": "*"},
		}, nil
	}

	// CORS headers for the response
	if response.Headers == nil {
		response.Headers = make(map[string]string)
	}
	response.Headers["Access-Control-Allow-Origin"] = "*"
	response.Headers["Access-Control-Allow-Methods"] = "OPTIONS,POST"
	response.Headers["Access-Control-Allow-Headers"] = "Content-Type, Authorization"

	return response, nil
}

func main() {
	// Start the Lambda function
	lambda.Start(assignRolesHandler)
}
