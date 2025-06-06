package entity

type (
	// input for OTP verification
	InputOtpVerify struct {
		VerifyOtp     string `json:"verify_otp"`
		VerifyKey     string `json:"verify_key"`
		VerifyKeyHash string `json:"verify_key_hash"`
		VerifyType    string `json:"verify_type"`
	}

	// out info otp verification
	OutOtpVerifyInfo struct {
		VerifyId      int64  `json:"verify_id"`
		VerifyOtp     string `json:"verify_otp"`
		VerifyKey     string `json:"verify_key"`
		VerifyKeyHash string `json:"verify_key_hash"`
		VerifyType    string `json:"verify_type"`
		IsVerified    bool   `json:"is_verified"`
		IsDeleted     bool   `json:"is_deleted"`
	}

	// out info base user
	OutInfoBaseUser struct {
		UserId       int64  `json:"user_id"`
		UserAccount  string `json:"user_account"`
		UserPassword string `json:"user_password"`
		UserSalt     string `json:"user_salt"`
	}

	// out info user for admin
	OutInforUserForAdmin struct {
		UserId         int64  `json:"user_id"`
		UserAccount    string `json:"user_account"`
		UserPassword   string `json:"user_password"`
		UserSalt       string `json:"user_salt"`
		UserLoginTime  string `json:"user_login_time"`
		UserLogoutTime string `json:"user_logout_time"`
		UserLoginIp    string `json:"user_login_ip"`
		UserCreatedAt  string `json:"user_created_at"`
		UserUpdatedAt  string `json:"user_updated_at"`
	}

	// input add user base
	InputAddUserBase struct {
		UserAccount  string `json:"user_account"`
		UserPassword string `json:"user_password"`
		UserSalt     string `json:"user_salt"`
	}

	// input update password user base
	InputUpdatePasswordUserBase struct {
		UserId          int64  `json:"user_id"`
		UserPasswordNew string `json:"user_password_new"`
	}

	// input login user base
	InputLoginUserBase struct {
		UserAccount  string `json:"user_account"`
		UserPassword string `json:"user_password"`
	}

	// out get user info
	OutGetUserInfo struct {
		UserId               int64  `json:"user_id"`
		UserAccount          string `json:"user_account"`
		UserNickname         string `json:"user_nickname"`
		UserAvatar           string `json:"user_avatar"`
		UserState            string `json:"user_state"`
		UserMobile           string `json:"user_mobile"`
		UserGender           string `json:"user_gender"`
		UserBirthday         string `json:"user_birthday"`
		UserEmail            string `json:"user_email"`
		UserIsAuthentication bool   `json:"user_is_authentication"`
		CreatedAt            string `json:"created_at"`
		UpdatedAt            string `json:"updated_at"`
	}

	// input AddUserAutoUserId
	InputAddUserAutoUserId struct {
		UserAccount          string `json:"user_account"`
		UserNickname         string `json:"user_nickname"`
		UserAvatar           string `json:"user_avatar"`
		UserState            string `json:"user_state"`
		UserMobile           string `json:"user_mobile"`
		UserGender           string `json:"user_gender"`
		UserBirthday         string `json:"user_birthday"`
		UserEmail            string `json:"user_email"`
		UserIsAuthentication bool   `json:"user_is_authentication"`
	}

	// input AddUserHaveUserId
	InputAddUserHaveUserId struct {
		UserId               int64  `json:"user_id"`
		UserAccount          string `json:"user_account"`
		UserNickname         string `json:"user_nickname"`
		UserAvatar           string `json:"user_avatar"`
		UserState            string `json:"user_state"`
		UserMobile           string `json:"user_mobile"`
		UserGender           string `json:"user_gender"`
		UserBirthday         string `json:"user_birthday"`
		UserEmail            string `json:"user_email"`
		UserIsAuthentication bool   `json:"user_is_authentication"`
	}

	// input EditUserByUserId
	InputEditUserByUserId struct {
		UserId       int64  `json:"user_id"`
		UserNickname string `json:"user_nickname"`
		UserAvatar   string `json:"user_avatar"`
		UserMobile   string `json:"user_mobile"`
		UserGender   string `json:"user_gender"`
		UserBirthday string `json:"user_birthday"`
		UserEmail    string `json:"user_email"`
	}

	// input enable two factor type email
	InputEnableTwoFactorTypeEmail struct {
		UserId            int64  `json:"user_id"`
		TwoFactorAuthType string `json:"two_factor_auth_type"`
		TwoFactorEmail    string `json:"two_factor_email"`
	}

	// input DisableTwoFactor
	InputDisableTwoFactor struct {
		UserId            int64  `json:"user_id"`
		TwoFactorAuthType string `json:"two_factor_auth_type"`
	}

	// input UpdateTwoFactorStatus
	InputUpdateTwoFactorStatus struct {
		UserId            int64  `json:"user_id"`
		TwoFactorAuthType string `json:"two_factor_auth_type"`
	}

	// input check verify two factor type
	InputCheckVerifyTwoFactorType struct {
		UserId            int64  `json:"user_id"`
		TwoFactorAuthType string `json:"two_factor_auth_type"`
	}

	// input GetTwoFactorStatus
	InputGetTwoFactorStatus struct {
		UserId            int64  `json:"user_id"`
		TwoFactorAuthType string `json:"two_factor_auth_type"`
	}

	// input add or update email two factor
	InputAddOrUpdateEmailTwoFactor struct {
		UserId         int64  `json:"user_id"`
		TwoFactorEmail string `json:"two_factor_email"`
	}

	// out info user two factor
	OutInfoUserTwoFactor struct {
		UserId              int64  `json:"user_id"`
		TwoFactorAuthType   string `json:"two_factor_auth_type"`
		TwoFactorAuthSecret string `json:"two_factor_auth_secret"`
		TwoFactorEmail      string `json:"two_factor_email"`
		TwoFactorPhone      string `json:"two_factor_phone"`
		TwoFactorIsActive   bool   `json:"two_factor_is_active"`
		TwoFactorCreatedAt  string `json:"two_factor_created_at"`
		TwoFactorUpdatedAt  string `json:"two_factor_updated_at"`
	}

	// input ReactivateTwoFactor
	InputReactivateTwoFactor struct {
		UserId            int64  `json:"user_id"`
		TwoFactorAuthType string `json:"two_factor_auth_type"`
	}

	// input RemoveTwoFactor
	InputRemoveTwoFactor struct {
		UserId            int64  `json:"user_id"`
		TwoFactorAuthType string `json:"two_factor_auth_type"`
	}

	// out get two factor method by id
	OutGetTwoFactorMethod struct {
		UserId              int64  `json:"user_id"`
		TwoFactorAuthType   string `json:"two_factor_auth_type"`
		TwoFactorAuthSecret string `json:"two_factor_auth_secret"`
		TwoFactorEmail      string `json:"two_factor_email"`
		TwoFactorPhone      string `json:"two_factor_phone"`
		TwoFactorIsActive   bool   `json:"two_factor_is_active"`
		TwoFactorCreatedAt  string `json:"two_factor_created_at"`
		TwoFactorUpdatedAt  string `json:"two_factor_updated_at"`
	}
)
