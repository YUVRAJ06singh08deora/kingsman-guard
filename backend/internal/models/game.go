package models

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Game struct {
	GameID         string    `json:"game_id" dynamodbav:"game_id"`
	GameCode       string    `json:"game_code" dynamodbav:"game_code"`
	Status         string    `json:"status" dynamodbav:"status"`
	NumberOfRounds int       `json:"number_of_rounds" dynamodbav:"number_of_rounds"`
	CreatedAt      time.Time `json:"created_at" dynamodbav:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" dynamodbav:"updated_at"`
}

type CreateGameRequest struct {
	NumberOfRounds int `json:"number_of_rounds" dynamodbav:"number_of_rounds"`
}

type EndGameRequest struct {
	GameCode string `json:"game_code" dynamodbav:"game_code"`
}
type StartGameRequest struct {
	GameCode string `json:"game_code" dynamodbav:"game_code"`
}

func (g *Game) ToDynamoDBAV() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"game_id":    &types.AttributeValueMemberS{Value: g.GameID},
		"game_code":  &types.AttributeValueMemberS{Value: g.GameCode},
		"created_at": &types.AttributeValueMemberS{Value: g.CreatedAt.Format(time.RFC3339)},
		"updated_at": &types.AttributeValueMemberS{Value: g.UpdatedAt.Format(time.RFC3339)},
	}
}
