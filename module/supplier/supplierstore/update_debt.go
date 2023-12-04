package supplierstore

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/constants"
	"qtk-store-api/module/supplier/suppliermodel"

	"gorm.io/gorm"
)

func (s *sqlStore) UpdateSupplierDebt(
	ctx context.Context,
	id string,
	data *suppliermodel.ReqUpdateDebtSupplier) error {
	db := s.db

	if err := db.Table(constants.TblSupplier).
		Where("id = ?", id).
		Update("debt", gorm.Expr("debt + ?", data.QuantityUpdate)).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
