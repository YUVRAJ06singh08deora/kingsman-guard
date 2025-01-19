# Create API Gateway REST API
resource "aws_api_gateway_rest_api" "api_gateway" {
  name = "GameAPIGateway"
  endpoint_configuration {
    types = ["REGIONAL"]
  }
}

# Cognito Authorizer
resource "aws_api_gateway_authorizer" "authorizer" {
  name          = "gameAuthorizer"
  type          = "COGNITO_USER_POOLS"
  rest_api_id   = aws_api_gateway_rest_api.api_gateway.id
  provider_arns = ["${aws_cognito_user_pool.kingsmen_guard_user_pool.arn}"]
}

# Create resource for create_game endpoint
resource "aws_api_gateway_resource" "create_game_resource" {
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  parent_id   = aws_api_gateway_rest_api.api_gateway.root_resource_id
  path_part   = "create_game"
}

# Create POST method for create_game endpoint
resource "aws_api_gateway_method" "create_game_method" {
  rest_api_id          = aws_api_gateway_rest_api.api_gateway.id
  resource_id          = aws_api_gateway_resource.create_game_resource.id
  http_method          = "POST"
  authorization        = "COGNITO_USER_POOLS"
  authorization_scopes = ["openid","aws.cognito.signin.user.admin"]
  authorizer_id        = aws_api_gateway_authorizer.authorizer.id
}

# Integration for create_game endpoint with Lambda
resource "aws_api_gateway_integration" "create_game_integration" {
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  resource_id = aws_api_gateway_resource.create_game_resource.id
  http_method = aws_api_gateway_method.create_game_method.http_method

  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.create_game_lambda.invoke_arn
}

