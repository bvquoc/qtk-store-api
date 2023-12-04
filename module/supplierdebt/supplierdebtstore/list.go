package supplierdebtstore

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/constants"
	"qtk-store-api/module/supplier/suppliermodel/filter"
	"qtk-store-api/module/supplierdebt/supplierdebtmodel"
	"time"

	"gorm.io/gorm"
)

func (s *sqlStore) ListSupplierDebt(
	ctx context.Context,
	supplierId string,
	filterSupplierDebt *filter.SupplierDebtFilter,
	paging *common.Paging) ([]supplierdebtmodel.SupplierDebt, error) {
	var result []supplierdebtmodel.SupplierDebt
	db := s.db

	db = db.Table(constants.TblSupplierDebt)

	db = db.Where("supplierId = ?", supplierId)

	handleFilter(db, filterSupplierDebt)

	dbTemp, errPaging := common.HandlePaging(db, paging)
	if errPaging != nil {
		return nil, errPaging
	}
	db = dbTemp

	if err := db.
		Order("createAt desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

func handleFilter(
	db *gorm.DB,
	filterSupplierDebt *filter.SupplierDebtFilter) {
	if filterSupplierDebt != nil {
		if filterSupplierDebt.DateFrom != nil {
			timeFrom := time.Unix(*filterSupplierDebt.DateFrom, 0)
			db = db.Where("createAt >= ?", timeFrom)
		}
		if filterSupplierDebt.DateTo != nil {
			timeTo := time.Unix(*filterSupplierDebt.DateTo, 0)
			db = db.Where("createAt <= ?", timeTo)
		}
	}
}
