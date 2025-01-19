package handlers

import (
	"context"
	"encoding/json"
	"kingsmenguard/internal/models"
	"kingsmenguard/internal/services"

	"github.com/aws/aws-lambda-go/events"
)

// JoinGame handles the logic for a player joining a game.
func JoinGame(ctx context.Context, request events.APIGatewayProxyRequest, gameService services.GameService) (events.APIGatewayProxyResponse, error) {
	// Check if the HTTP method is POST
	if !CheckHttpMethod(request.HTTPMethod, "POST") {
		return CreateJsonResponse(405, ApiResponse{Success: false, Message: "Invalid HTTP Method"}), nil
	}

	// Unmarshal the request body into a JoinGameRequest object
	var joinGameRequest models.JoinGameRequest
	err := json.Unmarshal([]byte(request.Body), &joinGameRequest)
	if err != nil {
		LogInfo(request.RequestContext, "ERROR", "Could not parse JSON")
		return CreateJsonResponse(400, ApiResponse{Success: false, Message: "Could not parse JSON"}), nil
	}

	// Call the GameService to handle the logic for the player joining the game
	err = gameService.JoinGame(ctx, &joinGameRequest)
	if err != nil {
		LogInfo(request.RequestContext, "ERROR", err.Error())
		return CreateJsonResponse(500, ApiResponse{Success: false, Message: "Could not join game"}), nil
	}

	// Log the success and return a response
	LogInfo(request.RequestContext, "SUCCESS", "Player joined game successfully")
	return CreateJsonResponse(200, ApiResponse{Success: true, Message: "Player Joined Successfully"}), nil
}
