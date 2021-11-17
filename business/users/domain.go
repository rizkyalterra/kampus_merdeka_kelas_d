package users

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Email     string
	Name      string
	Password  string
	Token     string
}

type UserUsecaseInterface interface {
	Login(domain Domain, ctx context.Context) (Domain, error)
	GetAllUsers(ctx context.Context) ([]Domain, error)
}

type UserRepoInterface interface {
	Login(domain Domain, ctx context.Context) (Domain, error)
	// GetAllUsers(ctx context.Context) ([]Domain, error)
}
