package handlers

import (
	"context"
	"kingsmenguard/internal/models"
	"kingsmenguard/internal/services"

	"github.com/aws/aws-lambda-go/events"
)

// GetLeaderboard handles the logic for retrieving the leaderboard data.
func GetLeaderboard(ctx context.Context, request events.APIGatewayProxyRequest, gameService services.GameService) (events.APIGatewayProxyResponse, error) {
	// Check if the HTTP method is GET
	if !CheckHttpMethod(request.HTTPMethod, "GET") {
		return CreateJsonResponse(405, ApiResponse{Success: false, Message: "Invalid HTTP Method"}), nil
	}

	// Retrieve game_code from query parameters
	gameCode := request.QueryStringParameters["game_code"]
	if gameCode == "" {
		return CreateJsonResponse(400, ApiResponse{Success: false, Message: "Game code is required"}), nil
	}

	// Initialize the leaderboard request with the game code
	leaderboardRequest := models.LeaderboardRequest{
		GameCode: gameCode, // Set the game_code from the query parameter
	}

	// Call the LeaderboardService to get the leaderboard data
	leaderboard, err := gameService.GetLeaderboard(ctx, leaderboardRequest)
	if err != nil {
		LogInfo(request.RequestContext, "ERROR", err.Error())
		return CreateJsonResponse(500, ApiResponse{Success: false, Message: "Could not fetch leaderboard data"}), nil
	}

	// Log the success and return the leaderboard response
	LogInfo(request.RequestContext, "SUCCESS", "Leaderboard fetched successfully")

	// Ensure that the leaderboard data is properly structured in the response
	return CreateJsonResponse(200, ApiResponse{
		Success: true,
		Message: "Leaderboard  Successfully fetched",
		Data:    leaderboard, // Directly passing the leaderboard data as the response Data field
	}), nil
}
