package productrepo

import (
	"context"
	"qtk-store-api/module/product/productmodel"
)

type CreateProductStore interface {
	// CreateProduct(ctx context.Context, productGeneral *productmodel.Product, productInfo *productmodel.ProductInfo) error
}

type createProductRepo struct {
	store CreateProductStore
}

func NewCreateProductRepo(store CreateProductStore) *createProductRepo {
	return &createProductRepo{store: store}
}

func (biz *createProductRepo) CreateProduct(ctx context.Context, data *productmodel.ReqCreateProduct) error {
	// if err := biz.store.CreateProduct(ctx, data); err != nil {
	// 	return err
	// }

	return nil
}
