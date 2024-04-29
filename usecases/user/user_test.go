package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go-wishlist-api/constant"
	"go-wishlist-api/drivers/mysql/mocks"
	"go-wishlist-api/entities"
	"go-wishlist-api/utils"
)

func TestRegister(t *testing.T) {

	expectedUser := &entities.User{
		ID:       1,
		Name:     "alfara",
		Email:    "alfara@gmail.com",
		Password: "password",
	}

	req := &entities.User{
		Name:     "alfara",
		Email:    "alfara@gmail.com",
		Password: "password",
	}

	t.Run("Success", func(t *testing.T) {

		mockRepo := new(mocks.MockUserRepository)
		uc := NewUserUseCase(mockRepo)

		mockRepo.On("FindByEmail", "alfara@gmail.com").Return(nil, nil)
		mockRepo.On("Create", mock.Anything).Return(expectedUser, nil)

		newUser, err := uc.Register(req)

		assert.NoError(t, err)
		assert.Equal(t, newUser.Email, req.Email)
	})

	t.Run("Empty input", func(t *testing.T) {

		mockRepo := new(mocks.MockUserRepository)
		uc := NewUserUseCase(mockRepo)

		req := &entities.User{}

		newUser, err := uc.Register(req)

		assert.Error(t, err)

		assert.Nil(t, newUser)
		assert.IsType(t, constant.ErrEmptyInput, err)
	})

	t.Run("Email already used", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		uc := NewUserUseCase(mockRepo)

		user := &entities.User{
			ID:       1,
			Name:     "alfara",
			Email:    "alfara@gmail.com",
			Password: "password",
		}

		expectedError := constant.ErrUserExist
		mockRepo.On("FindByEmail", "alfara@gmail.com").Return(user, nil)
		newUser, err := uc.Register(req)
		assert.Error(t, err)
		assert.Nil(t, newUser)
		assert.EqualError(t, err, expectedError.Error())
	})

	t.Run("Internal Server Error", func(t *testing.T) {

		mockRepo := new(mocks.MockUserRepository)
		uc := NewUserUseCase(mockRepo)

		expectedError := constant.ErrUserCreate

		mockRepo.On("FindByEmail", "alfara@gmail.com").Return(nil, nil)
		mockRepo.On("Create", mock.Anything).Return(nil, expectedError)

		newUser, err := uc.Register(req)

		assert.Error(t, err)
		assert.Nil(t, newUser)
		assert.EqualError(t, err, expectedError.Error())
	})
}

func TestLogin(t *testing.T) {

	req := &entities.User{
		Email:    "alfara@gmail.com",
		Password: "password",
	}

	t.Run("Success", func(t *testing.T) {

		mockRepo := new(mocks.MockUserRepository)
		uc := NewUserUseCase(mockRepo)

		password, _ := utils.HashPassword("password")

		expectedUser := &entities.User{
			Email:    "alfara@gmail.com",
			Password: password,
		}

		mockRepo.On("FindByEmail", "alfara@gmail.com").Return(expectedUser, nil)

		token, err := uc.Login(req)

		assert.NoError(t, err)
		assert.NotEmpty(t, token)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Wrong Email", func(t *testing.T) {

		mockRepo := new(mocks.MockUserRepository)
		uc := NewUserUseCase(mockRepo)

		mockRepo.On("FindByEmail", "alfara@gmail.com").Return(nil, nil)

		_, err := uc.Login(req)

		assert.Error(t, err)
		assert.Equal(t, constant.ErrUserNotExist.Error(), err.Error())
	})

	t.Run("Wrong Password", func(t *testing.T) {

		mockRepo := new(mocks.MockUserRepository)
		uc := NewUserUseCase(mockRepo)

		password, _ := utils.HashPassword("password")

		expectedUser := &entities.User{
			Email:    "alfara@gmail.com",
			Password: password,
		}

		mockRepo.On("FindByEmail", "alfara@gmail.com").Return(expectedUser, nil)

		req := &entities.User{
			Email:    "alfara@gmail.com",
			Password: "wrongpassword",
		}

		_, err := uc.Login(req)

		assert.Error(t, err)
		assert.Equal(t, constant.ErrUserWrongInput.Error(), err.Error())
		mockRepo.AssertExpectations(t)
	})
}
