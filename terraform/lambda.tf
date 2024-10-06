locals {
  lambda_runtime       = "provided.al2023" 
  lambda_handler       = "hello.handler"
  lambda_zip_file      = "./../deployment.zip"
}

# Create Lambda function
resource "aws_lambda_function" "my_lambda" {
  function_name = var.lambda_function_name
  role          = aws_iam_role.lambda_exec_role.arn
  handler       = local.lambda_handler
  runtime       = local.lambda_runtime
  filename      = local.lambda_zip_file
  source_code_hash = filebase64sha256(local.lambda_zip_file)
  architectures = ["arm64"]

# Environment variables
  environment {
    variables = {
      JWT_SECRET = random_password.jwt_secret.result
      JWT_EXPIRES_IN_N_HOURS = 24
      DDB_TABLE_NAME = var.dynamodb_table_name
      DDB_GSI_NAME= var.dynamodb_global_secondary_index_name

      APP_LOGIN_URL = var.app_login_url

      EMAIL_FROM = var.email_from
      EMAIL_PASSWORD = var.email_password
      EMAIL_SMTP_HOST = var.email_smtp_host
      EMAIL_SMTP_PORT = var.email_smtp_port
    }
  }

  # Optional: timeout, memory, etc.
  timeout = 30
  memory_size = 128
}

# Grant API Gateway permission to invoke Lambda
resource "aws_lambda_permission" "api_gateway_lambda" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.my_lambda.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_apigatewayv2_api.my_api.execution_arn}/*"
}