# CORS configuration for create_game
resource "aws_api_gateway_method" "create_game_cors" {
  rest_api_id   = aws_api_gateway_rest_api.api_gateway.id
  resource_id   = aws_api_gateway_resource.create_game_resource.id
  http_method   = "OPTIONS"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "create_game_cors_integration" {
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  resource_id = aws_api_gateway_resource.create_game_resource.id
  http_method = aws_api_gateway_method.create_game_cors.http_method

  integration_http_method = "POST"
  type                    = "MOCK"

  request_templates = {
    "application/json" = "{\"statusCode\": 200}"
  }

  content_handling = "CONVERT_TO_TEXT"
}

resource "aws_api_gateway_method_response" "create_game_cors_response200" {
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  resource_id = aws_api_gateway_resource.create_game_resource.id
  http_method = aws_api_gateway_method.create_game_cors.http_method
  status_code = "200"

  response_models = {
    "application/json" = "Empty"
  }

  response_parameters = {
    "method.response.header.Access-Control-Allow-Headers" = true
    "method.response.header.Access-Control-Allow-Methods" = true
    "method.response.header.Access-Control-Allow-Origin"  = true
  }
}

resource "aws_api_gateway_integration_response" "create_game_cors_integration_response" {
  depends_on = [
    aws_api_gateway_integration.create_game_cors_integration
  ]
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  resource_id = aws_api_gateway_resource.create_game_resource.id
  http_method = aws_api_gateway_method.create_game_cors.http_method
  status_code = "200"

  response_parameters = {
    "method.response.header.Access-Control-Allow-Headers" = "'Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token'",
    "method.response.header.Access-Control-Allow-Methods" = "'POST,OPTIONS'",
    "method.response.header.Access-Control-Allow-Origin"  = "'*'"
  }

  response_templates = {
    "application/json" = ""
  }
}

# Create resource for join_game endpoint
resource "aws_api_gateway_resource" "join_game_resource" {
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  parent_id   = aws_api_gateway_rest_api.api_gateway.root_resource_id
  path_part   = "join_game"
}

# Create POST method for join_game endpoint
resource "aws_api_gateway_method" "join_game_method" {
  rest_api_id          = aws_api_gateway_rest_api.api_gateway.id
  resource_id          = aws_api_gateway_resource.join_game_resource.id
  http_method          = "POST"
  authorization        = "COGNITO_USER_POOLS"
  authorization_scopes = ["openid","aws.cognito.signin.user.admin"]
  authorizer_id        = aws_api_gateway_authorizer.authorizer.id
}

# Integration for join_game endpoint with Lambda
resource "aws_api_gateway_integration" "join_game_integration" {
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  resource_id = aws_api_gateway_resource.join_game_resource.id
  http_method = aws_api_gateway_method.join_game_method.http_method

  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.join_game_lambda.invoke_arn
}

# CORS configuration for join_game
resource "aws_api_gateway_method" "join_game_cors" {
  rest_api_id   = aws_api_gateway_rest_api.api_gateway.id
  resource_id   = aws_api_gateway_resource.join_game_resource.id
  http_method   = "OPTIONS"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "join_game_cors_integration" {
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  resource_id = aws_api_gateway_resource.join_game_resource.id
  http_method = aws_api_gateway_method.join_game_cors.http_method

  integration_http_method = "POST"
  type                    = "MOCK"

  request_templates = {
    "application/json" = "{\"statusCode\": 200}"
  }

  content_handling = "CONVERT_TO_TEXT"
}

resource "aws_api_gateway_method_response" "join_game_cors_response200" {
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  resource_id = aws_api_gateway_resource.join_game_resource.id
  http_method = aws_api_gateway_method.join_game_cors.http_method
  status_code = "200"

  response_models = {
    "application/json" = "Empty"
  }

  response_parameters = {
    "method.response.header.Access-Control-Allow-Headers" = true
    "method.response.header.Access-Control-Allow-Methods" = true
    "method.response.header.Access-Control-Allow-Origin"  = true
  }
}

resource "aws_api_gateway_integration_response" "join_game_cors_integration_response" {
  depends_on = [
    aws_api_gateway_integration.join_game_cors_integration
  ]
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  resource_id = aws_api_gateway_resource.join_game_resource.id
  http_method = aws_api_gateway_method.join_game_cors.http_method
  status_code = "200"

  response_parameters = {
    "method.response.header.Access-Control-Allow-Headers" = "'Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token'",
    "method.response.header.Access-Control-Allow-Methods" = "'POST,OPTIONS'",
    "method.response.header.Access-Control-Allow-Origin"  = "'*'"
  }

  response_templates = {
    "application/json" = ""
  }
}


# Create resource for start_game endpoint
resource "aws_api_gateway_resource" "start_game_resource" {
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  parent_id   = aws_api_gateway_rest_api.api_gateway.root_resource_id
  path_part   = "start_game"
}

# Create POST method for start_game endpoint
resource "aws_api_gateway_method" "start_game_method" {
  rest_api_id          = aws_api_gateway_rest_api.api_gateway.id
  resource_id          = aws_api_gateway_resource.start_game_resource.id
  http_method          = "POST"
  authorization        = "COGNITO_USER_POOLS"
  authorization_scopes = ["openid","aws.cognito.signin.user.admin"]
  authorizer_id        = aws_api_gateway_authorizer.authorizer.id
}

# Integration for start_game endpoint with Lambda
resource "aws_api_gateway_integration" "start_game_integration" {
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  resource_id = aws_api_gateway_resource.start_game_resource.id
  http_method = aws_api_gateway_method.start_game_method.http_method

  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.start_game_lambda.invoke_arn
}

# CORS configuration for start_game
resource "aws_api_gateway_method" "start_game_cors" {
  rest_api_id   = aws_api_gateway_rest_api.api_gateway.id
  resource_id   = aws_api_gateway_resource.start_game_resource.id
  http_method   = "OPTIONS"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "start_game_cors_integration" {
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  resource_id = aws_api_gateway_resource.start_game_resource.id
  http_method = aws_api_gateway_method.start_game_cors.http_method

  integration_http_method = "POST"
  type                    = "MOCK"

  request_templates = {
    "application/json" = "{\"statusCode\": 200}"
  }

  content_handling = "CONVERT_TO_TEXT"
}

resource "aws_api_gateway_method_response" "start_game_cors_response200" {
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  resource_id = aws_api_gateway_resource.start_game_resource.id
  http_method = aws_api_gateway_method.start_game_cors.http_method
  status_code = "200"

  response_models = {
    "application/json" = "Empty"
  }

  response_parameters = {
    "method.response.header.Access-Control-Allow-Headers" = true
    "method.response.header.Access-Control-Allow-Methods" = true
    "method.response.header.Access-Control-Allow-Origin"  = true
  }
}

resource "aws_api_gateway_integration_response" "start_game_cors_integration_response" {
  depends_on = [
    aws_api_gateway_integration.start_game_cors_integration
  ]
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  resource_id = aws_api_gateway_resource.start_game_resource.id
  http_method = aws_api_gateway_method.start_game_cors.http_method
  status_code = "200"

  response_parameters = {
    "method.response.header.Access-Control-Allow-Headers" = "'Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token'",
    "method.response.header.Access-Control-Allow-Methods" = "'POST,OPTIONS'",
    "method.response.header.Access-Control-Allow-Origin"  = "'*'"
  }

  response_templates = {
    "application/json" = ""
  }
}

# Create resource for assign_roles endpoint
resource "aws_api_gateway_resource" "assign_roles_resource" {
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  parent_id   = aws_api_gateway_rest_api.api_gateway.root_resource_id
  path_part   = "assign_roles"
}

# Create POST method for assign_roles endpoint
resource "aws_api_gateway_method" "assign_roles_method" {
  rest_api_id          = aws_api_gateway_rest_api.api_gateway.id
  resource_id          = aws_api_gateway_resource.assign_roles_resource.id
  http_method          = "POST"
  authorization        = "COGNITO_USER_POOLS"
  authorization_scopes = ["openid","aws.cognito.signin.user.admin"]
  authorizer_id        = aws_api_gateway_authorizer.authorizer.id
}

# Integration for assign_roles endpoint with Lambda
resource "aws_api_gateway_integration" "assign_roles_integration" {
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  resource_id = aws_api_gateway_resource.assign_roles_resource.id
  http_method = aws_api_gateway_method.assign_roles_method.http_method

  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.assign_roles_lambda.invoke_arn
}

# CORS configuration for assign_roles
resource "aws_api_gateway_method" "assign_roles_cors" {
  rest_api_id   = aws_api_gateway_rest_api.api_gateway.id
  resource_id   = aws_api_gateway_resource.assign_roles_resource.id
  http_method   = "OPTIONS"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "assign_roles_cors_integration" {
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  resource_id = aws_api_gateway_resource.assign_roles_resource.id
  http_method = aws_api_gateway_method.assign_roles_cors.http_method

  integration_http_method = "POST"
  type                    = "MOCK"

  request_templates = {
    "application/json" = "{\"statusCode\": 200}"
  }

  content_handling = "CONVERT_TO_TEXT"
}

resource "aws_api_gateway_method_response" "assign_roles_cors_response200" {
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  resource_id = aws_api_gateway_resource.assign_roles_resource.id
  http_method = aws_api_gateway_method.assign_roles_cors.http_method
  status_code = "200"

  response_models = {
    "application/json" = "Empty"
  }

  response_parameters = {
    "method.response.header.Access-Control-Allow-Headers" = true
    "method.response.header.Access-Control-Allow-Methods" = true
    "method.response.header.Access-Control-Allow-Origin"  = true
  }
}

resource "aws_api_gateway_integration_response" "assign_roles_cors_integration_response" {
  depends_on = [
    aws_api_gateway_integration.assign_roles_cors_integration
  ]
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  resource_id = aws_api_gateway_resource.assign_roles_resource.id
  http_method = aws_api_gateway_method.assign_roles_cors.http_method
  status_code = "200"

  response_parameters = {
    "method.response.header.Access-Control-Allow-Headers" = "'Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token'",
    "method.response.header.Access-Control-Allow-Methods" = "'POST,OPTIONS'",
    "method.response.header.Access-Control-Allow-Origin"  = "'*'"
  }

  response_templates = {
    "application/json" = ""
  }
}

# Create resource for sepoy_guess endpoint
resource "aws_api_gateway_resource" "sepoy_guess_resource" {
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  parent_id   = aws_api_gateway_rest_api.api_gateway.root_resource_id
  path_part   = "sepoy_guess"
}

# Create POST method for sepoy_guess endpoint
resource "aws_api_gateway_method" "sepoy_guess_method" {
  rest_api_id          = aws_api_gateway_rest_api.api_gateway.id
  resource_id          = aws_api_gateway_resource.sepoy_guess_resource.id
  http_method          = "POST"
  authorization        = "COGNITO_USER_POOLS"
  authorization_scopes = ["openid","aws.cognito.signin.user.admin"]
  authorizer_id        = aws_api_gateway_authorizer.authorizer.id
}

# Integration for sepoy_guess endpoint with Lambda
resource "aws_api_gateway_integration" "sepoy_guess_integration" {
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  resource_id = aws_api_gateway_resource.sepoy_guess_resource.id
  http_method = aws_api_gateway_method.sepoy_guess_method.http_method

  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.sepoy_guess_lambda.invoke_arn
}

# CORS configuration for sepoy_guess
resource "aws_api_gateway_method" "sepoy_guess_cors" {
  rest_api_id   = aws_api_gateway_rest_api.api_gateway.id
  resource_id   = aws_api_gateway_resource.sepoy_guess_resource.id
  http_method   = "OPTIONS"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "sepoy_guess_cors_integration" {
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  resource_id = aws_api_gateway_resource.sepoy_guess_resource.id
  http_method = aws_api_gateway_method.sepoy_guess_cors.http_method

  integration_http_method = "POST"
  type                    = "MOCK"

  request_templates = {
    "application/json" = "{\"statusCode\": 200}"
  }

  content_handling = "CONVERT_TO_TEXT"
}

resource "aws_api_gateway_method_response" "sepoy_guess_cors_response200" {
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  resource_id = aws_api_gateway_resource.sepoy_guess_resource.id
  http_method = aws_api_gateway_method.sepoy_guess_cors.http_method
  status_code = "200"

  response_models = {
    "application/json" = "Empty"
  }

  response_parameters = {
    "method.response.header.Access-Control-Allow-Headers" = true
    "method.response.header.Access-Control-Allow-Methods" = true
    "method.response.header.Access-Control-Allow-Origin"  = true
  }
}

resource "aws_api_gateway_integration_response" "sepoy_guess_cors_integration_response" {
  depends_on = [
    aws_api_gateway_integration.sepoy_guess_cors_integration
  ]
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  resource_id = aws_api_gateway_resource.sepoy_guess_resource.id
  http_method = aws_api_gateway_method.sepoy_guess_cors.http_method
  status_code = "200"

  response_parameters = {
    "method.response.header.Access-Control-Allow-Headers" = "'Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token'",
    "method.response.header.Access-Control-Allow-Methods" = "'POST,OPTIONS'",
    "method.response.header.Access-Control-Allow-Origin"  = "'*'"
  }

  response_templates = {
    "application/json" = ""
  }
}


# Create resource for end_game endpoint
resource "aws_api_gateway_resource" "end_game_resource" {
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  parent_id   = aws_api_gateway_rest_api.api_gateway.root_resource_id
  path_part   = "end_game"
}

# Create POST method for end_game endpoint
resource "aws_api_gateway_method" "end_game_method" {
  rest_api_id          = aws_api_gateway_rest_api.api_gateway.id
  resource_id          = aws_api_gateway_resource.end_game_resource.id
  http_method          = "POST"
  authorization        = "COGNITO_USER_POOLS"
  authorization_scopes = ["openid","aws.cognito.signin.user.admin"]
  authorizer_id        = aws_api_gateway_authorizer.authorizer.id
}

# Integration for end_game endpoint with Lambda
resource "aws_api_gateway_integration" "end_game_integration" {
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  resource_id = aws_api_gateway_resource.end_game_resource.id
  http_method = aws_api_gateway_method.end_game_method.http_method

  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.end_game_lambda.invoke_arn
}

# CORS configuration for end_game
resource "aws_api_gateway_method" "end_game_cors" {
  rest_api_id   = aws_api_gateway_rest_api.api_gateway.id
  resource_id   = aws_api_gateway_resource.end_game_resource.id
  http_method   = "OPTIONS"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "end_game_cors_integration" {
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  resource_id = aws_api_gateway_resource.end_game_resource.id
  http_method = aws_api_gateway_method.end_game_cors.http_method

  integration_http_method = "POST"
  type                    = "MOCK"

  request_templates = {
    "application/json" = "{\"statusCode\": 200}"
  }

  content_handling = "CONVERT_TO_TEXT"
}

resource "aws_api_gateway_method_response" "end_game_cors_response200" {
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  resource_id = aws_api_gateway_resource.end_game_resource.id
  http_method = aws_api_gateway_method.end_game_cors.http_method
  status_code = "200"

  response_models = {
    "application/json" = "Empty"
  }

  response_parameters = {
    "method.response.header.Access-Control-Allow-Headers" = true
    "method.response.header.Access-Control-Allow-Methods" = true
    "method.response.header.Access-Control-Allow-Origin"  = true
  }
}

resource "aws_api_gateway_integration_response" "end_game_cors_integration_response" {
  depends_on = [
    aws_api_gateway_integration.end_game_cors_integration
  ]
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  resource_id = aws_api_gateway_resource.end_game_resource.id
  http_method = aws_api_gateway_method.end_game_cors.http_method
  status_code = "200"

  response_parameters = {
    "method.response.header.Access-Control-Allow-Headers" = "'Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token'",
    "method.response.header.Access-Control-Allow-Methods" = "'POST,OPTIONS'",
    "method.response.header.Access-Control-Allow-Origin"  = "'*'"
  }

  response_templates = {
    "application/json" = ""
  }
}

# Create Resource for get_leaderboard endpoint
resource "aws_api_gateway_resource" "get_leaderboard_resource" {
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  parent_id   = aws_api_gateway_rest_api.api_gateway.root_resource_id
  path_part   = "get_leaderboard"
}

# Create GET Method for get_leaderboard endpoint
resource "aws_api_gateway_method" "get_leaderboard_method" {
  rest_api_id          = aws_api_gateway_rest_api.api_gateway.id
  resource_id          = aws_api_gateway_resource.get_leaderboard_resource.id
  http_method          = "GET"
  authorization        = "COGNITO_USER_POOLS"
  authorization_scopes = ["openid", "aws.cognito.signin.user.admin"]
  authorizer_id        = aws_api_gateway_authorizer.authorizer.id
}

# Integration for get_leaderboard endpoint with Lambda
resource "aws_api_gateway_integration" "get_leaderboard_integration" {
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  resource_id = aws_api_gateway_resource.get_leaderboard_resource.id
  http_method = aws_api_gateway_method.get_leaderboard_method.http_method

  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.get_leaderboard_lambda.invoke_arn
}

# CORS Configuration for get_leaderboard

# CORS Method
resource "aws_api_gateway_method" "get_leaderboard_cors" {
  rest_api_id   = aws_api_gateway_rest_api.api_gateway.id
  resource_id   = aws_api_gateway_resource.get_leaderboard_resource.id
  http_method   = "OPTIONS"
  authorization = "NONE"
}

# CORS Integration
resource "aws_api_gateway_integration" "get_leaderboard_cors_integration" {
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  resource_id = aws_api_gateway_resource.get_leaderboard_resource.id
  http_method = aws_api_gateway_method.get_leaderboard_cors.http_method

  integration_http_method = "POST"
  type                    = "MOCK"

  request_templates = {
    "application/json" = "{\"statusCode\": 200}"
  }

  content_handling = "CONVERT_TO_TEXT"
}

# CORS Method Response (200 OK)
resource "aws_api_gateway_method_response" "get_leaderboard_cors_response200" {
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  resource_id = aws_api_gateway_resource.get_leaderboard_resource.id
  http_method = aws_api_gateway_method.get_leaderboard_cors.http_method
  status_code = "200"

  response_models = {
    "application/json" = "Empty"
  }

  response_parameters = {
    "method.response.header.Access-Control-Allow-Headers" = true
    "method.response.header.Access-Control-Allow-Methods" = true
    "method.response.header.Access-Control-Allow-Origin"  = true
  }
}

# CORS Integration Response (200 OK)
resource "aws_api_gateway_integration_response" "get_leaderboard_cors_integration_response" {
  depends_on = [
    aws_api_gateway_integration.get_leaderboard_cors_integration
  ]
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  resource_id = aws_api_gateway_resource.get_leaderboard_resource.id
  http_method = aws_api_gateway_method.get_leaderboard_cors.http_method
  status_code = "200"

  response_parameters = {
    "method.response.header.Access-Control-Allow-Headers" = "'Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token'",
    "method.response.header.Access-Control-Allow-Methods" = "'POST,OPTIONS'",
    "method.response.header.Access-Control-Allow-Origin"  = "'*'"
  }

  response_templates = {
    "application/json" = ""
  }
}


resource "aws_api_gateway_deployment" "game_api_deployment" {

  triggers = {
    redeployment = sha1(jsonencode([
      aws_api_gateway_resource.create_game_resource.id,
      aws_api_gateway_resource.join_game_resource.id,
      aws_api_gateway_resource.start_game_resource.id,
      aws_api_gateway_resource.assign_roles_resource.id,
      aws_api_gateway_resource.sepoy_guess_resource.id,
      aws_api_gateway_resource.end_game_resource.id,
      aws_api_gateway_resource.get_leaderboard_resource.id,
     

      aws_api_gateway_method.create_game_method.id,
      aws_api_gateway_method.join_game_method.id,
      aws_api_gateway_method.start_game_method.id,
      aws_api_gateway_method.assign_roles_method.id,
      aws_api_gateway_method.sepoy_guess_method.id,
      aws_api_gateway_method.end_game_method.id,
      aws_api_gateway_method.get_leaderboard_method.id,
    
      aws_api_gateway_integration.create_game_integration,
      aws_api_gateway_integration.join_game_integration,
      aws_api_gateway_integration.start_game_integration,
      aws_api_gateway_integration.assign_roles_integration,
      aws_api_gateway_integration.sepoy_guess_integration,
      aws_api_gateway_integration.end_game_integration,
      aws_api_gateway_integration.get_leaderboard_integration,

      aws_api_gateway_method.create_game_cors.id,
      aws_api_gateway_method.join_game_cors.id,
      aws_api_gateway_method.start_game_cors.id,
      aws_api_gateway_method.assign_roles_cors.id,
      aws_api_gateway_method.sepoy_guess_cors.id,
      aws_api_gateway_method.end_game_cors.id,
      aws_api_gateway_method.get_leaderboard_cors.id,

      aws_api_gateway_integration.create_game_cors_integration,
      aws_api_gateway_integration.join_game_cors_integration,
      aws_api_gateway_integration.start_game_cors_integration,
      aws_api_gateway_integration.assign_roles_cors_integration,
      aws_api_gateway_integration.sepoy_guess_cors_integration,
      aws_api_gateway_integration.end_game_cors_integration,
      aws_api_gateway_integration.get_leaderboard_cors_integration,

      aws_api_gateway_method_response.create_game_cors_response200.id,
      aws_api_gateway_method_response.join_game_cors_response200.id,
      aws_api_gateway_method_response.start_game_cors_response200.id,
      aws_api_gateway_method_response.assign_roles_cors_response200.id,
      aws_api_gateway_method_response.sepoy_guess_cors_response200.id,
      aws_api_gateway_method_response.end_game_cors_response200.id,
      aws_api_gateway_method_response.get_leaderboard_cors_response200.id,

      aws_api_gateway_integration_response.create_game_cors_integration_response.id,
      aws_api_gateway_integration_response.join_game_cors_integration_response.id,
      aws_api_gateway_integration_response.start_game_cors_integration_response.id,
      aws_api_gateway_integration_response.assign_roles_cors_integration_response.id,
      aws_api_gateway_integration_response.sepoy_guess_cors_integration_response.id,
      aws_api_gateway_integration_response.end_game_cors_integration_response.id,
      aws_api_gateway_integration_response.get_leaderboard_cors_integration_response.id
    ]))
  }
  
  rest_api_id = aws_api_gateway_rest_api.api_gateway.id
  
  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_api_gateway_stage" "api_stage" {
  deployment_id = aws_api_gateway_deployment.game_api_deployment.id
  rest_api_id   = aws_api_gateway_rest_api.api_gateway.id
  stage_name    = "prod"
}
