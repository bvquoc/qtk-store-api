package authorrepo

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/module/author/authormodel"
)

type ListAuthorStore interface {
	ListAuthor(ctx context.Context, filter *authormodel.Filter, propertiesContainSearchKey []string, paging *common.Paging) ([]authormodel.Author, error)
}

type listAuthorRepo struct {
	store ListAuthorStore
}

func NewListAuthorRepo(store ListAuthorStore) *listAuthorRepo {
	return &listAuthorRepo{store: store}
}

func (repo *listAuthorRepo) ListAuthor(ctx context.Context, filter *authormodel.Filter, paging *common.Paging) ([]authormodel.Author, error) {
	result, err := repo.store.ListAuthor(ctx, filter, []string{"name"}, paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
