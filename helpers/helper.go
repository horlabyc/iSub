package helper

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type SignedDetails struct {
	Email    string
	Username string
	UserId   primitive.ObjectID
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

func GenerateToken(username *string, email *string, userId *primitive.ObjectID) (string, string) {
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

func VerifyPassword(hashedPassword string, password string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	valid := true
	errorMsg := ""
	if err != nil {
		errorMsg = fmt.Sprintf("email or password is incorrect")
		valid = false
	}
	return valid, errorMsg
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(signedToken, &SignedDetails{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		msg = err.Error()
	}
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("provided token is invalid")
		return nil, msg
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is expired")
		return nil, msg
	}
	return claims, msg
}

func ConvertToObjectId(str string) (primitive.ObjectID, error) {
	oid, err := primitive.ObjectIDFromHex(str)
	return oid, err
}
