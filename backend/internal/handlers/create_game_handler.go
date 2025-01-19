package handlers

import (
	"context"
	"encoding/json"
	"kingsmenguard/internal/models"
	"kingsmenguard/internal/services"

	"github.com/aws/aws-lambda-go/events"
)

// CreateGame handles the creation of a new game.
func CreateGame(ctx context.Context, request events.APIGatewayProxyRequest, gameService services.GameService) (events.APIGatewayProxyResponse, error) {
	// Check if the HTTP method is POST
	if !CheckHttpMethod(request.HTTPMethod, "POST") {
		return CreateJsonResponse(405, ApiResponse{Success: false, Message: "Invalid HTTP Method"}), nil
	}

	// Unmarshal the request body into a Game object
	var gameReq models.CreateGameRequest
	err := json.Unmarshal([]byte(request.Body), &gameReq)
	if err != nil {
		LogInfo(request.RequestContext, "ERROR", "Could not parse JSON")
		return CreateJsonResponse(400, ApiResponse{Success: false, Message: "Could not parse JSON"}), nil
	}

	// Call the GameService to create the game in DynamoDB
	createdGame, err := gameService.CreateGame(ctx, &gameReq)
	if err != nil {
		LogInfo(request.RequestContext, "ERROR", err.Error())
		return CreateJsonResponse(500, ApiResponse{Success: false, Message: "Could not create game"}), nil
	}

	// Log the success and return a response
	LogInfo(request.RequestContext, "SUCCESS", "Game created successfully")

	// Return the game code and game ID as part of the response data
	return CreateJsonResponse(200, ApiResponse{
		Success: true,
		Message: "Game Created Successfully",
		Data:    map[string]string{"game_code": createdGame.GameCode, "game_id": createdGame.GameID},
	}), nil
}
