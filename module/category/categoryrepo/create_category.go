package categoryrepo

import (
	"context"
	"qtk-store-api/module/category/categorymodel"
)

type CreateCategory interface {
	CreateCategory(ctx context.Context, data *categorymodel.Category) error
}

type createCategoryRepo struct {
	store CreateCategory
}

func NewCreateCategoryRepo(store CreateCategory) *createCategoryRepo {
	return &createCategoryRepo{store: store}
}

func (biz *createCategoryRepo) CreateCategory(ctx context.Context, data *categorymodel.Category) error {
	if err := biz.store.CreateCategory(ctx, data); err != nil {
		return err
	}

	return nil
}
