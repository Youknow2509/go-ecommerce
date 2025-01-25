package impl

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/Youknow2509/go-ecommerce/internal/consts"
	"github.com/Youknow2509/go-ecommerce/internal/database"
	"github.com/Youknow2509/go-ecommerce/internal/model"
	"github.com/Youknow2509/go-ecommerce/internal/service"
	"github.com/Youknow2509/go-ecommerce/internal/utils"
	"github.com/Youknow2509/go-ecommerce/internal/utils/auth"
	"github.com/Youknow2509/go-ecommerce/internal/utils/crypto"
	"github.com/Youknow2509/go-ecommerce/internal/utils/random"
	"github.com/Youknow2509/go-ecommerce/internal/utils/sendto"
	"github.com/Youknow2509/go-ecommerce/internal/utils/sendto/create"
	"github.com/Youknow2509/go-ecommerce/response"
	"github.com/redis/go-redis/v9"
)

// struct
type sUserLogin struct {
	r *database.Queries
}

// new sUserLogin implementation interface for IUserLogin
func NewSUserLogin(r *database.Queries) service.IUserLogin {
	return &sUserLogin{
		r: r,
	}
}

// two-factor authentication
func (s *sUserLogin) IsTwoFactorEnabled(ctx context.Context, userId int) (codeResult int, err error) {
	// TODO
	return response.ErrCodeSuccess, nil
}

// setup authentication
func (s *sUserLogin) SetupTwoFactorAuth(ctx context.Context, in *model.SetupTwoFactorAuthInput) (codeResult int, err error) {
	// check if user is already
	isTwoFactorAuth, err := s.r.IsTwoFactorEnabled(ctx, in.UserId)
	if err != nil {
		return response.ErrCodeTwoFactorAuthSetupFailed, err
	}
	if isTwoFactorAuth > 0 {
		return response.ErrCodeTwoFactorAuthSetupFailed, errors.New("two factor authentication is already")
	}
	// create new type auth
	err = s.r.EnableTwoFactorTypeEmail(ctx, database.EnableTwoFactorTypeEmailParams{
		UserID:            in.UserId,
		TwoFactorAuthType: database.PreGoAccUserTwoFactor9999TwoFactorAuthTypeEMAIL,
		TwoFactorEmail:    in.TwoFactorEmail,
	})
	if err != nil {
		return response.ErrCodeTwoFactorAuthSetupFailed, err
	}
	// send otp to email
	keyUserTwoFactor := crypto.GetHash("2fa" + strconv.Itoa(int(in.UserId)))
	otp := random.GenerateSixDigitOtp()
	go global.Rdb.Set(ctx, keyUserTwoFactor, otp, time.Duration(consts.TIME_OTP_REGISTER)*time.Minute).Err()
	
	
	// TODO
	return response.ErrCodeSuccess, nil
}

// verify authentication
func (s *sUserLogin) VerifyTwoFactorAuth(ctx context.Context, in *model.TwoFactorVerificationInput) (codeResult int, err error) {
	// TODO
	return response.ErrCodeSuccess, nil
}

// Login implements service.IUserLogin.
func (s *sUserLogin) Login(ctx context.Context, in *model.LoginInput) (codeResult int, out model.LoginOutput, err error) {
	// check user in table user_base
	userBase, err := s.r.GetOneUserInfo(ctx, in.UserAccount)
	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}
	// check password
	if !crypto.ComparePasswordWithHash(in.UserPassword, userBase.UserSalt, userBase.UserPassword) {
		return response.ErrCodeAuthFailed, out, errors.New("password not match")
	}
	// TODO check two-factor authentication

	// upgrade state login
	go s.r.LoginUserBase(ctx, database.LoginUserBaseParams{
		UserLoginIp:  sql.NullString{String: "127.0.0.1", Valid: true},
		UserAccount:  in.UserAccount,
		UserPassword: userBase.UserPassword,
	})
	// create uuid
	subToken := utils.GenerateCliTokenUUID(int(userBase.UserID))
	log.Println("subToken: ", subToken)
	// get user info table
	infoUser, err := s.r.GetUser(ctx, uint64(userBase.UserID))
	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}
	// convert to json
	infoUserJson, err := json.Marshal(infoUser)
	if err != nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("convert json failed: %w", err)
	}
	// give infoUserJson to redis with key = subToken
	err = global.Rdb.Set(ctx, subToken, infoUserJson, time.Duration(consts.TIME_OTP_REGISTER)*time.Minute).Err()
	if err != nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("set redis failed: %w", err)
	}
	// create token
	out.Token, err = auth.CreateToken(subToken)
	if err != nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("create token failed: %w", err)
	}

	return response.ErrCodeSuccess, out, nil
}

