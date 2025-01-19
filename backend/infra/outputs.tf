# Output for Lambda ARNs
output "create_game_lambda_arn" {
  value = aws_lambda_function.create_game_lambda.arn
}

output "join_game_lambda_arn" {
  value = aws_lambda_function.join_game_lambda.arn
}

output "start_game_lambda_arn" {
  value = aws_lambda_function.start_game_lambda.arn
}

output "assign_roles_lambda_arn" {
  value = aws_lambda_function.assign_roles_lambda.arn
}

output "sepoy_guess_lambda_arn" {
  value = aws_lambda_function.sepoy_guess_lambda.arn
}


output "end_game_lambda_arn" {
  value = aws_lambda_function.end_game_lambda.arn
}

output "get_leaderboard_lambda_arn" {
  value = aws_lambda_function.get_leaderboard_lambda.arn
}


# Output for API Gateway URL
output "api_gateway_url" {
  value = aws_api_gateway_stage.api_stage.invoke_url
}

# Cognito User Pool and Client ID
output "user_pool_id" {
  value = aws_cognito_user_pool.kingsmen_guard_user_pool.id
}

output "client_id" {
  value = aws_cognito_user_pool_client.userpool_client.id
}

# Cognito Domain
output "cognito_domain" {
  value = aws_cognito_user_pool_domain.kingsmen_domain.domain
}
