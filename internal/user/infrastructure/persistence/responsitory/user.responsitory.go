package responsitory

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Youknow2509/go-ecommerce/internal/user/domain/model/entity"
	"github.com/Youknow2509/go-ecommerce/internal/user/domain/responsitory"
	"github.com/Youknow2509/go-ecommerce/internal/user/infrastructure/database"
)

// struct userResponsitory implements IUserResponsitory interface
type UserResponsitory struct {
	db *database.Queries
}

// #########################################################

// AddOrUpdateEmailTwoFactor implements responsitory.IUserResponsitory.
func (u *UserResponsitory) AddOrUpdateEmailTwoFactor(ctx context.Context, input *entity.InputAddOrUpdateEmailTwoFactor) error {
	return u.db.AddOrUpdateEmail(
		ctx,
		database.AddOrUpdateEmailParams{
			UserID:           uint32(input.UserId),
			TwoFactorEmail:   sql.NullString{String: input.TwoFactorEmail, Valid: true},
			TwoFactorEmail_2: sql.NullString{String: input.TwoFactorEmail, Valid: true},
		},
	)
}

// AddUserBase implements responsitory.IUserResponsitory.
func (u *UserResponsitory) AddUserBase(ctx context.Context, input *entity.InputAddUserBase) error {
	_, err := u.db.AddUserBase(
		ctx,
		database.AddUserBaseParams{
			UserAccount:  input.UserAccount,
			UserPassword: input.UserPassword,
			UserSalt:     input.UserSalt,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// AddUserInfoAutoId implements responsitory.IUserResponsitory.
func (u *UserResponsitory) AddUserInfoAutoId(ctx context.Context, input *entity.InputAddUserAutoUserId) error {
	_, err := u.db.AddUserAutoUserId(
		ctx,
		database.AddUserAutoUserIdParams{
			UserAccount:          input.UserAccount,
			UserNickname:         sql.NullString{String: input.UserNickname, Valid: true},
			UserAvatar:           sql.NullString{String: input.UserAvatar, Valid: true},
			UserState:            uint8(input.UserState),
			UserMobile:           sql.NullString{String: input.UserMobile, Valid: true},
			UserGender:           sql.NullInt16{Int16: int16(input.UserGender), Valid: true},
			UserBirthday:         sql.NullTime{Time: input.UserBirthday, Valid: true},
			UserEmail:            sql.NullString{String: input.UserEmail, Valid: true},
			UserIsAuthentication: uint8(input.UserIsAuthentication),
		},
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil // No rows affected, which is acceptable
		}
		return err
	}
	return nil
}

// AddUserInfoHaveUserId implements responsitory.IUserResponsitory.
func (u *UserResponsitory) AddUserInfoHaveUserId(ctx context.Context, input *entity.InputAddUserHaveUserId) error {
	_, err := u.db.AddUserHaveUserId(
		ctx,
		database.AddUserHaveUserIdParams{
			UserID:               uint64(input.UserId),
			UserNickname:         sql.NullString{String: input.UserNickname, Valid: true},
			UserAvatar:           sql.NullString{String: input.UserAvatar, Valid: true},
			UserState:            uint8(input.UserState),
			UserMobile:           sql.NullString{String: input.UserMobile, Valid: true},
			UserGender:           sql.NullInt16{Int16: int16(input.UserGender), Valid: true},
			UserBirthday:         sql.NullTime{Time: input.UserBirthday, Valid: true},
			UserEmail:            sql.NullString{String: input.UserEmail, Valid: true},
			UserIsAuthentication: uint8(input.UserIsAuthentication),
		},
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil // No rows affected, which is acceptable
		}
		return err
	}
	return nil
}

// CheckUserBaseExists implements responsitory.IUserResponsitory.
func (u *UserResponsitory) CheckUserBaseExists(ctx context.Context, user_account string) (bool, error) {
	_, err := u.db.CheckUserBaseExists(ctx, user_account)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CheckVerifyTwoFactorType implements responsitory.IUserResponsitory.
func (u *UserResponsitory) CheckVerifyTwoFactorType(ctx context.Context, input *entity.InputCheckVerifyTwoFactorType) error {
	response, err := u.db.VerifyTwoFactor(
		ctx,
		database.VerifyTwoFactorParams{
			UserID:            uint32(input.UserId),
			TwoFactorAuthType: database.UserTwoFactor001TwoFactorAuthType(input.TwoFactorAuthType),
		},
	)
	if err != nil {
		return err
	}
	if response == 0 {
		return fmt.Errorf("two factor type %s not found for user id %d", input.TwoFactorAuthType, input.UserId)
	}
	return nil
}

// CountActiveTwoFactorMethods implements responsitory.IUserResponsitory.
func (u *UserResponsitory) CountActiveTwoFactorMethods(ctx context.Context, user_id int64) (int64, error) {
	return u.db.CountActiveTwoFactorMethods(
		ctx,
		uint32(user_id),
	)
}

// DisableTwoFactor implements responsitory.IUserResponsitory.
func (u *UserResponsitory) DisableTwoFactor(ctx context.Context, input *entity.InputDisableTwoFactor) error {
	return u.db.DisableTwoFactor(
		ctx,
		database.DisableTwoFactorParams{
			UserID:            uint32(input.UserId),
			TwoFactorAuthType: database.UserTwoFactor001TwoFactorAuthType(input.TwoFactorAuthType),
		},
	)
}

// EditUserByUserId implements responsitory.IUserResponsitory.
func (u *UserResponsitory) EditUserByUserId(ctx context.Context, input *entity.InputEditUserByUserId) error {
	_, err := u.db.EditUserByUserId(
		ctx,
		database.EditUserByUserIdParams{
			UserID:       uint64(input.UserId),
			UserNickname: sql.NullString{String: input.UserNickname, Valid: true},
			UserAvatar:   sql.NullString{String: input.UserAvatar, Valid: true},
			UserMobile:   sql.NullString{String: input.UserMobile, Valid: true},
			UserGender:   sql.NullInt16{Int16: int16(input.UserGender), Valid: true},
			UserBirthday: sql.NullTime{Time: input.UserBirthday, Valid: true},
			UserEmail:    sql.NullString{String: input.UserEmail, Valid: true},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// EnableTwoFactorTypeEmail implements responsitory.IUserResponsitory.
func (u *UserResponsitory) EnableTwoFactorTypeEmail(ctx context.Context, input *entity.InputEnableTwoFactorTypeEmail) error {
	return u.db.EnableTwoFactorTypeEmail(
		ctx,
		database.EnableTwoFactorTypeEmailParams{
			UserID:            uint32(input.UserId),
			TwoFactorAuthType: database.UserTwoFactor001TwoFactorAuthType(input.TwoFactorAuthType),
			TwoFactorEmail:    sql.NullString{String: input.TwoFactorEmail, Valid: true},
		},
	)
}

// GetInfoOTPVerify implements responsitory.IUserResponsitory.
func (u *UserResponsitory) GetInfoOTPVerify(ctx context.Context, key_hash string) (*entity.OutOtpVerifyInfo, error) {
	respon, err := u.db.GetInfoOTP(
		ctx,
		key_hash,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no OTP verify info found for key hash: %s", key_hash)
		}
		return nil, err
	}
	return &entity.OutOtpVerifyInfo{
		VerifyId:      int64(respon.VerifyID),
		VerifyOtp:     respon.VerifyOtp,
		VerifyKey:     respon.VerifyKey,
		VerifyKeyHash: respon.VerifyKeyHash,
		VerifyType:    int(respon.VerifyType.Int32),
		IsVerified:    respon.IsVerified.Int32 == 1,
		IsDeleted:     respon.IsDeleted.Int32 == 1,
	}, nil
}

// GetOneUserInfo implements responsitory.IUserResponsitory.
func (u *UserResponsitory) GetOneUserInfo(ctx context.Context, user_account string) (*entity.OutInfoBaseUser, error) {
	respon, err := u.db.GetOneUserInfo(
		ctx,
		user_account,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user info found for account: %s", user_account)
		}
		return nil, err
	}
	return &entity.OutInfoBaseUser{
		UserId:       int64(respon.UserID),
		UserAccount:  respon.UserAccount,
		UserPassword: respon.UserPassword,
		UserSalt:     respon.UserSalt,
	}, nil
}

// GetOneUserInfoForAdmin implements responsitory.IUserResponsitory.
func (u *UserResponsitory) GetOneUserInfoForAdmin(ctx context.Context, user_account string) (*entity.OutInforUserForAdmin, error) {
	response, err := u.db.GetOneUserInfoAdmin(
		ctx,
		user_account,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user info found for account: %s", user_account)
		}
		return nil, err
	}
	return &entity.OutInforUserForAdmin{
		UserId:         int64(response.UserID),
		UserAccount:    response.UserAccount,
		UserPassword:   response.UserPassword,
		UserSalt:       response.UserSalt,
		UserLoginTime:  response.UserLoginTime.Time.Format("2006-01-02 15:04:05"),
		UserLogoutTime: response.UserLogoutTime.Time.Format("2006-01-02 15:04:05"),
		UserLoginIp:    response.UserLoginIp.String,
		UserCreatedAt:  response.UserCreatedAt.Time.Format("2006-01-02 15:04:05"),
		UserUpdatedAt:  response.UserUpdatedAt.Time.Format("2006-01-02 15:04:05"),
	}, nil
}

// GetTwoFactorMethodByID implements responsitory.IUserResponsitory.
func (u *UserResponsitory) GetTwoFactorMethodByID(ctx context.Context, user_id int64) (*entity.OutGetTwoFactorMethod, error) {
	response, err := u.db.GetTwoFactorMethodByID(
		ctx,
		uint32(user_id),
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no two factor method found for user id: %d", user_id)
		}
		return nil, err
	}
	return &entity.OutGetTwoFactorMethod{
		UserId:              int64(response.UserID),
		TwoFactorAuthType:   string(response.TwoFactorAuthType),
		TwoFactorAuthSecret: response.TwoFactorAuthSecret,
		TwoFactorEmail:      response.TwoFactorEmail.String,
		TwoFactorPhone:      response.TwoFactorPhone.String,
		TwoFactorIsActive:   response.TwoFactorIsActive,
		TwoFactorCreatedAt:  response.TwoFactorCreatedAt.Time.Format("2006-01-02 15:04:05"),
		TwoFactorUpdatedAt:  response.TwoFactorUpdatedAt.Time.Format("2006-01-02 15:04:05"),
	}, nil

}

// GetTwoFactorMethodByIDAndByType implements responsitory.IUserResponsitory.
func (u *UserResponsitory) GetTwoFactorMethodByIDAndByType(ctx context.Context, user_id int64, type_verify string) (*entity.OutGetTwoFactorMethod, error) {
	response, err := u.db.GetTwoFactorMethodByIDAndByType(
		ctx,
		database.GetTwoFactorMethodByIDAndByTypeParams{
			UserID:            uint32(user_id),
			TwoFactorAuthType: database.UserTwoFactor001TwoFactorAuthType(type_verify),
		},
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no two factor method found for user id: %d and type: %s", user_id, type_verify)
		}
		return nil, err
	}
	return &entity.OutGetTwoFactorMethod{
		UserId:              int64(response.UserID),
		TwoFactorAuthType:   string(response.TwoFactorAuthType),
		TwoFactorAuthSecret: response.TwoFactorAuthSecret,
		TwoFactorEmail:      response.TwoFactorEmail.String,
		TwoFactorPhone:      response.TwoFactorPhone.String,
		TwoFactorIsActive:   response.TwoFactorIsActive,
		TwoFactorCreatedAt:  response.TwoFactorCreatedAt.Time.Format("2006-01-02 15:04:05"),
		TwoFactorUpdatedAt:  response.TwoFactorUpdatedAt.Time.Format("2006-01-02 15:04:05"),
	}, nil
}

