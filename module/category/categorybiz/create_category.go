package categorybiz

import (
	"context"
	"qtk-store-api/component/generator"
	"qtk-store-api/constants"
	"qtk-store-api/middleware"
	"qtk-store-api/module/category/categorymodel"
)

type CreateCategoryRepo interface {
	CreateCategory(ctx context.Context, data *categorymodel.Category) error
}

type createCategoryBiz struct {
	gen       generator.IdGenerator
	repo      CreateCategoryRepo
	requester middleware.Requester
}

func NewCreateCategoryBiz(gen generator.IdGenerator, repo CreateCategoryRepo, requester middleware.Requester) *createCategoryBiz {
	return &createCategoryBiz{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *createCategoryBiz) CreateCategory(ctx context.Context, data *categorymodel.Category) error {
	if !biz.requester.IsHasFeature(constants.CategoryCreateFeatureCode) {
		return categorymodel.ErrCategoryCreateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := handleCategoryId(biz.gen, data); err != nil {
		return err
	}
	if err := biz.repo.CreateCategory(ctx, data); err != nil {
		return err
	}

	return nil
}

func handleCategoryId(gen generator.IdGenerator, data *categorymodel.Category) error {
	id, err := gen.IdProcess(&data.Id)
	if err != nil {
		return err
	}
	data.Id = *id
	return nil
}
