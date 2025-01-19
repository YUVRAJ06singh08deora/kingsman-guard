# 1. Games Table: Stores metadata about each game
resource "aws_dynamodb_table" "games_table" {
  name           = "Games"
  billing_mode   = "PROVISIONED"
  hash_key       = "game_code"
  read_capacity  = 5
  write_capacity = 5

  

  attribute {
    name = "game_code"
    type = "S"
  }

  tags = {
    Name = "Games Table"
  }
}

# 2. Players Table: Stores player details in each game
resource "aws_dynamodb_table" "players_table" {
  name           = "Players"
  billing_mode   = "PROVISIONED"
  hash_key       = "player_id"
  range_key      = "game_code"
  read_capacity  = 5
  write_capacity = 5

  attribute {
    name = "player_id"
    type = "S"
  }

  attribute {
    name = "game_code"
    type = "S"
  }

  attribute {
    name = "role"
    type = "S"
  }

  global_secondary_index {
    name               = "RoleIndex"
    hash_key           = "game_code"
    range_key          = "role"
    read_capacity      = 5
    write_capacity     = 5
    projection_type    = "ALL"
  }

  tags = {
    Name = "Players Table"
  }
}

