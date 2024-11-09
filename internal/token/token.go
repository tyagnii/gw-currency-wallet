// TODO: MOve the package to the internal directory
package token

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/tyagnii/gw-currency-wallet/internal/db/models"
	"os"
	"strconv"
	"time"
)

// todo: generate secrerte string on the fly
//
//	create token structure
//	make func token methods
var Secret string
var ExpireTime time.Duration

type Claims struct {
	jwt.RegisteredClaims
	Username string
}

// LoadEnvironment loads environment variables
func LoadEnvironment() {
	Secret = os.Getenv("JWT_SECRET")

	v, err := strconv.Atoi(os.Getenv("JWT_EXPIRE_TIME"))
	if err != nil {
		panic(err)
	}
	ExpireTime = time.Duration(v) * time.Hour
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
func ParseToken(tokenString string) (*jwt.Token, Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, *claims, err
	}

	return token, *claims, nil
}

func ValidateToken(token *jwt.Token) bool {
	return token.Valid
}
