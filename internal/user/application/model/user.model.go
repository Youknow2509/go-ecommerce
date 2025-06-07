package model

// input
type (
	// UserRegister is the data transfer object for user registration
	InputUserRegister struct {
		AcountName string `json:"account_name" binding:"required"`
		VerifyType int    `json:"verify_type" binding:"required"` // 1: email, 2: phone
	}

	// UserVerifyRegister is the data transfer object for user registration verification
	InputUserVerifyRegister struct {
		AcountName string `json:"account_name" binding:"required"`
		VerifyCode string `json:"verify_code" binding:"required"`
	}

	// UserCreatePassword is the data transfer object for creating a user password
	InputUserCreatePassword struct {
		AccountName         string `json:"account_name" binding:"required"`
		TokenCreatePassword string `json:"token_create_password" binding:"required"` // Token for creating password
		Password            string `json:"password" binding:"required,min=6,max=20"` // Password must be between 6 and 20 characters
	}
)

// output
type (
	// output after user verify
	OutputUserVerifyRegister struct {
		Token string `json:"token"`
	}
)
