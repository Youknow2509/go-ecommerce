package services

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/Youknow2509/go-ecommerce/consts"
	"github.com/Youknow2509/go-ecommerce/internal/auth/domain/cache"
	"github.com/Youknow2509/go-ecommerce/internal/user/application/model"
	"github.com/Youknow2509/go-ecommerce/internal/user/domain/model/entity"
	"github.com/Youknow2509/go-ecommerce/internal/user/domain/responsitory"
	"github.com/Youknow2509/go-ecommerce/internal/user/infrastructure/mq"
	"github.com/Youknow2509/go-ecommerce/internal/utils"
	"github.com/Youknow2509/go-ecommerce/internal/utils/crypto"
	"github.com/Youknow2509/go-ecommerce/internal/utils/random"
	"github.com/Youknow2509/go-ecommerce/response"
)

type (
	UserServiceImpl struct {
		repo             responsitory.IUserResponsitory
		distributedCache cache.ICacheService
		loacalCache      cache.ICacheService
		sendOtpService   mq.IKafkaService
	}
)

// ######################################################

// CreatePasswordUserService implements IUserService.
func (u *UserServiceImpl) CreatePasswordUserService(ctx context.Context, input *model.InputUserCreatePassword) (int, error) {
	/**
	 * - check token in distributed cache
	 * - create user base
	 * - create user info
	 * ...
	 */
	accountNameHash := crypto.GetHash(input.AccountName)
	key := utils.GetKeyUserRegisterToken(accountNameHash)
	// check token in distributed cache
	dataCache, err := u.distributedCache.Get(ctx, key)
	if err != nil {
		return response.ErrCodeCheckUserRegisterCache, err
	}
	if dataCache == nil {
		return response.ErrCodeUserRegisterTokenNotFound, nil
	}
	if dataCache != input.TokenCreatePassword {
		return response.ErrCodeVerifyTokenCreatePasswordFail, nil
	}
	// prepare data
	salt, _ := crypto.GenerateSalt(32)
	hashPassword := crypto.HashPasswordWithSalt(input.Password, salt)
	// create user base
	u.repo.AddUserBase(
		ctx,
		&entity.InputAddUserBase{
			UserAccount:  input.AccountName,
			UserPassword: hashPassword,
			UserSalt:     salt,
		},
	)
	if err != nil {
		return response.ErrCodeCreateUserBase, err
	}
	// create user info
	err = u.repo.AddUserInfoAutoId(
		ctx,
		&entity.InputAddUserAutoUserId{
			UserAccount:          input.AccountName,
			UserNickname:         input.AccountName,
			UserAvatar:           "",
			UserState:            0,
			UserMobile:           "",
			UserGender:           0,
			UserBirthday:         time.Now(),
			UserEmail:            "",
			UserIsAuthentication: 0,
		},
	)
	if err != nil {
		return response.ErrCodeCreateUserInfo, err
	}

	return response.ErrCodeSuccess, nil
}

// RegisterUserService implements IUserService.
func (u *UserServiceImpl) RegisterUserService(ctx context.Context, input *model.InputUserRegister) (int, error) {
	/**
	 * - message queue - check ....
	 * - Bloom Filter
	 * - db
	 * - distributed cache save otp code, save local cache
	 * - send otp code to user
	 * - return response code
	 */
	// prepare data
	accountNameHash := crypto.GetHash(strings.ToLower(input.AcountName))
	otpCode := random.GenerateSixDigitOtp()
	key := utils.GetKeyUserRegisterCache(accountNameHash)
	ttlTimeCache := time.Duration(consts.TTL_USER_REGISTER_CACHE) * time.Second
	// check distributed cache
	dataCache, err := u.distributedCache.Get(ctx, key)
	if err != nil {
		return response.ErrCodeCheckUserRegisterCache, err
	}
	if dataCache != nil && dataCache != "" {
		// block, spam, ....
		return response.ErrCodeSpamUserRegister, nil
	}
	// Check db
	ok, err := u.repo.CheckUserBaseExists(ctx, input.AcountName)
	if err != nil {
		return response.ErrCodeCheckUserBaseExists, err
	}
	if ok {
		return response.ErrCodeUserHasExist, nil
	}
	// save distributed cache
	err = u.distributedCache.SetWithTTL(ctx, key, otpCode, ttlTimeCache)
	if err != nil {
		return response.ErrCodeCheckUserRegisterCache, err
	}
	// send
	err = u.sendOtpService.Publish(
		ctx,
		consts.TOPIC_SEND_OTP,
		strconv.Itoa(input.VerifyType),
		0, // partition (default to 0)
		[]byte(strconv.Itoa(otpCode)),
	)
	if err != nil {
		// remove cache
		_ = u.distributedCache.Del(ctx, key)
		return response.ErrSendEmailOTP, err
	}
	return response.ErrCodeSuccess, nil
}

// VerifyRegisterUserService implements IUserService.
func (u *UserServiceImpl) VerifyRegisterUserService(ctx context.Context, input *model.InputUserVerifyRegister) (*model.OutputUserVerifyRegister, int, error) {
	/**
	 * - check distributed cache exist user name -> ok, spam, ....
	 * - check otp code
	 * - create token create password
	 * - write token to distributed cache
	 * - return response
	 */
	accountNameHash := crypto.GetHash(strings.ToLower(input.AcountName))
	key := utils.GetKeyUserRegisterCache(accountNameHash)
	// check distributed cache
	dataCache, err := u.distributedCache.Get(ctx, key)
	if err != nil {
		return nil, response.ErrCodeCheckUserRegisterCache, err
	}
	if dataCache == nil {
		return nil, response.ErrCodeUserRegisterCacheNotFound, nil
	}
	// spam, block, ....
	if dataCache != input.VerifyCode {
		return nil, response.ErrCodeVerifyOTPFail, nil
	}
	// create token
	token := utils.GetTokenUserRegister(accountNameHash)
	keyToken := utils.GetKeyUserRegisterToken(accountNameHash)
	// write token to distributed cache
	err = u.distributedCache.SetWithTTL(ctx, keyToken, token, time.Duration(consts.TTL_TOKEN_CREATE_PASSWORD_REGISTER)*time.Second)
	if err != nil {
		return nil, response.ErrCodeCheckUserRegisterCache, err
	}
	return &model.OutputUserVerifyRegister{
		Token: token,
	}, response.ErrCodeSuccess, nil
}

// ######################################################
// init and impl
func NewUserServiceImpl(
	repo responsitory.IUserResponsitory,
	distributedCache cache.ICacheService,
	loacalCache cache.ICacheService,
	sendOtpService mq.IKafkaService,
) IUserService {
	return &UserServiceImpl{
		repo:             repo,
		distributedCache: distributedCache,
		loacalCache:      loacalCache,
		sendOtpService:   sendOtpService,
	}
}
