package context

import (
	"context"
	"errors"

	"github.com/Youknow2509/go-ecommerce/internal/utils/cache"
)

type InfoUserUUID struct {
	UserId      uint64
	UserAccount string
}

// helper get uuid from context
func getSubjectUUID(ctx context.Context) (string, error) {
	sUUID, ok := ctx.Value("subjectUUID").(string)
	if !ok {
		return "", errors.New("uuid not found in context")
	}
	return sUUID, nil
}

// Get userId from uuid with context
func GetUserIdFromUUID(ctx context.Context) (uint64, error) {
	sUUID, err := getSubjectUUID(ctx)
	if err != nil {
		return 0, err
	}
	// get infoUser Redis from uuid
    var userInfo InfoUserUUID
	if err := cache.GetCache(ctx, sUUID, &userInfo); err != nil {
		return 0, err
	}

	return userInfo.UserId, nil
}
