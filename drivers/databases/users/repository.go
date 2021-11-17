package users

import (
	"context"
	"errors"
	"keld/business/users"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(gormDb *gorm.DB) users.UserRepoInterface {
	return &UserRepository{
		db: gormDb,
	}
}

func (repo *UserRepository) Login(domain users.Domain, ctx context.Context) (users.Domain, error) {
	userDb := FromDomain(domain)

	err := repo.db.Where("email = ? AND password = ?", userDb.Email, userDb.Password).First(&userDb).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return users.Domain{}, errors.New("Email not found")
		}
		return users.Domain{}, errors.New("Error in database")
	}
	return userDb.ToDomain(), nil
}

// func (repo *UserRepository) GetAllUsers(ctx context.Context) ([]users.Domain, error) {
// 	return []users.Domain{}, nil
// }
