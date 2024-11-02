package token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tyagnii/gw-currency-wallet/internal/models"
	"time"
)

var Secret string
var ExpireTime time.Duration

type Claims struct {
	jwt.RegisteredClaims
	Username string
}

// Init package variables
func init() {
	// todo: read secret from configuration
	Secret = "very secret string"

	// todo: read expiration time from configuration
	ExpireTime = time.Hour * 24
}

// NewToken creates new JWT token for given user
func NewToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ExpireTime)),
		},
		Username: user.Username,
	})

	tokenString, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken checks if provided token is valid and
// return error if its not
// or token structure if it is valid
func ParseToken(tokenString string) (bool, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})

	// TODO: handle error separately from token checks
	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, fmt.Errorf("invalid token")
	}

	return true, nil
}

func GetUsernameFromClaims(token *jwt.Token) (string, error) {
	return token.Claims.(jwt.MapClaims)["username"].(string), nil
}
