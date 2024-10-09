package pkg

import (
	"fmt"
	"os"
	"time"

	"burger.local/models"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

// GenerateToken generates a new JWT token
func GenerateToken(user models.User) (string, error) {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"iss": "burger.local",
		"aud": "burger.local",
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	})

	fmt.Printf("Token claims added: %v\n", claims)

	token, er := claims.SignedString(secretKey)
	if er != nil {
		return "", er
	}

	return token, nil
}
