package auth

import (
	"context"
)

func GetUsernameFromContext(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(UsernameKey).(string)
	return userID, ok
}
