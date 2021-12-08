package helper

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type SignedDetails struct {
	Email    string
	Username string
	UserId   uint
	jwt.StandardClaims
}

func LoadEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading environment file")
	}
	return os.Getenv(key)
}

var secret = LoadEnv("SECRET_KEY")

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ComparePasswordHash(password, hash string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	valid := true
	errorMessage := ""
	if err != nil {
		errorMessage = fmt.Sprintf("Email or Password is incorrect")
		valid = false
	}
	return valid, errorMessage
}

func GenerateToken(username *string, email *string, userId *uint) (string, string) {
	tokenClaims := &SignedDetails{
		UserId:   *userId,
		Username: *username,
		Email:    *email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims).SignedString([]byte(secret))
	if err != nil {
		log.Panic(err)
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(secret))
	if err != nil {
		log.Panic(err)
	}
	return token, refreshToken
}
