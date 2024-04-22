package user

import (
	"context"

	"github.com/ThailanTec/my-boiler-plate/internal/domain/user"
	"github.com/ThailanTec/my-boiler-plate/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user *user.User) error
	GetUser(ctx context.Context, userID uint) (*user.User, error)
	UpdateUser(ctx context.Context, user *user.User) error
	DeleteUser(ctx context.Context, userID uint) error
}

type userServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userServiceImpl{userRepository: userRepository}
}

func (s *userServiceImpl) CreateUser(ctx context.Context, user *user.User) error {
	// Validate user data before creating (implement validation logic)
	// Call the UserRepository's Create method to insert the user into the database
	return nil
}

func (s *userServiceImpl) GetUser(ctx context.Context, userID uint) (*user.User, error) {
	// Call the UserRepository's Read method to fetch the user with the given ID
	return nil, nil
}

func (s *userServiceImpl) UpdateUser(ctx context.Context, user *user.User) error {
	// Validate user data before updating (implement validation logic)
	// Call the UserRepository's Update method to update the user details
	return nil
}

func (s *userServiceImpl) DeleteUser(ctx context.Context, userID uint) error {
	// Call the UserRepository's Delete method to remove the user with the given ID
	return nil
}
