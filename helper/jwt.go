package helper

import (
	"MyGram/config"
	"MyGram/model/entity"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"math/rand"
	"time"
)

func GenerateUserToken(user entity.User) (string, error) {
	envConfig := config.LoadConfig()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.Id,
		"email":    user.Email,
		"username": user.Username,
		"age":      user.Age,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString([]byte(envConfig.JWT_SECRET_KEY))

	return tokenString, err
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

func ValidateUserToken(tokenString string) (user entity.UserReadJwt, err error) {
	envConfig := config.LoadConfig()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(envConfig.JWT_SECRET_KEY), nil
	})

	if err != nil {
		return user, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user.Id = int(claims["id"].(float64))
		user.Email = claims["email"].(string)
		user.Username = claims["username"].(string)
		user.Age = int(claims["age"].(float64))
		return user, nil
	} else if int64(claims["exp"].(float64)) < time.Now().Unix() {
		return user, fiber.NewError(fiber.StatusNotAcceptable, "Token is expired")
	} else if errors.Is(err, jwt.ErrInvalidKey) {
		return user, fiber.NewError(fiber.StatusNotAcceptable, "Invalid key")
	} else {
		return user, fiber.NewError(fiber.StatusNotAcceptable, "Invalid key")
	}
}

func AuthBearerValidation(ctx *fiber.Ctx, key string) (bool, error) {
	UserReadJwt, err := ValidateUserToken(key)
	if err != nil {
		return false, err
	}
	ctx.Locals("userRead", UserReadJwt)
	return true, nil
}

func GenerateSecretKey() {

	buf := make([]byte, 128)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}
	hx := hex.EncodeToString(buf)
	fmt.Println(" ==> " + hx)
}
