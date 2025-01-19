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
	"github.com/google/uuid"
)

// Helper function to generate a unique Game ID
func generateGameID() string {
	return uuid.New().String() // Returns a new unique game ID using UUID
}

// Helper function to generate a short Game Code (6 characters)
func generateGameCode() string {
	return uuid.New().String()[:6] // Trims the UUID to the first 6 characters to create a shorter Game Code
}

// CreateGame is responsible for creating a new game and storing it in DynamoDB
func (s *DefaultGameService) CreateGame(ctx context.Context, g *models.CreateGameRequest) (*models.Game, error) {
	// Initialize the game struct with values from the request
	game := models.Game{
		GameID:         generateGameID(),   // Generate a unique Game ID
		GameCode:       generateGameCode(), // Generate a unique Game Code
		CreatedAt:      time.Now(),         // Set the current timestamp as creation time
		UpdatedAt:      time.Now(),         // Set the current timestamp as update time
		NumberOfRounds: g.NumberOfRounds,   // Set the number of rounds from the request
		Status:         "created",          // Set the initial status as 'created'
	}

	// Marshal the Game struct into a format that can be stored in DynamoDB
	item, err := attributevalue.MarshalMap(game)
	if err != nil {
		// Return an error if marshaling fails
		return nil, fmt.Errorf("failed to marshal game struct: %w", err)
	}

	// Attempt to insert the game into DynamoDB
	_, err = db.DynamoInstance.Client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(GamesTableName), // Specify the DynamoDB table name
		Item:      item,                       // The marshaled game item to be inserted
	})
	if err != nil {
		// Return an error if inserting into DynamoDB fails
		return nil, fmt.Errorf("failed to insert game into DynamoDB: %w", err)
	}

	// Log the successful creation of the game with its Game ID and Game Code
	log.Printf("Game created successfully with GameID: %s and GameCode: %s", game.GameID, game.GameCode)

	// Return the created game object, including the generated GameID and GameCode
	return &game, nil
}
