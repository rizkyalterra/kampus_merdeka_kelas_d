package users

import (
	"context"
	"errors"
	_middleware "keld/app/middleware"
	"time"
)

type UserUseCase struct {
	repo UserRepoInterface
	ctx  time.Duration
	jwt  *_middleware.ConfigJWT
}

func NewUsecase(userRepo UserRepoInterface, contextTimeout time.Duration) UserUsecaseInterface {
	return &UserUseCase{
		repo: userRepo,
		ctx:  contextTimeout,
		// jwt:  configJWT,
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

	// user.Token = usecase.jwt.GenererateToken(user.Id)

	return user, nil
}

func (usecase *UserUseCase) GetAllUsers(ctx context.Context) ([]Domain, error) {
	return []Domain{}, nil
}

// 5, 7 => 12
func Addition(a, b int) int {
	sum := a + b
	if sum < 0 {
		sum = 0
	}
	return sum
}
