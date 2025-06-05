package dto

// list type dto
type (
	LoginDto struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6,max=100"`
	}
)
