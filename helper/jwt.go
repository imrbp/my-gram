package helper

import (
	"MyGram/config"
	"MyGram/model/entity"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateUserToken(user entity.User) (string, error) {
	envConfig := config.LoadConfig()
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"userId":   user.Id,
		"username": user.Username,
		"email":    user.Email,
		"age":      user.Age,
		"iat":      time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(envConfig.JWT_SECRET_KEY))

	return tokenString, err
}

func ValidateUser(tokenString string) (user entity.User, err error) {
	envConfig := config.LoadConfig()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(envConfig.JWT_SECRET_KEY), nil
	})

	if err != nil {
		return user, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user.Id = int(claims["userId"].(float64))
		user.Email = claims["email"].(string)
		user.Username = claims["username"].(string)
		user.Age = int(claims["age"].(float64))
		return user, nil
	}

	return user, nil
}

func HashPassword(password string) (hashedPassword string, err error) {
	hashes := sha256.New()
	_, err = hashes.Write([]byte(password))
	if err != nil {
		return "", err
	}

	passwordHashBytes := hashes.Sum(nil)
	passwordHashString := hex.EncodeToString(passwordHashBytes)
	return passwordHashString, nil
}
