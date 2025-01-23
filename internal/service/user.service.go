package service

import (
	"fmt"
	"time"

	"github.com/Youknow2509/go-ecommerce/global"
	rp "github.com/Youknow2509/go-ecommerce/internal/repo"
	"github.com/Youknow2509/go-ecommerce/internal/utils/crypto"
	"github.com/Youknow2509/go-ecommerce/internal/utils/random"
	"github.com/Youknow2509/go-ecommerce/response"
)

// type UserService struct {
// 	userRepo *rp.UserRepo
// }

// // NewUserService creates a new UserService
// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepo: rp.NewUserRepo(),
// 	}
// }

// // Get User Information Services
// func (u *UserService) GetUserInfoService() string {
// 	return u.userRepo.GetInfoUser()
// }

// Interface Version
type IUserService interface {
	RegisterService(email string, purpose string) int
}

type userService struct {
	userRepo     rp.IUserRepository
	userAuthRepo rp.IUserAuthRepository
	// ...
}

func NewUserService(
	userRepo rp.IUserRepository,
	userAuthRepo rp.IUserAuthRepository,
) IUserService {
	return &userService{
		userRepo:     userRepo,
		userAuthRepo: userAuthRepo,
	}
}

// RegisterService implements IUserService.
func (u *userService) RegisterService(email string, purpose string) int {
	// 0. hash email
	hashEmail := crypto.GetHash(email)
	fmt.Printf("Hash email is ::: %s\n", hashEmail)
	// 5. check opt is available
	// TODO

	// 6. user spam ...
	// TODO

	// 1. check email exists in db
	if u.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExist
	}
	// 2. new otp ...
	otp := random.GenerateSixDigitOtp()
	if purpose == "TEST_USER" {
		otp = 123456
	}

	fmt.Printf("OTP is ::: %d\n", otp)
	// 3. save otp in Redis with expiration time
	err := u.userAuthRepo.AddOtp(hashEmail, otp, int64(10*time.Minute))
	if err != nil {
		return response.ErrInvalidOTP
	}
	// 4. send otp to email 
	// err = create.FactoryCreateSendTo(sendto.TYPE_SENDGRID).SendTextEmailOTP([]string{email}, "lytranvinh.work@gmail.com", strconv.Itoa(otp))
	// err = create.FactoryCreateSendTo(sendto.TYPE_SENDGRID).SendTemplateEmailOTP([]string{email}, "lytranvinh.work@gmail.com", "otp-auth.html", map[string]interface{}{"otp": strconv.Itoa(otp)})
	// err = create.FactoryCreateSendTo(sendto.TYPE_API).SendAPIEmailOTP(email, "lytranvinh.work@gmail.com", strconv.Itoa(otp))
	if err != nil {
		return response.ErrSendEmailOTP
	}
	global.Logger.Info(fmt.Sprintf("OTP is sent to email: %s sucess", email))

	// // handle send to kafka
	// body := make(map[string]interface{})
	// body["email"] = email
	// body["otp"] = otp
	// // requestBody
	// requestBody, _ := json.Marshal(body)
	// // create message in kafaka
	// msg := kafka.Message{
    //     Key:   []byte("otp-auth"),
    //     Value: []byte(requestBody),
	// 	Time: time.Now(),
    // }
	// err = global.KafkaProducer.WriteMessages(context.Background(), msg)
	// if err!= nil {
    //     global.Logger.Error("Error sending message to Kafka: ", zap.Error(err))
    //     return response.ErrSendEmailOTP
    // }
    // 7. return success
    return response.ErrCodeSuccess

}
