variable "service_name" {
  description = "Name of the service"
  type        = string
}

variable "cost_center" {
  description = "Cost center of the service"
  type        = string
}

variable "lambda_function_name" {
  description = "Name of the Lambda function"
  type        = string
}

variable "jwt_secret" {
  description = "JWT secret"
  type        = string
}

variable "email_from" {
  description = "Email from"
  type        = string
}

variable "email_password" {
  description = "Email password"
  type        = string
}

variable "email_smtp_host" {
  description = "Email SMTP host"
  type        = string
}

variable "email_smtp_port" {
  description = "Email SMTP port"
  type        = number
}

variable "api_name" {
  description = "Name of the API Gateway"
  type        = string
}

variable "dynamodb_table_name" {
  description = "Name of the DynamoDB table"
  type        = string
}

variable "dynamodb_global_secondary_index_name" {
  description = "Name of the DynamoDB global secondary index"
  type        = string
}

variable "lambda_dynamodb_policy_name" {
  description = "Name of the Lambda DynamoDB policy"
  type        = string
}

variable "api_domain_name" {
  description = "Domain name of the API Gateway"
  type        = string
}

variable "domain_name" {
  description = "Domain name of the API Gateway"
  type        = string
}

variable "hosted_zone_id" {
  description = "Hosted zone ID of the domain"
  type        = string
}
