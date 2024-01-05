package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtPayload struct {
	Id int32 `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
}

type JwtParsed struct {

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

func ValidateJWT (tokenString string) (interface{}, error) {
  // validate the hashing algorithm
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})
	
	if err != nil {
		fmt.Printf("error: %v", err.Error());
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println(claims)
		return claims, nil;
	} else {
		fmt.Printf("error: %v ", err.Error())
		return nil, nil
	}
}