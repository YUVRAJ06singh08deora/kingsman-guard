resource "aws_cognito_user_pool" "kingsmen_guard_user_pool" {

  name                = "kingsmen_guard_user_pool"
  username_attributes = ["email"]
  auto_verified_attributes = ["email"]

  username_configuration {
    case_sensitive = false
  }



  schema {
    attribute_data_type = "String"
    mutable             = true
    name                = "email"
    required            = true
    string_attribute_constraints {
      min_length = 1
      max_length = 2048
    }
  }

  schema {
    attribute_data_type = "String"
    mutable             = true
    name                = "given_name"
    required            = true
    string_attribute_constraints {
      min_length = 1
      max_length = 2048
    }
  }

  schema {
    attribute_data_type = "String"
    mutable             = true
    name                = "family_name"
    required            = true
    string_attribute_constraints {
      min_length = 1
      max_length = 2048
    }
  }


  schema {
    attribute_data_type      = "String"
    developer_only_attribute = false
    mutable                  = true
    name                     = "PersonalInfo"
    required                 = false
    string_attribute_constraints {
      min_length = 1
      max_length = 2048
    }
  }

  schema {
    attribute_data_type      = "String"
    developer_only_attribute = false
    mutable                  = true
    name                     = "Role"
    required                 = false
    string_attribute_constraints {
      min_length = 1
      max_length = 2048
    }
  }
  

  # Email Verification Configuration
  email_verification_message = "Your verification code is {####}"
  email_verification_subject = "Verify your email address"

  account_recovery_setting {
    recovery_mechanism {
      name     = "verified_email"
      priority = 1
    }
  }
}

resource "aws_cognito_user_pool_client" "userpool_client" {
  name                                 = "kingmenclient"
  user_pool_id                         = aws_cognito_user_pool.kingsmen_guard_user_pool.id
  callback_urls                        = ["http://localhost:5173/authenticate"]
  logout_urls                          = ["http://localhost:5173/"]
  allowed_oauth_flows_user_pool_client = true
  allowed_oauth_flows                  = ["code", "implicit"]
  allowed_oauth_scopes                 = ["email", "openid","aws.cognito.signin.user.admin"]
  supported_identity_providers         = ["COGNITO"]
  access_token_validity = 1
  id_token_validity = 1
  refresh_token_validity = 2

  token_validity_units {
    access_token = "hours"
    id_token = "hours"
    refresh_token = "hours"
  } 
  
}


resource "aws_cognito_user_pool_domain" "kingsmen_domain" {
  domain       = "kingsmenauthentication"
  user_pool_id = aws_cognito_user_pool.kingsmen_guard_user_pool.id
}

resource "aws_cognito_user_pool_ui_customization" "kingsmen_ui_customization" {
  client_id = aws_cognito_user_pool_client.userpool_client.id

  css        = ".label-customizable {font-weight: 400;}"
  image_file = filebase64("lock.jpg")

  
  user_pool_id = aws_cognito_user_pool_domain.kingsmen_domain.user_pool_id
}