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
)

func (s *DefaultGameService) StartGame(ctx context.Context, gameCode string) error {

	// Step 1: Fetch the game details from DynamoDB
	key := struct {
		GameCode string `dynamodbav:"game_code" json:"game_code"`
	}{GameCode: gameCode}

	avs, err := attributevalue.MarshalMap(key)
	if err != nil {
		log.Println("Error marshalling key:", err.Error())
		return err
	}

	getGameInput := &dynamodb.GetItemInput{
		TableName: aws.String(GamesTableName), // Replace with your DynamoDB table name
		Key:       avs,
	}

	gameResult, err := db.DynamoInstance.Client.GetItem(ctx, getGameInput)
	if err != nil {
		log.Println("Error fetching game from DynamoDB", err.Error())
		return err
	}

	if gameResult.Item == nil {
		log.Println("Game not found")
		return fmt.Errorf("game not found")
	}

	// Unmarshal the result into the game struct
	var game models.Game
	err = attributevalue.UnmarshalMap(gameResult.Item, &game)
	if err != nil {
		log.Println("Error unmarshalling game data", err.Error())
		return err
	}

	// Step 2: Update the game's status to "active"
	game.Status = "active"
	game.UpdatedAt = time.Now() // Ensure the updated time is set

	// Marshal the updated game object
	updatedGameItem, err := attributevalue.MarshalMap(game)
	if err != nil {
		log.Println("Error marshalling updated game data", err.Error())
		return err
	}

	// Step 3: Update the game status in DynamoDB
	updateGameInput := &dynamodb.PutItemInput{
		TableName: aws.String(GamesTableName), // Replace with your DynamoDB table name
		Item:      updatedGameItem,
	}

	_, err = db.DynamoInstance.Client.PutItem(ctx, updateGameInput)
	if err != nil {
		log.Println("Error updating game status in DynamoDB", err.Error())
		return err
	}

	log.Println("Game status updated to 'active' successfully")
	return nil
}
