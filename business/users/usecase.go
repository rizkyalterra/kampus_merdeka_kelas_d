package users

import (
	"context"
	"errors"
	"time"
)

type UserUseCase struct {
	repo UserRepoInterface
	ctx  time.Duration
}

func NewUsecase(userRepo UserRepoInterface, contextTimeout time.Duration) UserUsecaseInterface {
	return &UserUseCase{
		repo: userRepo,
		ctx:  contextTimeout,
	}
}

func (usecase *UserUseCase) Login(domain Domain, ctx context.Context) (Domain, error) {
	if domain.Email == "" {
		return Domain{}, errors.New("Email Empty")
	}
	if domain.Password == "" {
		return Domain{}, errors.New("Password Empty")
	}
	user, err := usecase.repo.Login(domain, ctx)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (usecase *UserUseCase) GetAllUsers(ctx context.Context) ([]Domain, error) {
	return []Domain{}, nil
}
