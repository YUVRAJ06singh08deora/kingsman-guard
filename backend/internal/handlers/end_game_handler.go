package handlers

import (
	"context"
	"encoding/json"
	"kingsmenguard/internal/models"
	"kingsmenguard/internal/services"

	"github.com/aws/aws-lambda-go/events"
)

// EndGame handles the logic to end a game and returns player details
func EndGame(ctx context.Context, request events.APIGatewayProxyRequest, gameService services.GameService) (events.APIGatewayProxyResponse, error) {
	// Check if the HTTP method is POST
	if !CheckHttpMethod(request.HTTPMethod, "POST") {
		return CreateJsonResponse(405, ApiResponse{Success: false, Message: "Invalid HTTP Method"}), nil
	}

	// Unmarshal the request body into an EndGame request object
	var endGameRequest models.EndGameRequest
	err := json.Unmarshal([]byte(request.Body), &endGameRequest)
	if err != nil {
		LogInfo(request.RequestContext, "ERROR", "Could not parse JSON")
		return CreateJsonResponse(400, ApiResponse{Success: false, Message: "Could not parse JSON"}), nil
	}

	// Call the GameService to end the game and retrieve player details
	players, err := gameService.EndGame(ctx, endGameRequest.GameCode)
	if err != nil {
		LogInfo(request.RequestContext, "ERROR", err.Error())
		return CreateJsonResponse(500, ApiResponse{Success: false, Message: "Could not end game"}), nil
	}

	// Log the success and return the response along with player details
	LogInfo(request.RequestContext, "SUCCESS", "Game ended successfully")

	// Prepare the response body with player details in the Data field
	responseBody := ApiResponse{
		Success: true,
		Message: "Game Ended Successfully",
		Data:    players, // Include the list of players here
	}

	// Marshal the response body
	body, err := json.Marshal(responseBody)
	if err != nil {
		LogInfo(request.RequestContext, "ERROR", "Could not marshal response body")
		return CreateJsonResponse(500, ApiResponse{Success: false, Message: "Internal Server Error"}), nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(body),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
