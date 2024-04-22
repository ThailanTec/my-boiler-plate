package repository

import (
	"context"

	"github.com/ThailanTec/my-boiler-plate/internal/config"
	"github.com/ThailanTec/my-boiler-plate/internal/domain/user"
)

type UserRepository interface {
	Create(ctx context.Context, user *user.User) error
	Read(ctx context.Context, userID uint) (*user.User, error)
	Update(ctx context.Context, user *user.User) error
	Delete(ctx context.Context, userID uint) error
}

type userRepositoryImpl struct {
	db *config.Postgres
}

func NewUserRepository(db *config.Postgres) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) Create(ctx context.Context, user *user.User) error {
	return nil
}

func (r *userRepositoryImpl) Read(ctx context.Context, userID uint) (*user.User, error) {
	return nil, nil
}

func (r *userRepositoryImpl) Update(ctx context.Context, user *user.User) error {
	return nil
}

func (r *userRepositoryImpl) Delete(ctx context.Context, userID uint) error {
	return nil
}
