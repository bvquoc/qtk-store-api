package supplierbiz

import (
	"context"
	"qtk-store-api/component/generator"
	"qtk-store-api/constants"
	"qtk-store-api/middleware"
	"qtk-store-api/module/supplier/suppliermodel"
)

type CreateSupplierRepo interface {
	CreateSupplier(
		ctx context.Context,
		data *suppliermodel.ReqCreateSupplier,
	) error
}

type createSupplierBiz struct {
	gen       generator.IdGenerator
	repo      CreateSupplierRepo
	requester middleware.Requester
}

func NewCreateSupplierBiz(
	gen generator.IdGenerator,
	repo CreateSupplierRepo,
	requester middleware.Requester) *createSupplierBiz {
	return &createSupplierBiz{
		gen:       gen,
		repo:      repo,
		requester: requester,
	}
}

func (biz *createSupplierBiz) CreateSupplier(
	ctx context.Context,
	data *suppliermodel.ReqCreateSupplier) error {
	if !biz.requester.IsHasFeature(constants.SupplierCreateFeatureCode) {
		return suppliermodel.ErrSupplierCreateNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	data.Round()

	if err := handleSupplierId(biz.gen, data); err != nil {
		return err
	}
	if err := biz.repo.CreateSupplier(ctx, data); err != nil {
		return err
	}

	return nil
}

func handleSupplierId(gen generator.IdGenerator, data *suppliermodel.ReqCreateSupplier) error {
	id, err := gen.IdProcess(data.Id)
	if err != nil {
		return err
	}

	data.Id = id

	return nil
}
