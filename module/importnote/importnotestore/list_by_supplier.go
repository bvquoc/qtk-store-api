package importnotestore

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/constants"
	"qtk-store-api/module/importnote/importnotemodel"
	"qtk-store-api/module/supplier/suppliermodel/filter"
	"time"

	"gorm.io/gorm"
)

func (s *sqlStore) ListAllImportNoteBySupplier(
	supplierId string,
	filter *filter.SupplierImportFilter,
	ctx context.Context,
	paging *common.Paging) ([]importnotemodel.ImportNote, error) {
	var result []importnotemodel.ImportNote
	db := s.db

	db = db.Table(constants.TblImportNote)

	handleSupplierImportFilter(db, filter)

	dbTemp, errPaging := common.HandlePaging(db, paging)
	if errPaging != nil {
		return nil, errPaging
	}
	db = dbTemp

	if err := db.
		Where("supplierId = ?", supplierId).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

func handleSupplierImportFilter(
	db *gorm.DB,
	filterSupplierDebt *filter.SupplierImportFilter) {
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
