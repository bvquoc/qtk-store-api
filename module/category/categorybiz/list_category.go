package categorybiz

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/constants"
	"qtk-store-api/middleware"
	"qtk-store-api/module/category/categorymodel"
)

type ListCategoryRepo interface {
	ListCategory(ctx context.Context, filter *categorymodel.Filter, paging *common.Paging) ([]categorymodel.Category, error)
}

type listCategoryBiz struct {
	repo      ListCategoryRepo
	requester middleware.Requester
}

func NewListCategoryRepo(repo ListCategoryRepo, requester middleware.Requester) *listCategoryBiz {
	return &listCategoryBiz{repo: repo, requester: requester}
}

func (biz *listCategoryBiz) ListCategory(ctx context.Context, filter *categorymodel.Filter, paging *common.Paging) ([]categorymodel.Category, error) {
	if !biz.requester.IsHasFeature(constants.CategoryViewFeatureCode) {
		return nil, categorymodel.ErrCategoryViewNoPermission
	}

	result, err := biz.repo.ListCategory(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
