package authorrepo

import (
	"context"
	"qtk-store-api/module/author/authormodel"
)

type CreateAuthor interface {
	CreateAuthor(ctx context.Context, data *authormodel.Author) error
}

type createAuthorRepo struct {
	store CreateAuthor
}

func NewCreateAuthorRepo(store CreateAuthor) *createAuthorRepo {
	return &createAuthorRepo{store: store}
}

func (biz *createAuthorRepo) CreateAuthor(ctx context.Context, data *authormodel.Author) error {
	if err := biz.store.CreateAuthor(ctx, data); err != nil {
		return err
	}

	return nil
}
