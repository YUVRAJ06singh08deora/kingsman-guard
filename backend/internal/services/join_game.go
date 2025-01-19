package services

import (
	"context"
	"fmt"
	"kingsmenguard/internal/models"
	"kingsmenguard/pkg/db"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// JoinGame handles a player joining a game.
func (s *DefaultGameService) JoinGame(ctx context.Context, gameReq *models.JoinGameRequest) error {
	// Validate that the player has a valid game ID and player ID
	if gameReq.GameCode == "" || gameReq.PlayerID == "" {
		return fmt.Errorf("No Values received for player ID or game ID")
	}

	// Step 1: Count the number of players in the game
	playerCount, err := countPlayersInGame(ctx, gameReq.GameCode)
	if err != nil {
		return fmt.Errorf("failed to count players for game %s: %v", gameReq.GameCode, err)
	}

	// Step 2: Check if the player count is 4
	if playerCount == 4 {
		return fmt.Errorf("players already full for game %s", gameReq.GameCode)
	}

	// Step 3: Prepare the player object
	var p models.Player
	p.PlayerID = gameReq.PlayerID
	p.PlayerName = gameReq.PlayerName
	p.GameCode = gameReq.GameCode
	p.JoinedAt = time.Now()
	p.UpdatedAt = time.Now()
	p.CurrentRoundScore = 0
	p.RoundNumber = 0
	p.TotalScore = 0
	p.Role = "unassigned"
	p.RoleAssignedAt = time.Now()

	// Step 4: Marshal the Player struct to DynamoDB attribute values
	item, err := attributevalue.MarshalMap(p)
	if err != nil {
		return fmt.Errorf("failed to marshal player struct: %w", err)
	}

	// Step 5: Insert the player into DynamoDB
	_, err = db.DynamoInstance.Client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:           aws.String(PlayersTableName),
		Item:                item,
		ConditionExpression: aws.String("attribute_not_exists(player_id)"),
	})

	if err != nil {
		return fmt.Errorf("failed to insert player with PlayerID %s into DynamoDB: %w", p.PlayerID, err)
	}

	// Log success
	log.Printf("Player %s successfully joined game %s at %s", p.PlayerID, p.GameCode, p.JoinedAt)

	// Return success (optional: return player object or ID for further use)
	return nil
}

// countPlayersInGame counts the number of players already in the game with the given game_code.
func countPlayersInGame(ctx context.Context, gameCode string) (int, error) {
	// Step 1: Define the query to count players by game_code using the RoleIndex GSI
	queryInput := &dynamodb.QueryInput{
		TableName:              aws.String(PlayersTableName),
		IndexName:              aws.String("RoleIndex"), // Using the GSI that has game_code as the hash key
		KeyConditionExpression: aws.String("game_code = :gameCode"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":gameCode": &types.AttributeValueMemberS{Value: gameCode},
		},
	}

	// Step 2: Query DynamoDB to get the count of players for the specified game_code
	result, err := db.DynamoInstance.Client.Query(ctx, queryInput)
	if err != nil {
		log.Println("Error querying DynamoDB for player count:", err)
		return 0, err
	}

	// Step 3: Return the count of players (the number of items returned by the query)
	return int(result.Count), nil
}
