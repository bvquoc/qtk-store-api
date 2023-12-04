package authorbiz

import (
	"context"
	"qtk-store-api/component/generator"
	"qtk-store-api/constants"
	"qtk-store-api/middleware"
	"qtk-store-api/module/author/authormodel"
)

type CreateAuthorRepo interface {
	CreateAuthor(ctx context.Context, data *authormodel.Author) error
}

type createAuthorBiz struct {
	gen       generator.IdGenerator
	repo      CreateAuthorRepo
	requester middleware.Requester
}

func NewCreateAuthorBiz(gen generator.IdGenerator, repo CreateAuthorRepo, requester middleware.Requester) *createAuthorBiz {
	return &createAuthorBiz{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *createAuthorBiz) CreateAuthor(ctx context.Context, data *authormodel.Author) error {
	if !biz.requester.IsHasFeature(constants.AuthorCreateFeatureCode) {
		return authormodel.ErrAuthorCreateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := handleAuthorId(biz.gen, data); err != nil {
		return err
	}
	if err := biz.repo.CreateAuthor(ctx, data); err != nil {
		return err
	}

	return nil
}

func handleAuthorId(gen generator.IdGenerator, data *authormodel.Author) error {
	id, err := gen.IdProcess(&data.Id)
	if err != nil {
		return err
	}
	data.Id = *id
	return nil
}
