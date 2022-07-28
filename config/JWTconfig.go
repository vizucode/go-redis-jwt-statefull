package config

import (
	"os"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
)

var JWT_SECRET = []byte(os.Getenv("JWT_SECRET_KEY"))
var JWT_DURATION, _ = strconv.Atoi(os.Getenv("JWT_DURATION"))

// var JWT_SECRET = []byte("secret")
// var JWT_DURATION = 1

type JWTClaims struct {
	Name  string
	Email string
	jwt.RegisteredClaims
}
