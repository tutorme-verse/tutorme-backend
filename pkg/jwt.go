package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"time"
)

// Secret key for signing JWTs (keep this secure in production)
var jwtKey = []byte("my_secret_key")

// Claims defines the structure of the JWT payload
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}
