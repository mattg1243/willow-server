package utils

import (
	"errors"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type JwtPayload struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}


var secretString = os.Getenv("JWT_SECRET")

func GenerateJWT(p JwtPayload) (string, error) {
	
	secretKey := []byte(secretString)
	claims := &jwt.MapClaims{
			"id": p.Id,
			"email": p.Email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	return token.SignedString(secretKey);
}

func ValidateJWT (tokenString string) (*JwtPayload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtPayload{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretString), nil
	})
	fmt.Print(token)
	if err != nil {
		fmt.Printf("error: %v", err.Error());
		return &JwtPayload{}, err
	}

	if claims, ok := token.Claims.(*JwtPayload); ok && token.Valid {
		return claims, nil
	} else {
		return &JwtPayload{}, errors.New("Error decoding token; invalid token")
	}
}
