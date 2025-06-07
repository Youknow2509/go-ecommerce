package dto

type (
	// UserRegisterDto is the data transfer object for user registration
	UserRegisterDto struct {
		AcountName string `json:"account_name" binding:"required"`
		VerifyType int    `json:"verify_type" binding:"required"` // 1: email, 2: phone
	}

	// UserVerifyRegisterDto is the data transfer object for user registration verification
	UserVerifyRegisterDto struct {
		AcountName string `json:"account_name" binding:"required"`
		VerifyCode string `json:"verify_code" binding:"required"`
	}

	// UserCreatePasswordDto is the data transfer object for creating a user password
	UserCreatePasswordDto struct {
		AccountName         string `json:"account_name" binding:"required"`          // Account name must be unique
		TokenCreatePassword string `json:"token_create_password" binding:"required"` // Token for creating password
		Password            string `json:"password" binding:"required,min=6,max=20"` // Password must be between 6 and 20 characters
	}
)
