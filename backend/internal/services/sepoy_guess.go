package services

import (
	"context"
	"kingsmenguard/internal/models"
	"kingsmenguard/pkg/db"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (s *DefaultGameService) ProcessSepoyGuess(ctx context.Context, sepoyGuess *models.SepoyGuessRequest) error {
	// Step 1: Fetch the selected player's details from the Players table
	playerSelectedId := sepoyGuess.SelectedPlayerID
	gameCode := sepoyGuess.GameCode
	roundID := sepoyGuess.RoundID
	guess := sepoyGuess.GuessedRole
	log.Println("the role to be guessed is", guess)
	// Query the Players table to fetch all players' details for the game
	getPlayersInput := &dynamodb.QueryInput{
		TableName:              aws.String(PlayersTableName),
		IndexName:              aws.String("RoleIndex"), // Using the GSI to get all players in the game
		KeyConditionExpression: aws.String("game_code = :game_code"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":game_code": &types.AttributeValueMemberS{Value: gameCode},
		},
	}

	playersResult, err := db.DynamoInstance.Client.Query(ctx, getPlayersInput)
	if err != nil {
		log.Println("Error querying players from DynamoDB", err.Error())
		return err
	}

	// Initialize variables to store player roles
	var thiefRole, kingRole, queenRole string

	// Iterate through players and assign roles
	for _, p := range playersResult.Items {
		playerData := models.Player{}
		err := attributevalue.UnmarshalMap(p, &playerData)
		if err != nil {
			log.Println("Error unmarshalling player role", err.Error())
			return err
		}

		// Assign roles to the variables
		if playerData.Role == "Thief" {
			thiefRole = playerData.PlayerID
		} else if playerData.Role == "King" {
			kingRole = playerData.PlayerID
		} else if playerData.Role == "Queen" {
			queenRole = playerData.PlayerID
		}
	}

	// Step 2: Check if the selected player ID matches the thief's ID
	if playerSelectedId == thiefRole {
		// Correct guess scenario
		// 1. Sepoy gets +150 points
		err = s.updatePlayerPoints(ctx, gameCode, roundID, sepoyGuess.SepoyId, 150)
		if err != nil {
			log.Println("Error updating Sepoy points", err.Error())
			return err
		}

		// 2. King gets +100 points
		err = s.updatePlayerPoints(ctx, gameCode, roundID, kingRole, 100)
		if err != nil {
			log.Println("Error updating King points", err.Error())
			return err
		}

		// 3. Queen gets +50 points
		err = s.updatePlayerPoints(ctx, gameCode, roundID, queenRole, 50)
		if err != nil {
			log.Println("Error updating Queen points", err.Error())
			return err
		}

		// 4. Thief gets -25 points
		err = s.updatePlayerPoints(ctx, gameCode, roundID, thiefRole, -25)
		if err != nil {
			log.Println("Error updating Thief points", err.Error())
			return err
		}

	} else {
		// Incorrect guess scenario

		// If the selected player is the King
		if playerSelectedId == kingRole {
			// 1. Sepoy gets -50 points
			err = s.updatePlayerPoints(ctx, gameCode, roundID, sepoyGuess.SepoyId, -50)
			if err != nil {
				log.Println("Error updating Sepoy points", err.Error())
				return err
			}

			// 2. King gets +100 points
			err = s.updatePlayerPoints(ctx, gameCode, roundID, kingRole, 100)
			if err != nil {
				log.Println("Error updating King points", err.Error())
				return err
			}

			// 3. Queen gets +50 points
			err = s.updatePlayerPoints(ctx, gameCode, roundID, queenRole, 50)
			if err != nil {
				log.Println("Error updating Queen points", err.Error())
				return err
			}

			// 4. Thief gets +25 points
			err = s.updatePlayerPoints(ctx, gameCode, roundID, thiefRole, 25)
			if err != nil {
				log.Println("Error updating Thief points", err.Error())
				return err
			}

		} else if playerSelectedId == queenRole {
			// If the selected player is the Queen
			// 1. Sepoy gets -25 points
			err = s.updatePlayerPoints(ctx, gameCode, roundID, sepoyGuess.SepoyId, -25)
			if err != nil {
				log.Println("Error updating Sepoy points", err.Error())
				return err
			}

			// 2. King gets +100 points
			err = s.updatePlayerPoints(ctx, gameCode, roundID, kingRole, 100)
			if err != nil {
				log.Println("Error updating King points", err.Error())
				return err
			}

			// 3. Queen gets +50 points
			err = s.updatePlayerPoints(ctx, gameCode, roundID, queenRole, 50)
			if err != nil {
				log.Println("Error updating Queen points", err.Error())
				return err
			}

			// 4. Thief gets +25 points
			err = s.updatePlayerPoints(ctx, gameCode, roundID, thiefRole, 25)
			if err != nil {
				log.Println("Error updating Thief points", err.Error())
				return err
			}
		}
	}

	log.Println("Sepoy guess processed successfully")
	return nil
}
