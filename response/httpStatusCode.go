package response

const (
	ErrCodeSuccess      = 20001 // Success,
	ErrCodeParamInvalid = 20003 // Email is invalid
	ErrInvalidToken     = 20004 // Invalid token
	// Register Code
	ErrCodeUserHasExist = 50001 // User has exist
)

// message
var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "email is invalid",
	ErrInvalidToken:     "token is invalid",
	ErrCodeUserHasExist: "user has exist", 
}
