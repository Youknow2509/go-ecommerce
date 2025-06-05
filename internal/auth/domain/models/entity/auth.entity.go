package entity

// struct domain auth
type (
	Login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// ###########################

	Token struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	// ###########################

	SetupTwoFactorAuth struct {
		UserId uint32 `json:"user_id"`
		Type   string `json:"type"`            // e.g., "email", "sms", "app"
		Email  string `json:"email,omitempty"` // optional, if type is "email"
		Phone  string `json:"phone,omitempty"` // optional, if type is "sms"
		App    string `json:"app,omitempty"`   // optional, if type is "app"
	}

	// ###########################

	VerifyTwoFactorAuth struct {
		UserId uint32 `json:"user_id"`
		Code   string `json:"code"` // The verification code sent to the user
		Type   string `json:"type"` // e.g., "email", "sms", "app"
	}

	// ###########################

	InputTokenCreate struct {
		UserId                uint32 `json:"user_id"`
		AccessToken           string `json:"access_token"`
		RefreshToken          string `json:"refresh_token"`
		AccessTokenExpiresAt  int64  `json:"access_token_expires_at"`
		RefreshTokenExpiresAt int64  `json:"refresh_token_expires_at"`
	}
)
