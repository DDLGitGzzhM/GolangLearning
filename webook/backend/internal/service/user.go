package service

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"GolangLearning/webook/backend/internal/domain"
	"GolangLearning/webook/backend/internal/repository"
)

var UserEmailDuplicate = repository.UserEmailDuplicate

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(UserRepo *repository.UserRepository) *UserService {
	return &UserService{
		UserRepo: UserRepo,
	}
}

func (svc *UserService) SignUp(ctx context.Context, user domain.User) error {
	encryptPwd, err := bcrypt.GenerateFromPassword([]byte(user.PassWord), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PassWord = string(encryptPwd)
	return svc.UserRepo.Create(ctx, user)
}
