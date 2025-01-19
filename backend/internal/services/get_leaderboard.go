package services

import (
	"context"
	"fmt"
	"kingsmenguard/internal/models"
	"kingsmenguard/pkg/db"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// GetLeaderboard retrieves the leaderboard data from DynamoDB
func (s *DefaultGameService) GetLeaderboard(ctx context.Context, request models.LeaderboardRequest) ([]models.Player, error) {
	// Prepare the DynamoDB query to fetch leaderboard data
	// In this example, assuming we want to fetch sorted leaderboard entries
	// You can adjust this to your exact use case (e.g., paginated queries, filtering, etc.)
	// Query the Players table to fetch all players' details for the game
	getLeaderBoardInput := &dynamodb.QueryInput{
		TableName:              aws.String(PlayersTableName),
		IndexName:              aws.String("RoleIndex"), // Using the GSI to get all players in the game
		KeyConditionExpression: aws.String("game_code = :game_code"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":game_code": &types.AttributeValueMemberS{Value: request.GameCode},
		},
	}

	playersResult, err := db.DynamoInstance.Client.Query(ctx, getLeaderBoardInput)

	if err != nil {
		// Return an error if the query fails
		return nil, fmt.Errorf("failed to query leaderboard: %w", err)
	}

	// Initialize a slice to hold the leaderboard entries
	var leaderboard []models.Player

	// Unmarshal the result items into LeaderboardEntry objects
	for _, item := range playersResult.Items {
		var entry models.Player
		if err := attributevalue.UnmarshalMap(item, &entry); err != nil {
			// Return an error if unmarshaling fails
			return nil, fmt.Errorf("failed to unmarshal leaderboard item: %w", err)
		}
		leaderboard = append(leaderboard, entry) // Append the entry to the leaderboard slice
	}

	// Log the successful retrieval of the leaderboard
	log.Printf("Successfully retrieved %d leaderboard entries", len(leaderboard))

	// Return the leaderboard entries
	return leaderboard, nil
}
