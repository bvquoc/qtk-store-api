package supplierrepo

import (
	"context"
	"qtk-store-api/module/supplier/suppliermodel"
)

type CreateSupplierStore interface {
	CreateSupplier(ctx context.Context, data *suppliermodel.ReqCreateSupplier) error
}

type createSupplierRepo struct {
	store CreateSupplierStore
}

func NewCreateSupplierRepo(store CreateSupplierStore) *createSupplierRepo {
	return &createSupplierRepo{store: store}
}

func (biz *createSupplierRepo) CreateSupplier(
	ctx context.Context,
	data *suppliermodel.ReqCreateSupplier) error {
	if err := biz.store.CreateSupplier(ctx, data); err != nil {
		return err
	}

	return nil
}
