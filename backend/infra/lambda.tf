# Lambda function for creating a game
resource "aws_lambda_function" "create_game_lambda" {
  function_name    = "CreateGameLambda"
  handler          = "create_game.main"
  runtime          = "provided.al2"
  role             = aws_iam_role.lambda_execution_role.arn
  filename         = "${path.module}/../dist/create_game.zip"
  source_code_hash = filebase64sha256("${path.module}/../dist/create_game.zip")
}

# Lambda function for players to join a game
resource "aws_lambda_function" "join_game_lambda" {
  function_name    = "JoinGameLambda"
  handler          = "join_game.main"
  runtime          = "provided.al2"
  role             = aws_iam_role.lambda_execution_role.arn
  filename         = "${path.module}/../dist/join_game.zip"
  source_code_hash = filebase64sha256("${path.module}/../dist/join_game.zip")
}

# Lambda function for starting the game
resource "aws_lambda_function" "start_game_lambda" {
  function_name    = "StartGameLambda"
  handler          = "start_game.main"
  runtime          = "provided.al2"
  role             = aws_iam_role.lambda_execution_role.arn
  filename         = "${path.module}/../dist/start_game.zip"
  source_code_hash = filebase64sha256("${path.module}/../dist/start_game.zip")
}

# Lambda function for assigning roles to players
resource "aws_lambda_function" "assign_roles_lambda" {
  function_name    = "AssignRolesLambda"
  handler          = "assign_roles.main"
  runtime          = "provided.al2"
  role             = aws_iam_role.lambda_execution_role.arn
  filename         = "${path.module}/../dist/assign_roles.zip"
  source_code_hash = filebase64sha256("${path.module}/../dist/assign_roles.zip")
}

# Lambda function for Sepoy to make a guess about the Thief
resource "aws_lambda_function" "sepoy_guess_lambda" {
  function_name    = "SepoyGuessLambda"
  handler          = "sepoy_guess.main"
  runtime          = "provided.al2"
  role             = aws_iam_role.lambda_execution_role.arn
  filename         = "${path.module}/../dist/sepoy_guess.zip"
  source_code_hash = filebase64sha256("${path.module}/../dist/sepoy_guess.zip")
}


# Lambda function for ending the game and showing the final leaderboard
resource "aws_lambda_function" "end_game_lambda" {
  function_name    = "EndGameLambda"
  handler          = "end_game.main"
  runtime          = "provided.al2"
  role             = aws_iam_role.lambda_execution_role.arn
  filename         = "${path.module}/../dist/end_game.zip"
  source_code_hash = filebase64sha256("${path.module}/../dist/end_game.zip")
}

# Lambda function for ending the game and showing the final leaderboard
resource "aws_lambda_function" "get_leaderboard_lambda" {
  function_name    = "GetLeaderboardLambda"
  handler          = "get_leaderboard.main"
  runtime          = "provided.al2"
  role             = aws_iam_role.lambda_execution_role.arn
  filename         = "${path.module}/../dist/get_leaderboard.zip"
  source_code_hash = filebase64sha256("${path.module}/../dist/get_leaderboard.zip")
}

