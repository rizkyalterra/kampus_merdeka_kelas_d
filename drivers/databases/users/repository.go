package users

import (
	"context"
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
		return users.Domain{}, err
	}
	return userDb.ToDomain(), nil
}
func (repo *UserRepository) GetAllUsers(ctx context.Context) ([]users.Domain, error) {
	return []users.Domain{}, nil
}