// GetTwoFactorStatus implements responsitory.IUserResponsitory.
func (u *UserResponsitory) GetTwoFactorStatus(ctx context.Context, input *entity.InputGetTwoFactorStatus) (bool, error) {
	return u.db.GetTwoFactorStatus(ctx, database.GetTwoFactorStatusParams{
		UserID:            uint32(input.UserId),
		TwoFactorAuthType: database.UserTwoFactor001TwoFactorAuthType(input.TwoFactorAuthType),
	})
}

// GetUserTwoFactoryMethods implements responsitory.IUserResponsitory.
func (u *UserResponsitory) GetUserTwoFactoryMethods(ctx context.Context, user_id int64) ([]*entity.OutInfoUserTwoFactor, error) {
	rows, err := u.db.GetUserTwoFactoryMethods(ctx, uint32(user_id))
	if err != nil {
		return nil, err
	}
	var result []*entity.OutInfoUserTwoFactor
	for _, r := range rows {
		result = append(result, &entity.OutInfoUserTwoFactor{
			UserId:              int64(r.UserID),
			TwoFactorAuthType:   string(r.TwoFactorAuthType),
			TwoFactorAuthSecret: r.TwoFactorAuthSecret,
			TwoFactorEmail:      r.TwoFactorEmail.String,
			TwoFactorPhone:      r.TwoFactorPhone.String,
			TwoFactorIsActive:   r.TwoFactorIsActive,
			TwoFactorCreatedAt:  r.TwoFactorCreatedAt.Time.Format("2006-01-02 15:04:05"),
			TwoFactorUpdatedAt:  r.TwoFactorUpdatedAt.Time.Format("2006-01-02 15:04:05"),
		})
	}
	return result, nil
}

