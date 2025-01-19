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

func getLeaderboardHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	leaderboardService := services.NewGameService()
	log.Printf("Received request: Path=%s, HTTPMethod=%s, Body=%s", request.Path, request.HTTPMethod, request.Body)
	var err error
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

	if request.HTTPMethod == "GET" {
		// Call the handler for getting the leaderboard
		response, err = handlers.GetLeaderboard(ctx, request, leaderboardService)

		if err != nil {
			log.Printf("Error getting leaderboard: %v", err)
			return events.APIGatewayProxyResponse{
				StatusCode: 500,
				Body:       "Internal server error: failed to get leaderboard",
				Headers:    map[string]string{"Access-Control-Allow-Origin": "*"},
			}, nil
		}
	} else {
		log.Printf("Unsupported HTTP method: %s", request.HTTPMethod)
		return events.APIGatewayProxyResponse{
			StatusCode: 405,
			Body:       "Method Not Allowed",
			Headers:    map[string]string{"Access-Control-Allow-Origin": "*"},
		}, nil
	}

	// CORS headers
	if response.Headers == nil {
		response.Headers = make(map[string]string)
	}
	response.Headers["Access-Control-Allow-Origin"] = "*"
	response.Headers["Access-Control-Allow-Methods"] = "OPTIONS,GET"
	response.Headers["Access-Control-Allow-Headers"] = "Content-Type, Authorization"

	return response, nil
}

func main() {
	// Start the Lambda function
	lambda.Start(getLeaderboardHandler)
}
