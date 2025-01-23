package utils

import "fmt"

func GetUserKey(key string) string {
	return fmt.Sprintf("u:%s:otp", key)
}