// Register implements service.IUserLogin.
func (s *sUserLogin) Register(ctx context.Context, in *model.RegisterInput) (codeResult int, err error) {
	// logic
	// 1. hash email
	fmt.Printf("Verify key: %s\n", in.VerifyKey)
	fmt.Printf("Verify type: %c\n", in.VerifyType)

	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))
	fmt.Printf("Hash key: %s\n", hashKey)

	// 2. check user exists in user database
	userFound, err := s.r.CheckUserBaseExists(ctx, in.VerifyKey)
	if err != nil {
		return response.ErrCodeUserHasExist, err
	}

	if userFound > 0 {
		return response.ErrCodeUserHasExist, errors.New("user has exist")
	}

	// 3. create otp
	userKey := utils.GetUserKey(hashKey) // fmt.Sprintf("u:%s:otp", key)
	otpFound, err := global.Rdb.Get(ctx, userKey).Result()

	fmt.Println("userKey::", userKey)
	fmt.Println("otpFound::", otpFound)
	fmt.Println("Err:: ", err)

	// utils...
	switch {
	case errors.Is(err, redis.Nil):
		fmt.Println("key does not exist")
	case err != nil:
		fmt.Println("get failed:: ", err)
		return response.ErrInvalidOTP, err
	case otpFound != "":
		return response.ErrCodeOTPNotExist, errors.New("OTP exists but not registered")
	}

	// 4. generate otp
	otpNew := random.GenerateSixDigitOtp()
	if in.VerifyPurpose == "TEST_USER" {
		otpNew = 123456
	}
	fmt.Printf("New OTP is ::: %d\n", otpNew)

	// 5. save otp in Redis with expiration time
	timeExpire := time.Duration(consts.TIME_OTP_REGISTER) * time.Minute
	err = global.Rdb.SetEx(ctx, userKey, strconv.Itoa(otpNew), timeExpire).Err()
	if err != nil {
		return response.ErrInvalidOTP, err
	}

	// 6. send otp
	switch in.VerifyType {
	case consts.EMAIL:
		// send email
		email := in.VerifyKey
		err = create.FactoryCreateSendTo(sendto.TYPE_SENDGRID).SendTemplateEmailOTP([]string{email}, consts.EMAIL_HOST, "otp-auth.html", map[string]interface{}{"otp": strconv.Itoa(otpNew)})
		if err != nil {
			return response.ErrSendEmailOTP, err
		}
		global.Logger.Info(fmt.Sprintf("OTP is sent to email: %s sucess", email))

		// 7. save OTP to database
		result, err := s.r.InsertOTPVerify(
			ctx,
			database.InsertOTPVerifyParams{
				VerifyOtp:     strconv.Itoa(otpNew),
				VerifyKey:     in.VerifyKey,
				VerifyKeyHash: hashKey,
				VerifyType:    sql.NullInt32{Int32: 1, Valid: true},
			},
		)
		if err != nil {
			return response.ErrInvalidOTP, err
		}

		// 8. get last id
		lastIdVerifyUser, err := result.LastInsertId()
		if err != nil {
			return response.ErrSendEmailOTP, err
		}
		global.Logger.Info(fmt.Sprintf("Last id verify user: %d", lastIdVerifyUser))
	case consts.MOBILE:
		// send sms
		// TODO
		return response.ErrCodeSuccess, nil
	}

	return response.ErrCodeSuccess, nil
}

// VerifyOTP implements service.IUserLogin.
func (s *sUserLogin) VerifyOTP(ctx context.Context, in *model.VerifyInput) (out model.VerifyOTPOutput, err error) {
	// get hash key
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))

	// get otp
	otpFound, err := global.Rdb.Get(ctx, utils.GetUserKey(hashKey)).Result()
	switch {
	case errors.Is(err, redis.Nil):
		fmt.Println("key does not exist")
	case err != nil:
		fmt.Println("get failed:: ", err)
		return out, err
	}

	if in.VerifyCode != otpFound {
		// TODO - neu sai 3 lan trong 1 phut
		return out, errors.New("OTP not match")
	}

	infoOTP, err := s.r.GetInfoOTP(ctx, hashKey)
	if err != nil {
		return out, err
	}

	// upgrade status verify
	err = s.r.UpdateUserVerificationStatus(ctx, hashKey)
	if err != nil {
		return out, err
	}

	// output
	out.Token = infoOTP.VerifyKeyHash
	out.Message = "success"

	return out, err
}

// UpdatePasswordRegister implements service.IUserLogin.
func (s *sUserLogin) UpdatePasswordRegister(ctx context.Context, in *model.UpdatePasswordInput) (userId int, err error) {
	// token is already
	infoOTP, err := s.r.GetInfoOTP(ctx, in.Token)
	if err != nil {
		return response.ErrCodeUserOTPNotExist, err
	}
	// check otp verify
	if infoOTP.IsVerified.Int32 == 0 {
		return response.ErrCodeOTPDontVerify, errors.New("OTP not verify")
	}
	// check token exists in user_base
	// update userbase password
	salt, err := crypto.GenerateSalt(16)
	if err != nil {
		return response.ErrCodeGeneratorSalt, err
	}
	passworkHash := crypto.HashPasswordWithSalt(in.Password, salt)

	userBase := database.AddUserBaseParams{
		UserAccount:  infoOTP.VerifyKey,
		UserPassword: passworkHash,
		UserSalt:     salt,
	}
	// add userBase to user_base table
	newUserBase, err := s.r.AddUserBase(ctx, userBase)
	if err != nil {
		return response.ErrCodeAddUserBase, err
	}

	user_id, err := newUserBase.LastInsertId()
	if err != nil {
		return response.ErrCodeAddUserBase, err
	}

	// add user_id to user_info table

	newUserInfo, err := s.r.AddUserHaveUserId(ctx, database.AddUserHaveUserIdParams{
		UserID:               uint64(user_id),
		UserAccount:          infoOTP.VerifyKey,
		UserNickname:         sql.NullString{String: infoOTP.VerifyKey, Valid: true},
		UserAvatar:           sql.NullString{String: "", Valid: true},
		UserState:            1,
		UserMobile:           sql.NullString{String: "", Valid: true},
		UserGender:           sql.NullInt16{Int16: 0, Valid: true},
		UserBirthday:         sql.NullTime{Time: time.Time{}, Valid: false},
		UserEmail:            sql.NullString{String: infoOTP.VerifyKey, Valid: true},
		UserIsAuthentication: 0,
	})
	if err != nil {
		return response.ErrCodeAddUserInfo, err
	}

	user_id, err = newUserInfo.LastInsertId()
	if err != nil {
		return response.ErrCodeAddUserBase, err
	}
	return int(user_id), nil
}
