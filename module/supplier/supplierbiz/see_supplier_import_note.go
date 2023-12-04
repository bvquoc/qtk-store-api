package supplierbiz

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/constants"
	"qtk-store-api/middleware"
	"qtk-store-api/module/supplier/suppliermodel"
	"qtk-store-api/module/supplier/suppliermodel/filter"
)

type SeeSupplierImportNoteRepo interface {
	SeeSupplierImportNote(
		ctx context.Context,
		supplierId string,
		filter *filter.SupplierImportFilter,
		paging *common.Paging) (*suppliermodel.ResImportNoteSupplier, error)
}

type seeSupplierImportNoteBiz struct {
	repo      SeeSupplierImportNoteRepo
	requester middleware.Requester
}

func NewSeeSupplierImportNoteBiz(
	repo SeeSupplierImportNoteRepo,
	requester middleware.Requester) *seeSupplierImportNoteBiz {
	return &seeSupplierImportNoteBiz{
		repo:      repo,
		requester: requester,
	}
}

func (biz *seeSupplierImportNoteBiz) SeeSupplierImportNote(
	ctx context.Context,
	supplierId string,
	filter *filter.SupplierImportFilter,
	paging *common.Paging) (*suppliermodel.ResImportNoteSupplier, error) {
	if !biz.requester.IsHasFeature(constants.SupplierViewFeatureCode) {
		return nil, suppliermodel.ErrSupplierViewNoPermission
	}

	supplier, err := biz.repo.SeeSupplierImportNote(
		ctx, supplierId, filter, paging)
	if err != nil {
		return nil, err
	}

	return supplier, nil
}