// GetValidOtp implements responsitory.IUserResponsitory.
func (u *UserResponsitory) GetValidOtp(ctx context.Context, key_hash string) error {
	_, err := u.db.GetValidOtp(ctx, key_hash)
	return err
}

// InsertOTPVerify implements responsitory.IUserResponsitory.
func (u *UserResponsitory) InsertOTPVerify(ctx context.Context, input *entity.InputOtpVerify) error {
	_, err := u.db.InsertOTPVerify(ctx, database.InsertOTPVerifyParams{
		VerifyOtp:     input.VerifyOtp,
		VerifyKey:     input.VerifyKey,
		VerifyKeyHash: input.VerifyKeyHash,
		VerifyType:    sql.NullInt32{Int32: int32(input.VerifyType), Valid: true},
	})
	return err
}

// IsTwoFactorEnabled implements responsitory.IUserResponsitory.
func (u *UserResponsitory) IsTwoFactorEnabled(ctx context.Context, user_id int64) (bool, error) {
	count, err := u.db.IsTwoFactorEnabled(ctx, uint32(user_id))
	return count > 0, err
}

// LoginUserBase implements responsitory.IUserResponsitory.
func (u *UserResponsitory) LoginUserBase(ctx context.Context, input *entity.InputLoginUserBase) error {
	return u.db.LoginUserBase(ctx, database.LoginUserBaseParams{
		UserLoginIp:  sql.NullString{String: input.UserIpLogin, Valid: true},
		UserAccount:  input.UserAccount,
		UserPassword: input.UserPassword,
	})
}

