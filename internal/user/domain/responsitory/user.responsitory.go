package responsitory

import (
	"context"
	"fmt"

	"github.com/Youknow2509/go-ecommerce/internal/user/domain/model/entity"
)

// IUserResponsitory defines the interface for user-related operations in the repository layer.
type (
	IUserResponsitory interface {
		// user verify
		GetValidOtp(ctx context.Context, key_hash string) error
		UpdateUserVerificationStatus(ctx context.Context, key_hash string) error
		InsertOTPVerify(ctx context.Context, input *entity.InputOtpVerify) error
		GetInfoOTPVerify(ctx context.Context, key_hash string) (*entity.OutOtpVerifyInfo, error)

		// user base
		GetOneUserInfo(ctx context.Context, user_account string) (*entity.OutInfoBaseUser, error)
		GetOneUserInfoForAdmin(ctx context.Context, user_id int64) (*entity.OutInforUserForAdmin, error)
		CheckUserBaseExists(ctx context.Context, user_account string) (bool, error)
		AddUserBase(ctx context.Context, input *entity.InputAddUserBase) error
		UpdatePasswordUserBase(ctx context.Context, input *entity.InputUpdatePasswordUserBase) error
		LoginUserBase(ctx context.Context, input *entity.InputLoginUserBase) error
		LogoutUserBase(ctx context.Context, user_account string) error

		// user info
		RemoveUserInfo(ctx context.Context, user_id int64) error
		AddUserInfoAutoId(ctx context.Context, input *entity.InputAddUserAutoUserId) error
		AddUserInfoHaveUserId(ctx context.Context, input *entity.InputAddUserHaveUserId) error
		EditUserByUserId(ctx context.Context, input *entity.InputEditUserByUserId) error

		// user two factor
		EnableTwoFactorTypeEmail(ctx context.Context, input *entity.InputEnableTwoFactorTypeEmail) error
		DisableTwoFactor(ctx context.Context, input *entity.InputDisableTwoFactor) error
		UpdateTwoFactorStatus(ctx context.Context, input *entity.InputUpdateTwoFactorStatus) error
		CheckVerifyTwoFactorType(ctx context.Context, input *entity.InputCheckVerifyTwoFactorType) error
		GetTwoFactorStatus(ctx context.Context, input *entity.InputGetTwoFactorStatus) (bool, error)
		IsTwoFactorEnabled(ctx context.Context, user_id int64) (bool, error)
		AddOrUpdateEmailTwoFactor(ctx context.Context, input *entity.InputAddOrUpdateEmailTwoFactor) error
		GetUserTwoFactoryMethods(ctx context.Context, user_id int64) ([]*entity.OutInfoUserTwoFactor, error)
		ReactivateTwoFactor(ctx context.Context, input *entity.InputReactivateTwoFactor) error
		RemoveTwoFactor(ctx context.Context, input *entity.InputRemoveTwoFactor) error
		CountActiveTwoFactorMethods(ctx context.Context, user_id int64) (int64, error)
		GetTwoFactorMethodByID(ctx context.Context, user_id int64) (*entity.OutGetTwoFactorMethod, error)
		GetTwoFactorMethodByIDAndByType(ctx context.Context, user_id int64) (*entity.OutGetTwoFactorMethod, error)
	}
)
// ########################################################

var (
	vIResponsitory IUserResponsitory
)
// ########################################################

// init interface user responsitory
func InitUserResponsitory(i IUserResponsitory) {
	vIResponsitory = i
}

// GetUserResponsitory returns the current user responsitory interface.
func GetUserResponsitory() (IUserResponsitory, error) {
	if vIResponsitory == nil {
		return nil, fmt.Errorf("responsitory user is not initialized, please call InitUserResponsitory first")
	}
	return vIResponsitory, nil
}