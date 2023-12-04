package supplierbiz

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/constants"
	"qtk-store-api/middleware"
	"qtk-store-api/module/supplier/suppliermodel"
	"qtk-store-api/module/supplier/suppliermodel/filter"
)

type SeeSupplierDebtRepo interface {
	SeeSupplierDebt(
		ctx context.Context,
		supplierId string,
		filter *filter.SupplierDebtFilter,
		paging *common.Paging) (*suppliermodel.ResDebtSupplier, error)
}

type seeSupplierDebtBiz struct {
	repo      SeeSupplierDebtRepo
	requester middleware.Requester
}

func NewSeeSupplierDebtBiz(
	repo SeeSupplierDebtRepo,
	requester middleware.Requester) *seeSupplierDebtBiz {
	return &seeSupplierDebtBiz{
		repo:      repo,
		requester: requester,
	}
}

func (biz *seeSupplierDebtBiz) SeeSupplierDebt(
	ctx context.Context,
	supplierId string,
	filterSupplierDebt *filter.SupplierDebtFilter,
	paging *common.Paging) (*suppliermodel.ResDebtSupplier, error) {
	if !biz.requester.IsHasFeature(constants.SupplierViewFeatureCode) {
		return nil, suppliermodel.ErrSupplierViewNoPermission
	}

	supplier, err := biz.repo.SeeSupplierDebt(
		ctx, supplierId, filterSupplierDebt, paging)
	if err != nil {
		return nil, err
	}

	return supplier, nil
}
