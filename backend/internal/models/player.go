package models

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Player struct {
	PlayerID          string    `json:"player_id" dynamodbav:"player_id"`
	PlayerName        string    `json:"player_name" dynamodbav:"player_name"`
	GameCode          string    `json:"game_code" dynamodbav:"game_code"`
	Role              string    `json:"role" dynamodbav:"role"`
	TotalScore        int       `json:"total_score" dynamodbav:"total_score"`
	RoundNumber       int       `json:"round_number" dynamodbav:"round_number"`
	CurrentRoundScore int       `json:"current_round_score" dynamodbav:"current_round_score"`
	JoinedAt          time.Time `json:"joined_at" dynamodbav:"joined_at"`
	UpdatedAt         time.Time `json:"updated_at" dynamodbav:"updated_at"`
	RoleAssignedAt    time.Time `json:"role_assigned_at" dynamodbav:"role_assigned_at"`
}

type JoinGameRequest struct {
	PlayerID   string `json:"player_id" dynamodbav:"player_id"`
	GameCode   string `json:"game_code" dynamodbav:"game_code"`
	PlayerName string `json:"player_name" dynamodbav:"player_name"`
}

type AssignRoleRequest struct {
	Player1ID   string `json:"player_1_id" dynamodbav:"player_1_id"`
	Player2ID   string `json:"player_2_id" dynamodbav:"player_2_id"`
	Player3ID   string `json:"player_3_id" dynamodbav:"player_3_id"`
	Player4ID   string `json:"player_4_id" dynamodbav:"player_4_id"`
	GameCode    string `json:"game_code" dynamodbav:"game_code"`
	RoundNumber int    `json:"round_number" dynamodbav:"round_number"`
}

type LeaderboardRequest struct {
	GameCode string `json:"game_code" dynamodbav:"game_code"`
}

func (p *Player) ToDynamoDBAV() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"player_id":  &types.AttributeValueMemberS{Value: p.PlayerID},
		"game_code":  &types.AttributeValueMemberS{Value: p.GameCode},
		"role":       &types.AttributeValueMemberS{Value: p.Role},
		"joined_at":  &types.AttributeValueMemberS{Value: p.JoinedAt.Format(time.RFC3339)},
		"updated_at": &types.AttributeValueMemberS{Value: p.UpdatedAt.Format(time.RFC3339)},
	}
}
