package request

import "keld/business/users"

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *UserLogin) ToDomain() *users.Domain {
	return &users.Domain{
		Email:    user.Email,
		Password: user.Password,
	}
}
