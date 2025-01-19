package models

type SepoyGuessRequest struct {
	GameCode         string `json:"game_code" dynamodbav:"game_code"`                   // The game ID
	RoundID          int    `json:"round_id" dynamodbav:"round_id"`                     // The round ID
	SepoyId          string `json:"sepoy_id" dynamodbav:"sepoy_id"`                     // The Sepoy player ID
	SelectedPlayerID string `json:"selected_player_id" dynamodbav:"selected_player_id"` // The ID of the player who has been guessed for theif
	GuessedRole      string `json:"guess" dynamodbav:"guess"`                           // The role the Sepoy is guessing (King, Queen, Thief)
}
