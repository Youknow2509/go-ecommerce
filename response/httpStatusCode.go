package response

const (
	ErrCodeSuccess      = 20001 // Success,
	ErrCodeParamInvalid = 20003 // Email is invalid

	ErrInvalidToken   = 30001 // Invalid token
	ErrInvalidOTP     = 30002 // Invalid otp
	ErrSendEmailOTP   = 30003 // Send email failed
	ErrSchemaValidate = 30004

	ErrCodeAuthFailed = 401 // Auth failed

	// Register Code
	ErrCodeUserHasExist                  = 50001 // User has exist
	ErrCodeCheckUserBaseExists           = 50007 // Check user base exists
	ErrCodeBindRegisterInput             = 50002
	ErrCodeBindVerifyInput               = 50003
	ErrCodeVerifyOTPFail                 = 50004
	ErrCodeBindUpdatePasswordInput       = 50005
	ErrCodeBindLoginInput                = 50006
	ErrCodeCheckUserRegisterCache        = 50008 // Check user register cache
	ErrCodeSpamUserRegister              = 50009 // Spam user register
	ErrCodeUserRegisterCacheNotFound     = 50010 // User register cache not found
	ErrCodeUserRegisterTokenNotFound     = 50011 // User register token not found
	ErrCodeVerifyTokenCreatePasswordFail = 50012 // Verify token create password failed
	ErrCodeCreateUserBase                = 50013 // Create user base failed
	ErrCodeCreateUserInfo                = 50014 // Create user info failed

	// Login Code
	ErrCodeOTPNotExist = 60001

	ErrCodeUserOTPNotExist = 60002
	ErrCodeOTPDontVerify   = 60003

	ErrCodeUpdatePasswordRegister = 100000

	// crypto code
	ErrCodeCryptoHash    = 70001
	ErrCodeGeneratorSalt = 70002

	// database code
	ErrCodeAddUserBase      = 80001
	ErrCodeQueryUserBase    = 80002
	ErrCodeUpdateUserBase   = 80003
	ErrCodeDeleteUserBase   = 80004
	ErrCodeUserBaseNotFound = 80005

	ErrCodeAddUserInfo  = 90001
	ErrCodeUserNotFound = 90002

	// two factor authentication code
	ErrCodeTwoFactorAuthSetupFailed = 9002
	ErrCodeTwoFactorAuthFailed      = 9003

	// rate limit code
	ErrCodeTooManyRequests = 429

	// product
	ErrCodeCreateProductError = 100001

	// cron job
	ErrCodeCronAddJobFailed     = 200001
	ErrCodeCronStartFailed      = 200002
	ErrCodeCronStopFailed       = 200003
	ErrCodeCronRemoveJobFailed  = 200004
	ErrCodeCronJobNotFound      = 200005
	ErrCodeCronJobAlreadyExists = 200006
)

// message
var msg = map[int]string{
	ErrCodeCreateUserInfo:                "create user info failed",
	ErrCodeCreateUserBase:                "create user base failed",
	ErrCodeUserRegisterTokenNotFound:     "user register token not found",
	ErrCodeVerifyTokenCreatePasswordFail: "verify token create password failed",
	ErrCodeUserRegisterCacheNotFound:     "user register cache not found",
	ErrCodeSpamUserRegister:              "spam user register",
	ErrCodeCheckUserRegisterCache:        "check user register cache",
	ErrCodeCheckUserBaseExists:           "check user base exists",
	ErrCodeCronAddJobFailed:              "add cron job failed",
	ErrCodeCronStartFailed:               "start cron job failed",
	ErrCodeCronStopFailed:                "stop cron job failed",
	ErrCodeCronRemoveJobFailed:           "remove cron job failed",
	ErrCodeCronJobNotFound:               "cron job not found",
	ErrCodeCronJobAlreadyExists:          "cron job already exists",
	ErrCodeCreateProductError:            "create product error",
	ErrCodeSuccess:                       "success",
	ErrCodeParamInvalid:                  "email is invalid",
	ErrInvalidToken:                      "token is invalid",
	ErrInvalidOTP:                        "otp is invalid",
	ErrSendEmailOTP:                      "send email otp failed",
	ErrCodeUserHasExist:                  "user has exist",
	ErrCodeBindRegisterInput:             "bind register input failed",
	ErrCodeBindVerifyInput:               "bind verify input failed",
	ErrCodeVerifyOTPFail:                 "verify otp failed",
	ErrCodeBindUpdatePasswordInput:       "bind update password input failed",
	ErrCodeOTPNotExist:                   "otp exists but not registered",
	ErrCodeUserOTPNotExist:               "user otp does not exist",
	ErrCodeOTPDontVerify:                 "otp does not verify",
	ErrCodeCryptoHash:                    "crypto hash failed",
	ErrCodeGeneratorSalt:                 "generator salt failed",
	ErrCodeAddUserBase:                   "add user base failed",
	ErrCodeQueryUserBase:                 "query user base failed",
	ErrCodeUpdateUserBase:                "update user base failed",
	ErrCodeDeleteUserBase:                "delete user base failed",
	ErrCodeUserBaseNotFound:              "user base not found",
	ErrCodeAddUserInfo:                   "add user info failed",
	ErrCodeUpdatePasswordRegister:        "update password register failed",
	ErrCodeUserNotFound:                  "user not found",
	ErrCodeAuthFailed:                    "auth failed",
	ErrCodeBindLoginInput:                "bind login input failed",
	ErrCodeTwoFactorAuthSetupFailed:      "two factor authentication setup failed",
	ErrCodeTwoFactorAuthFailed:           "two factor authentication failed",
	ErrCodeTooManyRequests:               "too many requests",
}
