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

