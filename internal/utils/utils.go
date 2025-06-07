package utils

import (
	"fmt"

	"github.com/Youknow2509/go-ecommerce/internal/utils/crypto"
	"github.com/Youknow2509/go-ecommerce/internal/utils/random"
)

// get key user register cache
func GetKeyUserRegisterCache(accountNameHash string) string {
	return fmt.Sprintf("user:register:%s", accountNameHash)
}

// create token for user register
func GetTokenUserRegister(data string) string {
	return crypto.GetHash(fmt.Sprintf("%s:%d", data, random.GenerateSixDigitOtp()))
}

// get key token for create password register
func GetKeyUserRegisterToken(accountNameHash string) string {
	return fmt.Sprintf("user:register:token:%s", accountNameHash)
}

