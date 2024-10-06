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


# Attach the Cloudwatch policy to the Lambda IAM role
resource "aws_iam_role_policy_attachment" "lambda_cloudwatch_logs_attach" {
  role       = aws_iam_role.lambda_exec_role.name
  policy_arn = aws_iam_policy.cloudwatch_policy.arn
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

