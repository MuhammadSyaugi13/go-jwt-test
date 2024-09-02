package config

import (
	"github.com/golang-jwt/jwt/v4"
)

var JWT_KEY = []byte("fsadf87aser329423j5lt930f9")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
