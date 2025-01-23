package response

const (
	ErrCodeSuccess      = 20001 // Success,
	ErrCodeParamInvalid = 20003 // Email is invalid

	ErrInvalidToken = 30001 // Invalid token
	ErrInvalidOTP   = 30002 // Invalid otp
	ErrSendEmailOTP = 30003 // Send email failed

	// Register Code
	ErrCodeUserHasExist = 50001 // User has exist

	// Login Code
	ErrCodeOTPNotExist = 60001
)

// message
var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "email is invalid",

	ErrInvalidToken: "token is invalid",
	ErrInvalidOTP:   "otp is invalid",
	ErrSendEmailOTP: "send email otp failed",

	ErrCodeUserHasExist: "user has exist",

	ErrCodeOTPNotExist: "otp exists but not registered",
}
