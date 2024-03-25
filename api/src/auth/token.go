package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthToken struct {
	ExpiresAt time.Time
	IssuedAt  time.Time
	UserID    string
}

func GenerateToken(userID string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"iat": jwt.NewNumericDate(now),
		"exp": jwt.NewNumericDate(now.Add(time.Minute * 10)),
	})
	return token.SignedString([]byte(secret))
}

func DecodeToken(tokenString string) (AuthToken, error) {
	secret := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return AuthToken{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return AuthToken{}, fmt.Errorf("something is wrong")
	}

	return createAuthTokenFromClaims(claims)
}

func (authToken AuthToken) IsValid() bool {
	now := time.Now()
	return now.After(authToken.IssuedAt) && now.Before(authToken.ExpiresAt)
}

func createAuthTokenFromClaims(claims jwt.MapClaims) (AuthToken, error) {
	expiresAt, err := claims.GetExpirationTime()
	if err != nil {
		return AuthToken{}, err
	}

	issuedAt, err := claims.GetIssuedAt()
	if err != nil {
		return AuthToken{}, err
	}

	userID, err := claims.GetSubject()
	if err != nil {
		return AuthToken{}, err
	}

	return AuthToken{
		ExpiresAt: expiresAt.Time,
		IssuedAt:  issuedAt.Time,
		UserID:    userID,
	}, nil
}
