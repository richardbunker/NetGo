# Generate a random string for JWT_SECRET
resource "random_password" "jwt_secret" {
  length  = 32
  special = false
}
