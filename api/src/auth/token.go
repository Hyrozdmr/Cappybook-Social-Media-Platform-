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

var secret = os.Getenv("JWT_SECRET")

func GenerateToken(userID string) (string, error) {
	now := time.Now()
	/*
	 * Creates a token variable that creates a new webtoken, signs it, and maps its three values with a userID,
	 * a time and date which will be now, and an expiration time with now plus 10 minutes
	 */
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"iat": jwt.NewNumericDate(now),
		"exp": jwt.NewNumericDate(now.Add(time.Minute * 10)),
	})
	// This is then returned and signed into a string of bytes,
	//  and into our secret environment variable
	return token.SignedString([]byte(secret))
}

func DecodeToken(tokenString string) (AuthToken, error) {
	/*
	 * First two lines take the code and parse it using the jwt package. Within it,
	 * there is a pointer to the jwt.Token, presumably to pull this function and apply it to the tokenString.
	 */
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// If it is ok, it will move on to the token.Header line, otherwise it will raise an error
		// It will then return the decoded token back into secrets and no error if there is none
		return secret, nil
	})
	// From here the authToken is created proper with the MapClaims function,
	// which is then output as the AuthToken proper or an error if there is one
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok {
		return createAuthTokenFromClaims(claims)
	} else {
		return AuthToken{}, err
	}
}

func (authToken AuthToken) IsValid() bool {
	now := time.Now()
	// Basically checks if the AuthToken is valid by verifying the time it was
	// created is within the parameters set earlier,
	// that the time it was issued is not after now, and before the expiry time
	return now.After(authToken.IssuedAt) && now.Before(authToken.ExpiresAt)
}

func createAuthTokenFromClaims(claims jwt.MapClaims) (AuthToken, error) {
	/*
	 * From the previously issued AuthToken from claims in the decode function,
	 * this function is what powers the ending of the Decode Token and so is only callable from within this package
	 */
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
	// Function checks the expiration time, when it was issued, the user id,
	// using jwt.MapClaims to do so, and returning errors if there are any.
	// If there are no errors, it will return the AuthToken in the format of the
	// AuthToken struct
	return AuthToken{
		ExpiresAt: expiresAt.Time,
		IssuedAt:  issuedAt.Time,
		UserID:    userID,
	}, nil
}
