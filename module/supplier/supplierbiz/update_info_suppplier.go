package supplierbiz

import (
	"context"
	"qtk-store-api/constants"
	"qtk-store-api/middleware"
	"qtk-store-api/module/supplier/suppliermodel"
)

type UpdateInfoSupplierRepo interface {
	UpdateSupplierInfo(
		ctx context.Context,
		supplierId string,
		data *suppliermodel.ReqUpdateInfoSupplier,
	) error
}

type updateInfoSupplierBiz struct {
	repo      UpdateInfoSupplierRepo
	requester middleware.Requester
}

func NewUpdateInfoSupplierBiz(
	repo UpdateInfoSupplierRepo,
	requester middleware.Requester) *updateInfoSupplierBiz {
	return &updateInfoSupplierBiz{repo: repo, requester: requester}
}

func (biz *updateInfoSupplierBiz) UpdateInfoSupplier(
	ctx context.Context,
	id string,
	data *suppliermodel.ReqUpdateInfoSupplier) error {
	if !biz.requester.IsHasFeature(constants.SupplierUpdateInfoFeatureCode) {
		return suppliermodel.ErrSupplierUpdateInfoNoPermission
	}

	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.repo.UpdateSupplierInfo(ctx, id, data); err != nil {
		return err
	}

	return nil
}
