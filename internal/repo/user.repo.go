package repo

type UserRepo struct {
}

// New UserRepo creates a new UserRepo
func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

// Get Information User
func (u *UserRepo) GetInfoUser() string {
	return "some one"
}
