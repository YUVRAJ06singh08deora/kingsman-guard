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

// UpdatePlayerPoints updates the points of a specific player
func (s *DefaultGameService) updatePlayerPoints(ctx context.Context, gameCode string, roundID int, playerID string, pointsToAdd int) error {
	// Query to get the current player points
	playerKeyCondExpr := "player_id = :player_id and game_code = :game_code"
	playerExpressionAttrs := map[string]types.AttributeValue{
		":player_id": &types.AttributeValueMemberS{Value: playerID},
		":game_code": &types.AttributeValueMemberS{Value: gameCode},
	}

	getPlayerInput := &dynamodb.QueryInput{
		TableName:                 aws.String(PlayersTableName),
		KeyConditionExpression:    aws.String(playerKeyCondExpr),
		ExpressionAttributeValues: playerExpressionAttrs,
	}

	playerResult, err := db.DynamoInstance.Client.Query(ctx, getPlayerInput)
	if err != nil {
		log.Println("Error querying player from DynamoDB", err.Error())
		return err
	}

	if len(playerResult.Items) == 0 {
		log.Println("Player not found")
		return fmt.Errorf("player not found")
	}

	player := models.Player{}
	err = attributevalue.UnmarshalMap(playerResult.Items[0], &player)
	if err != nil {
		log.Println("Error unmarshalling player data", err.Error())
		return err
	}

	// Step 1: Update the CurrentRoundScore with the pointsToAdd
	newCurrentRoundScore := pointsToAdd

	// Step 2: Add the CurrentRoundScore to the TotalScore
	newTotalScore := player.TotalScore + newCurrentRoundScore

	// Step 3: Update both CurrentRoundScore and TotalScore in the DynamoDB table
	updatePlayerInput := &dynamodb.UpdateItemInput{
		TableName: aws.String(PlayersTableName),
		Key: map[string]types.AttributeValue{
			"player_id": &types.AttributeValueMemberS{Value: playerID},
			"game_code": &types.AttributeValueMemberS{Value: gameCode},
		},
		UpdateExpression: aws.String("SET current_round_score = :new_current_round_score, total_score = :new_total_score"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":new_current_round_score": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", newCurrentRoundScore)},
			":new_total_score":         &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", newTotalScore)},
		},
	}

	_, err = db.DynamoInstance.Client.UpdateItem(ctx, updatePlayerInput)
	if err != nil {
		log.Println("Error updating player points in DynamoDB", err.Error())
		return err
	}

	log.Printf("Player %s's scores updated: CurrentRoundScore = %d, TotalScore = %d for Game %s", playerID, newCurrentRoundScore, newTotalScore, gameCode)
	return nil
}
