package services

import (
	"context"
	"fmt"
	"kingsmenguard/internal/models"
	"kingsmenguard/pkg/db"
	"log"
	"math/rand"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// AssignRoles assigns roles to four players in a game
func (s *DefaultGameService) AssignRoles(ctx context.Context, assignRequest *models.AssignRoleRequest) error {

	// Step 1: Fetch the game
	queryInput := &dynamodb.QueryInput{
		TableName:              aws.String(GamesTableName),
		KeyConditionExpression: aws.String("game_code = :gameCode"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":gameCode": &types.AttributeValueMemberS{Value: assignRequest.GameCode},
		},
	}

	result, err := db.DynamoInstance.Client.Query(ctx, queryInput)
	if err != nil {
		log.Println("Error querying DynamoDB for player count:", err)
		return err
	}

	game := models.Game{}
	err = attributevalue.UnmarshalMap(result.Items[0], &game)
	if err != nil {
		log.Println("Error unmarshalling game data", err.Error())
		return err
	}
	log.Println("No of Rounds in Game : ", game.NumberOfRounds)
	log.Println("Incomminf Round number : ", assignRequest.RoundNumber)

	if game.NumberOfRounds < assignRequest.RoundNumber {
		game.Status = "completed"
		return fmt.Errorf("roundsOver")
	}
	if assignRequest.RoundNumber == 1 {
		game.Status = "active"
	}
	if assignRequest.RoundNumber == game.NumberOfRounds {
		game.Status = "completed"
	}
	// Marshal the updated Game struct to DynamoDB attribute values
	game_item, err := attributevalue.MarshalMap(game)
	if err != nil {
		log.Println("Error marshalling the player:", err)
		return err
	}
	// Update the game status in DynamoDB
	_, err = db.DynamoInstance.Client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(GamesTableName),
		Item:      game_item,
	})

	// Step 1: Validate the input
	if assignRequest.Player1ID == "" || assignRequest.Player2ID == "" || assignRequest.Player3ID == "" || assignRequest.Player4ID == "" {
		return fmt.Errorf("Not Enough Players to Assign Roles")
	}

	if assignRequest.GameCode == "" {
		return fmt.Errorf("Invalid Game Code")
	}

	if assignRequest.RoundNumber <= 0 {
		return fmt.Errorf("Invalid Round Number")
	}

	// Step 2: Get player details from DynamoDB
	var player1, player2, player3, player4 *models.Player

	player1, err = getPlayer(ctx, assignRequest.Player1ID, assignRequest.GameCode)
	if err != nil {
		return fmt.Errorf("error fetching player1: %v", err)
	}

	player2, err = getPlayer(ctx, assignRequest.Player2ID, assignRequest.GameCode)
	if err != nil {
		return fmt.Errorf("error fetching player2: %v", err)
	}

	player3, err = getPlayer(ctx, assignRequest.Player3ID, assignRequest.GameCode)
	if err != nil {
		return fmt.Errorf("error fetching player3: %v", err)
	}

	player4, err = getPlayer(ctx, assignRequest.Player4ID, assignRequest.GameCode)
	if err != nil {
		return fmt.Errorf("error fetching player4: %v", err)
	}

	// Step 3: Create a list of players
	players := []*models.Player{player1, player2, player3, player4}

	// Step 4: Create a list of roles
	roles := []string{"King", "Queen", "Thief", "Sepoy"}

	// Step 5: Create a new random source and generator, and shuffle the roles list
	source := rand.NewSource(time.Now().UnixNano()) // Get a new random source based on current time
	randGen := rand.New(source)                     // Create a new random generator using the source

	// Shuffle the roles list using the new random generator
	randGen.Shuffle(len(roles), func(i, j int) {
		roles[i], roles[j] = roles[j], roles[i]
	})

	// Step 6: Assign the shuffled roles to players
	for i, player := range players {
		// Ensure player has game_code to be part of the key
		player.Role = roles[i]
		player.CurrentRoundScore = 0
		player.RoundNumber = assignRequest.RoundNumber
		player.RoleAssignedAt = time.Now()

		// Step 7: Marshal the updated Player struct to DynamoDB attribute values
		item, err := attributevalue.MarshalMap(player)
		if err != nil {
			log.Println("Error marshalling the player:", err)
			return err
		}

		// Step 8: Update the player's role in DynamoDB
		_, err = db.DynamoInstance.Client.PutItem(ctx, &dynamodb.PutItemInput{
			TableName: aws.String(PlayersTableName),
			Item:      item,
		})
		if err != nil {
			log.Println("Error updating player role in DynamoDB:", err)
			return err
		}
	}

	// Step 9: Return success after all players have been updated
	log.Println("Roles assigned successfully to all players")
	return nil
}

// getPlayer retrieves a player from DynamoDB using the PlayerID and GameCode.
func getPlayer(ctx context.Context, playerID, gameCode string) (*models.Player, error) {
	// Step 1: Define the key for the player using the correct AttributeValue type
	key := map[string]types.AttributeValue{
		"player_id": &types.AttributeValueMemberS{Value: playerID},
		"game_code": &types.AttributeValueMemberS{Value: gameCode},
	}

	// Step 2: Fetch the player details from the DynamoDB Players table
	getPlayerInput := &dynamodb.GetItemInput{
		TableName: aws.String(PlayersTableName),
		Key:       key,
	}

	result, err := db.DynamoInstance.Client.GetItem(ctx, getPlayerInput)
	if err != nil {
		log.Println("Error fetching player from DynamoDB:", err)
		return nil, err
	}

	// Step 3: Check if the player exists
	if result.Item == nil {
		log.Println("Player not found")
		return nil, fmt.Errorf("player not found")
	}

	// Step 4: Unmarshal the result into the Player struct
	var player models.Player
	err = attributevalue.UnmarshalMap(result.Item, &player)
	if err != nil {
		log.Println("Error unmarshalling player data:", err)
		return nil, err
	}

	// Step 5: Return the player object
	return &player, nil
}
