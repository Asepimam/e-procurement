package context

import (
	"context"
	"e-procurement/pkg/constans"
	"errors"
)

func GetUserIDFromContext(ctx context.Context) (string, error) {
	if ctx == nil {
		return "", errors.New("context is nil")
	}

	userID, ok := ctx.Value(constans.ContextUserIDKey).(string)
	if !ok || userID == "" {
		return "", errors.New("user ID not found in context")
	}

	return userID, nil
}