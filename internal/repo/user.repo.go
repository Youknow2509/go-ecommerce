package repo

// type UserRepo struct {
// }

// // New UserRepo creates a new UserRepo
// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// // Get Information User
// func (u *UserRepo) GetInfoUser() string {
// 	return "some one"
// }

// Interface Version
type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {

}

func NewUserRepository() IUserRepository {
    return &userRepository{}
}

func (u *userRepository) GetUserByEmail(email string) bool {
	return true
}
