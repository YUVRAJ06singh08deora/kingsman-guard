package handlers

import (
	"context"
	"encoding/json"
	"kingsmenguard/internal/models"
	"kingsmenguard/internal/services"

	"github.com/aws/aws-lambda-go/events"
)

// StartGame handles the logic for starting a new game round.
func StartGame(ctx context.Context, request events.APIGatewayProxyRequest, gameService services.GameService) (events.APIGatewayProxyResponse, error) {
	// Check if the HTTP method is POST
	if !CheckHttpMethod(request.HTTPMethod, "POST") {
		return CreateJsonResponse(405, ApiResponse{Success: false, Message: "Invalid HTTP Method"}), nil
	}

	// Unmarshal the request body into a StartGameRequest object
	var startGameRequest models.StartGameRequest
	err := json.Unmarshal([]byte(request.Body), &startGameRequest)
	if err != nil {
		LogInfo(request.RequestContext, "ERROR", "Could not parse JSON")
		return CreateJsonResponse(400, ApiResponse{Success: false, Message: "Could not parse JSON"}), nil
	}

	// Call the GameService to start the game
	err = gameService.StartGame(ctx, startGameRequest.GameCode)
	if err != nil {
		LogInfo(request.RequestContext, "ERROR", err.Error())
		return CreateJsonResponse(500, ApiResponse{Success: false, Message: "Could not start the game"}), nil
	}

	// Log the success and return a response
	LogInfo(request.RequestContext, "SUCCESS", "Game started successfully")
	return CreateJsonResponse(200, ApiResponse{Success: true, Message: "Game Started Successfully"}), nil
}
