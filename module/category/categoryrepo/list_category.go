package categoryrepo

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/module/category/categorymodel"
)

type ListCategoryStore interface {
	ListCategory(ctx context.Context, filter *categorymodel.Filter, propertiesContainSearchKey []string, paging *common.Paging) ([]categorymodel.Category, error)
}

type listCategoryRepo struct {
	store ListCategoryStore
}

func NewListCategoryRepo(store ListCategoryStore) *listCategoryRepo {
	return &listCategoryRepo{store: store}
}

func (repo *listCategoryRepo) ListCategory(ctx context.Context, filter *categorymodel.Filter, paging *common.Paging) ([]categorymodel.Category, error) {
	result, err := repo.store.ListCategory(ctx, filter, []string{"name"}, paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
