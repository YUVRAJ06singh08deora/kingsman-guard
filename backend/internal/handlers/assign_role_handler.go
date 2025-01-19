package handlers

import (
	"context"
	"encoding/json"
	"kingsmenguard/internal/models"
	"kingsmenguard/internal/services"

	"github.com/aws/aws-lambda-go/events"
)

// AssignRoles handles the assignment of roles to players in a game.
func AssignRoles(ctx context.Context, request events.APIGatewayProxyRequest, gameService services.GameService) (events.APIGatewayProxyResponse, error) {
	// Check if the HTTP method is POST
	if !CheckHttpMethod(request.HTTPMethod, "POST") {
		return CreateJsonResponse(405, ApiResponse{Success: false, Message: "Invalid HTTP Method"}), nil
	}

	// Unmarshal the request body into an AssignRoles request object
	var assignRequest models.AssignRoleRequest
	err := json.Unmarshal([]byte(request.Body), &assignRequest)
	if err != nil {
		LogInfo(request.RequestContext, "ERROR", "Could not parse JSON")
		return CreateJsonResponse(400, ApiResponse{Success: false, Message: "Could not parse JSON"}), nil
	}

	// Call the GameService to assign roles in the game
	err = gameService.AssignRoles(ctx, &assignRequest)
	if err != nil {
		LogInfo(request.RequestContext, "ERROR", err.Error())
		return CreateJsonResponse(500, ApiResponse{Success: false, Message: err.Error()}), nil
	}

	// Log the success and return a response
	LogInfo(request.RequestContext, "SUCCESS", "Roles assigned successfully")
	return CreateJsonResponse(200, ApiResponse{Success: true, Message: "Roles Assigned Successfully"}), nil
}
