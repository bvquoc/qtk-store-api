package userbiz

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/component/generator"
	"qtk-store-api/component/hasher"
	"qtk-store-api/middleware"
	"qtk-store-api/module/user/usermodel"
)

type CreateUserRepo interface {
	CheckRoleExist(ctx context.Context, roleId string) error
	CreateUser(ctx context.Context, data *usermodel.ReqCreateUser) error
}

type createUserBiz struct {
	gen       generator.IdGenerator
	repo      CreateUserRepo
	hasher    hasher.Hasher
	requester middleware.Requester
}

func NewCreateUserBiz(
	gen generator.IdGenerator,
	repo CreateUserRepo,
	hasher hasher.Hasher,
	requester middleware.Requester) *createUserBiz {
	return &createUserBiz{
		gen:       gen,
		repo:      repo,
		hasher:    hasher,
		requester: requester,
	}
}

func (biz *createUserBiz) CreateUser(
	ctx context.Context,
	data *usermodel.ReqCreateUser) error {
	if biz.requester.GetRole().Id != common.RoleAdminId {
		return usermodel.ErrUserCreateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.repo.CheckRoleExist(ctx, data.RoleId); err != nil {
		return err
	}

	salt := common.GenSalt(50)
	data.Password = biz.hasher.Hash(common.DefaultPass + salt)
	data.Salt = salt

	if err := handleUserId(biz.gen, data); err != nil {
		return err
	}

	if err := biz.repo.CreateUser(ctx, data); err != nil {
		return err
	}

	return nil
}

func handleUserId(gen generator.IdGenerator, data *usermodel.ReqCreateUser) error {
	id, err := gen.GenerateId()
	if err != nil {
		return err
	}

	data.Id = id
	return nil
}
