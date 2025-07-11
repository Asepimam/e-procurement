package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
    jwtSecret []byte
}

// NewJWT initializes a new JWT instance with the provided secret key
func NewJWT(secret string) *JWT {
    return &JWT{
        jwtSecret: []byte(secret),
    }
}

// GenerateToken generates a JWT token with the given user ID
func (j *JWT) GenerateToken(userID string,position string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "position": position,
        // set expiration time to 3 days
        "exp":     time.Now().Add(time.Hour * 24 * 3).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // Use HS256 for HMAC
    return token.SignedString(j.jwtSecret)
}

// ValidateToken validates a JWT token and returns the claims
func (j *JWT) ValidateToken(tokenString string) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
        if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return j.jwtSecret, nil
    })

    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return nil, errors.New("invalid token claims")
    }

    return claims, nil
}

func (j *JWT) RefreshToken()(string, error) {return "", nil}