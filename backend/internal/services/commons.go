package services

import (
	"context"
	"kingsmenguard/internal/models"
)

type GameService interface {
	CreateGame(ctx context.Context, g *models.CreateGameRequest) (*models.Game, error)
	AssignRoles(ctx context.Context, player *models.AssignRoleRequest) error
	GetLeaderboard(ctx context.Context, request models.LeaderboardRequest) ([]models.Player, error)
	EndGame(ctx context.Context, gameCode string) ([]models.Player, error)
	JoinGame(ctx context.Context, p *models.JoinGameRequest) error
	ProcessSepoyGuess(ctx context.Context, sepoyGuess *models.SepoyGuessRequest) error
	StartGame(ctx context.Context, gameID string) error
	updatePlayerPoints(ctx context.Context, game_code string, RoundID int, playerID string, pointsToAdd int) error
}

const GamesTableName = "Games"
const PlayersTableName = "Players"
const LeaderboardTableName = "Leaderboard"
const RoundsTableName = "Rounds"
const GameActionsTableName = "GameActions"

type DefaultGameService struct{}

func NewGameService() GameService {
	return &DefaultGameService{}
}
