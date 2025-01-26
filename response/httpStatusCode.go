package response

const (
	ErrCodeSuccess      = 20001 // Success,
	ErrCodeParamInvalid = 20003 // Email is invalid

	ErrInvalidToken = 30001 // Invalid token
	ErrInvalidOTP   = 30002 // Invalid otp
	ErrSendEmailOTP = 30003 // Send email failed

	ErrCodeAuthFailed = 40001 // Auth failed

	// Register Code
	ErrCodeUserHasExist = 50001 // User has exist
	ErrCodeBindRegisterInput = 50002
	ErrCodeBindVerifyInput = 50003
	ErrCodeVerifyOTPFail = 50004
	ErrCodeBindUpdatePasswordInput = 50005
	ErrCodeBindLoginInput = 50006

	// Login Code
	ErrCodeOTPNotExist = 60001

	ErrCodeUserOTPNotExist = 60002
	ErrCodeOTPDontVerify = 60003

	ErrCodeUpdatePasswordRegister = 100000
	
	// crypto code
	ErrCodeCryptoHash = 70001
	ErrCodeGeneratorSalt = 70002

	// database code
	ErrCodeAddUserBase = 80001
    ErrCodeQueryUserBase = 80002
    ErrCodeUpdateUserBase = 80003
    ErrCodeDeleteUserBase = 80004
    ErrCodeUserBaseNotFound = 80005
	
	ErrCodeAddUserInfo = 90001
	ErrCodeUserNotFound = 90002

	// two factor authentication code
	ErrCodeTwoFactorAuthSetupFailed = 9002
	ErrCodeTwoFactorAuthFailed = 9003
)

// message
var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "email is invalid",
	ErrInvalidToken: "token is invalid",
	ErrInvalidOTP:   "otp is invalid",
	ErrSendEmailOTP: "send email otp failed",
	ErrCodeUserHasExist: "user has exist",
	ErrCodeBindRegisterInput: "bind register input failed",
	ErrCodeBindVerifyInput: "bind verify input failed",
	ErrCodeVerifyOTPFail: "verify otp failed",
	ErrCodeBindUpdatePasswordInput: "bind update password input failed",
	ErrCodeOTPNotExist: "otp exists but not registered",
	ErrCodeUserOTPNotExist: "user otp does not exist",
	ErrCodeOTPDontVerify: "otp does not verify",
	ErrCodeCryptoHash: "crypto hash failed",
	ErrCodeGeneratorSalt: "generator salt failed",
	ErrCodeAddUserBase: "add user base failed",
	ErrCodeQueryUserBase: "query user base failed",
	ErrCodeUpdateUserBase: "update user base failed",
	ErrCodeDeleteUserBase: "delete user base failed",
	ErrCodeUserBaseNotFound: "user base not found",
	ErrCodeAddUserInfo: "add user info failed",
	ErrCodeUpdatePasswordRegister: "update password register failed",
	ErrCodeUserNotFound: "user not found",
	ErrCodeAuthFailed: "auth failed",
	ErrCodeBindLoginInput: "bind login input failed",
	ErrCodeTwoFactorAuthSetupFailed: "two factor authentication setup failed",
	ErrCodeTwoFactorAuthFailed: "two factor authentication failed",
}
