package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func GetUserKey(key string) string {
	return fmt.Sprintf("u:%s:otp", key)
}

// create uuid
func GenerateCliTokenUUID(userId int) string {
	newUUID := uuid.New()
	// convert uuid to string, remove -
	uuidStr := strings.ReplaceAll(newUUID.String(), "-", "")
	return strconv.Itoa(userId) + "clitoken" + uuidStr
}