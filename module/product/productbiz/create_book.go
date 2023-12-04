package productbiz

import (
	"context"
	"qtk-store-api/component/generator"
	"qtk-store-api/middleware"
	"qtk-store-api/module/product/productmodel"
)

type CreateProductRepo interface {
	CreateProduct(ctx context.Context, data *productmodel.Product) error
}

type createProductBiz struct {
	gen       generator.IdGenerator
	repo      CreateProductRepo
	requester middleware.Requester
}

func NewCreateProductBiz(
	gen generator.IdGenerator,
	repo CreateProductRepo,
	requester middleware.Requester) *createProductBiz {
	return &createProductBiz{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *createProductBiz) CreateProduct(ctx context.Context, data *productmodel.Product) error {
	//if !biz.requester.IsHasFeature(common.ProductCreateFeatureCode) {
	//	return productmodel.ErrProductCreateNoPermission
	//}
	//
	//if err := data.Validate(); err != nil {
	//	return err
	//}
	//
	//if err := handleProductId(biz.gen, data); err != nil {
	//	return err
	//}
	//if err := biz.repo.CreateProduct(ctx, data); err != nil {
	//	return err
	//}
	//
	return nil
}

func handleProductId(gen generator.IdGenerator, data *productmodel.Product) error {
	id, err := gen.IdProcess(&data.ID)
	if err != nil {
		return err
	}
	data.ID = *id
	return nil
}
