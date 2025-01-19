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

func (s *DefaultGameService) EndGame(ctx context.Context, gameCode string) ([]models.Player, error) {
	// Step 1: Update the Game's status to "Ended" in the Games table
	game := models.Game{}
	gameKey := map[string]types.AttributeValue{
		"game_code": &types.AttributeValueMemberS{Value: gameCode},
	}

	// Get the current game details
	getGameInput := &dynamodb.GetItemInput{
		TableName: aws.String(GamesTableName),
		Key:       gameKey,
	}

	result, err := db.DynamoInstance.Client.GetItem(ctx, getGameInput)
	if err != nil {
		log.Println("Error fetching game from DynamoDB", err.Error())
		return nil, err
	}

	if result.Item == nil {
		log.Println("Game not found")
		return nil, fmt.Errorf("game not found")
	}

	// Unmarshal game data
	err = attributevalue.UnmarshalMap(result.Item, &game)
	if err != nil {
		log.Println("Error unmarshalling game data", err.Error())
		return nil, err
	}

	// Set the game status as "Inactive"
	game.Status = "Inactive"
	game.UpdatedAt = time.Now()

	// Marshal the updated game object
	item, err := attributevalue.MarshalMap(game)
	if err != nil {
		log.Println("Error marshalling game for update", err.Error())
		return nil, err
	}

	// Update the game in DynamoDB
	_, err = db.DynamoInstance.Client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(GamesTableName),
		Item:      item,
	})
	if err != nil {
		log.Println("Error updating game status in DynamoDB", err.Error())
		return nil, err
	}

	// Step 3: Fetch all players' details for the game
	getPlayersInput := &dynamodb.QueryInput{
		TableName:              aws.String(PlayersTableName),
		IndexName:              aws.String("RoleIndex"),
		KeyConditionExpression: aws.String("game_code = :game_code"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":game_code": &types.AttributeValueMemberS{Value: gameCode},
		},
	}

	PlayersResult, err := db.DynamoInstance.Client.Query(ctx, getPlayersInput)
	if err != nil {
		log.Println("Error querying updated players from DynamoDB", err.Error())
		return nil, err
	}

	// Unmarshal the updated player data into a slice of players
	var Players []models.Player
	for _, item := range PlayersResult.Items {
		var player models.Player
		err := attributevalue.UnmarshalMap(item, &player)
		if err != nil {
			log.Println("Error unmarshalling updated player data", err.Error())
			return nil, err
		}
		Players = append(Players, player)
	}

	// Step 4: Return the updated player details
	log.Println("All player details for game", gameCode, ":", Players)

	log.Println("Game ended successfully, leaderboard updated.")
	return Players, nil // Return the players and nil error (success)
}
