package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

/**
 * Get key OTP Two Factor in cache 
 */
func GetTwoFactorKey(key string) string {
	return fmt.Sprintf("u:%s:2fa", key)
}

/**
 * Get key OTP in cache
 */
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