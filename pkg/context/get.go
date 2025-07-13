package context

import (
	"context"
	"e-procurement/pkg/constans"
	"errors"
	"log"
)

func GetUserIDFromContext(ctx context.Context) (string, error) {
	if ctx == nil {
		return "", errors.New("context is nil")
	}

	userID, ok := ctx.Value(constans.ContextUserIDKey).(string)
	log.Printf("Extracted user ID from context: %s", userID)
	if !ok || userID == "" {
		return "", errors.New("user ID not found in context")
	}

	return userID, nil
}