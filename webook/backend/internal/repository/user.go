package repository

import (
	"context"

	"GolangLearning/webook/backend/internal/domain"
	"GolangLearning/webook/backend/internal/repository/dao"
)

var UserEmailDuplicate = dao.UserEmailDuplicate

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{dao: dao}
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	return r.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.PassWord,
	})
}