// LogoutUserBase implements responsitory.IUserResponsitory.
func (u *UserResponsitory) LogoutUserBase(ctx context.Context, user_account string) error {
	return u.db.LogoutUserBase(ctx, user_account)
}

// ReactivateTwoFactor implements responsitory.IUserResponsitory.
func (u *UserResponsitory) ReactivateTwoFactor(ctx context.Context, input *entity.InputReactivateTwoFactor) error {
	return u.db.ReactivateTwoFactor(ctx, database.ReactivateTwoFactorParams{
		UserID:            uint32(input.UserId),
		TwoFactorAuthType: database.UserTwoFactor001TwoFactorAuthType(input.TwoFactorAuthType),
	})
}

// RemoveTwoFactor implements responsitory.IUserResponsitory.
func (u *UserResponsitory) RemoveTwoFactor(ctx context.Context, input *entity.InputRemoveTwoFactor) error {
	return u.db.RemoveTwoFactor(ctx, database.RemoveTwoFactorParams{
		UserID:            uint32(input.UserId),
		TwoFactorAuthType: database.UserTwoFactor001TwoFactorAuthType(input.TwoFactorAuthType),
	})
}

// RemoveUserInfo implements responsitory.IUserResponsitory.
func (u *UserResponsitory) RemoveUserInfo(ctx context.Context, user_id int64) error {
	return u.db.RemoveUser(ctx, uint64(user_id))
}

// UpdatePasswordUserBase implements responsitory.IUserResponsitory.
func (u *UserResponsitory) UpdatePasswordUserBase(ctx context.Context, input *entity.InputUpdatePasswordUserBase) error {
	return u.db.UpdatePassword(ctx, database.UpdatePasswordParams{
		UserPassword: input.UserPasswordNew,
		UserID:       int32(input.UserId),
	})
}

// UpdateTwoFactorStatus implements responsitory.IUserResponsitory.
func (u *UserResponsitory) UpdateTwoFactorStatus(ctx context.Context, input *entity.InputUpdateTwoFactorStatus) error {
	return u.db.UpdateTwoFactorStatus(ctx, database.UpdateTwoFactorStatusParams{
		UserID:            uint32(input.UserId),
		TwoFactorAuthType: database.UserTwoFactor001TwoFactorAuthType(input.TwoFactorAuthType),
	})
}

// UpdateUserVerificationStatus implements responsitory.IUserResponsitory.
func (u *UserResponsitory) UpdateUserVerificationStatus(ctx context.Context, key_hash string) error {
	return u.db.UpdateUserVerificationStatus(ctx, key_hash)
}

// #########################################################

// new instant and impl IUserResponsitory interface
func NewUserResponsitory(db *database.Queries) responsitory.IUserResponsitory {
	return &UserResponsitory{
		db: db,
	}
}
