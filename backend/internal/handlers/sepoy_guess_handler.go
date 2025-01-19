package handlers

import (
	"context"
	"encoding/json"
	"kingsmenguard/internal/models"
	"kingsmenguard/internal/services"

	"github.com/aws/aws-lambda-go/events"
)

// SepoyGuess handles the logic for the player making a guess during the game round.
func SepoyGuess(ctx context.Context, request events.APIGatewayProxyRequest, gameService services.GameService) (events.APIGatewayProxyResponse, error) {
	// Check if the HTTP method is POST
	if !CheckHttpMethod(request.HTTPMethod, "POST") {
		return CreateJsonResponse(405, ApiResponse{Success: false, Message: "Invalid HTTP Method"}), nil
	}

	// Unmarshal the request body into a SepoyGuessRequest object
	var guessRequest models.SepoyGuessRequest
	err := json.Unmarshal([]byte(request.Body), &guessRequest)
	if err != nil {
		LogInfo(request.RequestContext, "ERROR", "Could not parse JSON")
		return CreateJsonResponse(400, ApiResponse{Success: false, Message: "Could not parse JSON"}), nil
	}

	// Call the GameService to process the player's guess
	err = gameService.ProcessSepoyGuess(ctx, &guessRequest)
	if err != nil {
		LogInfo(request.RequestContext, "ERROR", err.Error())
		return CreateJsonResponse(500, ApiResponse{Success: false, Message: "Could not process guess"}), nil
	}

	// Log the success and return a response
	LogInfo(request.RequestContext, "SUCCESS", "Guess processed successfully")
	return CreateJsonResponse(200, ApiResponse{Success: true, Message: "Guess Processed Successfully"}), nil
}
