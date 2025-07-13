package auth

import (
	"context"
	"e-procurement/pkg/constans"
	response "e-procurement/pkg/responses"
	"net/http"
	"strings"
	"time"
)

type AuthHttp struct {
	jwt    *JWT
}

func NewAuthMiddleware(jwt *JWT) *AuthHttp {
	return &AuthHttp{jwt: jwt}
}


func (m *AuthHttp) VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := extractTokenFromHeader(r)
		if err != nil {
			response.Error(w, http.StatusUnauthorized, err.Error())
			return
		}

		claims, err := m.jwt.ValidateToken(token)
		if err != nil {
			response.Error(w, http.StatusUnauthorized, "Invalid token: "+err.Error())
			return
		}

		if err := validateTokenClaims(claims); err != nil {
			
			response.Error(w, http.StatusUnauthorized, err.Error())
			return
		}

		userID, position, err := extractUserIDAndPosition(claims)
		if err != nil {

			response.Error(w, http.StatusUnauthorized, err.Error())
			return
		}
		
		ctx := context.WithValue(r.Context(), constans.ContextUserIDKey, userID)
		ctx = context.WithValue(ctx, constans.ContextPositionKey, position)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func extractTokenFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errorString("Authorization header is missing")
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == authHeader {
		return "", errorString("Invalid token format: missing 'Bearer' prefix")
	}
	return token, nil
}

func validateTokenClaims(claims map[string]interface{}) error {
	exp, ok := claims["exp"].(float64)
	if !ok {
		return errorString("Token missing 'exp' claim")
	}
	if time.Now().Unix() > int64(exp) {
		return errorString("Token has expired")
	}
	return nil
}

func extractUserIDAndPosition(claims map[string]interface{}) (string, string, error) {
	userIDRaw, ok := claims["user_id"]
	if !ok {
		return "", "", errorString("Token missing 'user_id' claim")
	}
	userID, ok := userIDRaw.(string)
	if !ok {
		return "", "", errorString("Invalid 'user_id' type in token")
	}

	positionRaw, ok := claims["position"]
	if !ok {
		return "", "", errorString("Token missing 'position' claim")
	}
	position, ok := positionRaw.(string)
	if !ok {
		return "", "", errorString("Invalid 'position' type in token")
	}

	return userID, position, nil
}

type errorString string

func (e errorString) Error() string {
	return string(e)
}
