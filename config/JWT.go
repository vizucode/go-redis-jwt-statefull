package config

import "os"

var JWT_TOP_SECRET = []byte(os.Getenv("JWT_SECRET"))
