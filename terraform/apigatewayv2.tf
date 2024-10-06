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

# Custom domain for API Gateway
resource "aws_apigatewayv2_domain_name" "custom_domain" {
  domain_name = var.api_domain_name
  domain_name_configuration {
    certificate_arn = aws_acm_certificate.api_cert.arn
    endpoint_type   = "REGIONAL"
    security_policy = "TLS_1_2"  
  }
}

# Base path mapping for the custom domain
resource "aws_apigatewayv2_api_mapping" "custom_domain_mapping" {
  api_id      = aws_apigatewayv2_api.my_api.id
  domain_name = aws_apigatewayv2_domain_name.custom_domain.domain_name
  stage       = aws_apigatewayv2_stage.default_stage.id
}
