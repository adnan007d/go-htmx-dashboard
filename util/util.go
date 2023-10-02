package util

import (
	"context"
	"go-htmx-dashboard/internal/database"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateAUser(Db *database.Queries, email string, password string) (database.User, error) {


	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return database.User{}, err
	}

	user, err := Db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		Email:    email,
		Password: string(hash),
	})

	return user, err

}

type JWTClaims struct {
	Id string
	jwt.RegisteredClaims
}

func GenerateJWTToken(id string) (string, jwt.Claims, error) {
	key := []byte(os.Getenv("JWT_SECRET"))

	issuedTime := time.Now()
	expirationTime := issuedTime.Add(time.Hour * 24)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(issuedTime),
			ID:        id,
		},
	})

	tokenString, err := token.SignedString(key)

	return tokenString, token.Claims, err

}
