package service

import (
	"context"

	"GolangLearning/webook/backend/internal/domain"
	"GolangLearning/webook/backend/internal/repository"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(UserRepo *repository.UserRepository) *UserService {
	return &UserService{
		UserRepo: UserRepo,
	}
}

func (svc *UserService) SignUp(ctx context.Context, user domain.User) error {
	return svc.UserRepo.Create(ctx, user)
}
