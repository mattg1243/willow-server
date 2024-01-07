package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtPayload struct {
	Id string `json:"id"`
	Email string `json:"email"`
	Exp string `json:"exp"`
	jwt.Claims
}

func GenerateJWT(p JwtPayload) (string, error) {
	secretString := os.Getenv("JWT_SECRET")
	secretKey := []byte(secretString)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = p.Id
	claims["email"] = p.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	return token.SignedString(secretKey);
}

func ValidateJWT (tokenString string) (JwtPayload, error) {
  // validate the hashing algorithm
	token, err := jwt.ParseWithClaims(tokenString, &JwtPayload{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(tokenString), nil
	})
	
	if err != nil {
		fmt.Printf("error: %v", err.Error());
		return JwtPayload{}, err
	}

	if claims, ok := token.Claims.(*JwtPayload); ok {
		fmt.Println(claims)
		return JwtPayload{}, nil;
	} else {
		fmt.Printf("error: %v ", err.Error())
		return JwtPayload{}, nil
	}
}

func ParseUserIdFromClaims (claims interface{}) (uuid.UUID, error) {
	claimsParsed, ok := claims.(*JwtPayload)
	if !ok {
		return uuid.UUID{}, errors.New("Claims could not be parsed into a JwtPayload struct")
	}

	return uuid.Parse(claimsParsed.Id)
}