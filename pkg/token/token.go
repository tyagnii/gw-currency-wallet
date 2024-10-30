package token

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/tyagnii/gw-currency-wallet/internal/models"
	"time"
)

var Secret string
var ExpireTime time.Duration

// Init package variables
func init() {
	// todo: read secret from configuration
	Secret = "very secret string"

	// todo: read expiration time from configuration
	ExpireTime = time.Hour * 24
}

// NewToken creates new JWT token for given user
func NewToken(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(ExpireTime).Unix()

	tokenString, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
