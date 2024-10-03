provider "aws" {
  region = "ap-southeast-2"

  default_tags {
    tags = {
      ServiceName = var.service_name
      CostCenter  = var.cost_center
    }
  }
}

locals {
  lambda_runtime       = "provided.al2023" 
  lambda_handler       = "hello.handler"
  lambda_zip_file      = "deployment.zip"
}


# IAM Role for Lambda
resource "aws_iam_role" "lambda_exec_role" {
  name = "lambda_exec_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action    = "sts:AssumeRole"
      Effect    = "Allow"
      Principal = {
        Service = "lambda.amazonaws.com"
      }
    }]
  })
}

# Lambda Cloudwatch IAM Policy
resource "aws_iam_policy" "cloudwatch_policy" {
  name = "lambda_cloudwatch_logs_policy"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ]
        Effect   = "Allow"
        Resource = "arn:aws:logs:*:*:*"
      }
    ]
  })
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
      JWT_SECRET = var.jwt_secret
      JWT_EXPIRES_IN_N_HOURS = 24
      DDB_TABLE_NAME = var.dynamodb_table_name
      DDB_GSI_NAME= var.dynamodb_global_secondary_index_name

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

# Attach the Cloudwatch policy to the Lambda IAM role
resource "aws_iam_role_policy_attachment" "lambda_cloudwatch_logs_attach" {
  role       = aws_iam_role.lambda_exec_role.name
  policy_arn = aws_iam_policy.cloudwatch_policy.arn
}

# Create the DynamoDB Table
resource "aws_dynamodb_table" "my_table" {
  name           = var.dynamodb_table_name
  billing_mode   = "PAY_PER_REQUEST"
  hash_key       = "PK"
  range_key      = "SK"

  attribute {
    name = "PK"
    type = "S"
  }
  attribute {
    name = "SK"
    type = "S"
  }
  attribute {
    name = "GSIPK"
    type = "S"
  }
  attribute {
    name = "GSISK"
    type = "S"
  }

  global_secondary_index {
    name               = var.dynamodb_global_secondary_index_name
    hash_key           = "GSIPK"
    range_key          = "GSISK"
    projection_type    = "ALL"
  }
}

# Create a custom policy for accessing the specific DynamoDB table
resource "aws_iam_policy" "lambda_dynamodb_policy" {
  name        = var.lambda_dynamodb_policy_name
  
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect   = "Allow"
        Action   = [
          "dynamodb:GetItem",
          "dynamodb:PutItem",
          "dynamodb:UpdateItem",
          "dynamodb:DeleteItem",
          "dynamodb:Query",
        ]
        Resource = concat(
          [aws_dynamodb_table.my_table.arn],
          [for gsi in aws_dynamodb_table.my_table.global_secondary_index : "${aws_dynamodb_table.my_table.arn}/index/${gsi.name}"]
        )
      }
    ]
  })
}

# Attach the policy to the Lambda IAM role
resource "aws_iam_role_policy_attachment" "lambda_dynamodb_policy_attachment" {
  role       = aws_iam_role.lambda_exec_role.name
  policy_arn = aws_iam_policy.lambda_dynamodb_policy.arn
}

# API Gateway REST API
resource "aws_apigatewayv2_api" "my_api" {
  name          = var.api_name
  protocol_type = "HTTP"
}

# API Gateway integration with Lambda
resource "aws_apigatewayv2_integration" "lambda_integration" {
  api_id             = aws_apigatewayv2_api.my_api.id
  integration_type   = "AWS_PROXY"
  integration_uri    = aws_lambda_function.my_lambda.invoke_arn
  integration_method = "POST"
  payload_format_version = "2.0"
}

# API Gateway route
resource "aws_apigatewayv2_route" "default_route" {
  api_id    = aws_apigatewayv2_api.my_api.id
  route_key = "ANY /{proxy+}"
  target = "integrations/${aws_apigatewayv2_integration.lambda_integration.id}"
}

# Default stage for the API Gateway
resource "aws_apigatewayv2_stage" "default_stage" {
  api_id      = aws_apigatewayv2_api.my_api.id
  name        = "$default"
  auto_deploy = true
}

# Grant API Gateway permission to invoke Lambda
resource "aws_lambda_permission" "api_gateway_lambda" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.my_lambda.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_apigatewayv2_api.my_api.execution_arn}/*"
}

# Obtain an SSL certificate from AWS ACM
resource "aws_acm_certificate_validation" "cert_validation" {
  certificate_arn         = aws_acm_certificate.api_cert.arn
  
  # Iterate over the domain validation options and get the fqdn for each validation record
  validation_record_fqdns = [
    for dvo in aws_acm_certificate.api_cert.domain_validation_options : aws_route53_record.validation[dvo.domain_name].fqdn
  ]
}

# Custom domain for API Gateway
resource "aws_apigatewayv2_domain_name" "custom_domain" {
  domain_name = var.api_domain_name
  domain_name_configuration {
    certificate_arn = aws_acm_certificate.api_cert.arn
    endpoint_type   = "REGIONAL"
    security_policy = "TLS_1_2"  
  }
  depends_on = [aws_acm_certificate_validation.cert_validation]
}

# Route 53 hosted zone for the custom domain
resource "aws_route53_zone" "hosted_zone" {
  name = var.domain_name
}

# Route 53 DNS record for the custom domain
resource "aws_route53_record" "custom_domain" {
  zone_id = aws_route53_zone.hosted_zone.zone_id
  name    = var.api_domain_name
  type    = "A"
  alias {
    name                   = aws_apigatewayv2_domain_name.custom_domain.domain_name_configuration[0].target_domain_name
    zone_id                = aws_apigatewayv2_domain_name.custom_domain.domain_name_configuration[0].hosted_zone_id
    evaluate_target_health = false
  }
}

# Base path mapping for the custom domain
resource "aws_apigatewayv2_api_mapping" "custom_domain_mapping" {
  api_id      = aws_apigatewayv2_api.my_api.id
  domain_name = aws_apigatewayv2_domain_name.custom_domain.domain_name
  stage       = aws_apigatewayv2_stage.default_stage.id
}


# Obtain an SSL certificate from AWS ACM
resource "aws_acm_certificate" "api_cert" {
  domain_name       = var.api_domain_name
  validation_method = "DNS"

}

# DNS validation record using Route 53
resource "aws_route53_record" "validation" {
  for_each = {
    for dvo in aws_acm_certificate.api_cert.domain_validation_options : dvo.domain_name => {
      name  = dvo.resource_record_name
      type  = dvo.resource_record_type
      value = dvo.resource_record_value
    }
  }

  name    = each.value.name
  type    = each.value.type
  zone_id = aws_route53_zone.hosted_zone.zone_id
  records = [each.value.value]
  ttl     = 60
}
