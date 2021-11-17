package users_test

import (
	"context"
	"errors"
	"keld/business/users"
	"keld/business/users/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepoInterfaceMock mocks.UserRepoInterface
var userUseCaseInterface users.UserUsecaseInterface
var userDataDummyLogin users.Domain

func setup() {
	userUseCaseInterface = users.NewUsecase(&userRepoInterfaceMock, time.Hour*1)
	userDataDummyLogin = users.Domain{
		Id:       1,
		Name:     "Alterra",
		Email:    "alterra@gmail.com",
		Password: "123",
		Token:    "",
	}
}

func TestLogin(t *testing.T) {
	setup()
	t.Run("Success Login", func(t *testing.T) {
		userRepoInterfaceMock.On("Login", mock.AnythingOfType("users.Domain"), mock.Anything).Return(userDataDummyLogin, nil).Once()

		var requestLoginDomain = users.Domain{
			Email:    "Alterra@gmail.com",
			Password: "123",
		}
		domain, err := userUseCaseInterface.Login(requestLoginDomain, context.Background())

		assert.Equal(t, nil, err)
		assert.Equal(t, userDataDummyLogin, domain)
	})

	t.Run("Login with Email Empty", func(t *testing.T) {
		var requestLoginDomain = users.Domain{
			Email:    "",
			Password: "123",
		}
		domain, err := userUseCaseInterface.Login(requestLoginDomain, context.Background())

		assert.Equal(t, "Email Empty", err.Error())
		assert.Equal(t, users.Domain{}, domain)
	})

	t.Run("Login with Password Empty", func(t *testing.T) {
		var requestLoginDomain = users.Domain{
			Email:    "alterra@gmail.com",
			Password: "",
		}
		domain, err := userUseCaseInterface.Login(requestLoginDomain, context.Background())

		assert.Equal(t, "Password Empty", err.Error())
		assert.Equal(t, users.Domain{}, domain)
	})

	t.Run("Email not found in database", func(t *testing.T) {
		userRepoInterfaceMock.On("Login", mock.AnythingOfType("users.Domain"), mock.Anything).Return(users.Domain{}, errors.New("Email not found")).Once()

		var requestLoginDomain = users.Domain{
			Email:    "Alterra2@gmail.com",
			Password: "123",
		}
		domain, err := userUseCaseInterface.Login(requestLoginDomain, context.Background())

		assert.Equal(t, errors.New("Email not found"), err)
		assert.Equal(t, users.Domain{}, domain)
	})

}
