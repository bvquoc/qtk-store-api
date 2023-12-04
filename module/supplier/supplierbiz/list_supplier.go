package supplierbiz

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/constants"
	"qtk-store-api/middleware"
	"qtk-store-api/module/supplier/suppliermodel"
	"qtk-store-api/module/supplier/suppliermodel/filter"
)

type ListSupplierRepo interface {
	ListSupplier(
		ctx context.Context,
		filter *filter.Filter,
		paging *common.Paging,
	) ([]suppliermodel.Supplier, error)
}

type listSupplierBiz struct {
	repo      ListSupplierRepo
	requester middleware.Requester
}

func NewListSupplierRepo(
	repo ListSupplierRepo,
	requester middleware.Requester) *listSupplierBiz {
	return &listSupplierBiz{repo: repo, requester: requester}
}

func (biz *listSupplierBiz) ListSupplier(
	ctx context.Context,
	filter *filter.Filter,
	paging *common.Paging) ([]suppliermodel.Supplier, error) {
	if !biz.requester.IsHasFeature(constants.SupplierViewFeatureCode) {
		return nil, suppliermodel.ErrSupplierViewNoPermission
	}

	result, err := biz.repo.ListSupplier(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
