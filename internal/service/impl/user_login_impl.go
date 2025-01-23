package impl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Youknow2509/go-ecommerce/global"
	"github.com/Youknow2509/go-ecommerce/internal/consts"
	"github.com/Youknow2509/go-ecommerce/internal/database"
	"github.com/Youknow2509/go-ecommerce/internal/model"
	"github.com/Youknow2509/go-ecommerce/internal/service"
	"github.com/Youknow2509/go-ecommerce/internal/utils"
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

// Login implements service.IUserLogin.
func (s *sUserLogin) Login(ctx context.Context) error {
	return nil
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
	fmt.Println("Err:: ",err)

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
func (s *sUserLogin) VerifyOTP(ctx context.Context) error {
	return nil
}
