package inventorychecknotedetailstore

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/constants"
	"qtk-store-api/module/inventorychecknotedetail/inventorychecknotedetailmodel"
)

func (s *sqlStore) ListInventoryCheckNoteDetail(
	ctx context.Context,
	inventoryCheckNoteId string,
	paging *common.Paging) ([]inventorychecknotedetailmodel.InventoryCheckNoteDetail, error) {
	var result []inventorychecknotedetailmodel.InventoryCheckNoteDetail
	db := s.db

	db = db.Table(constants.TblInventoryCheckNoteDetail)

	db = db.Where("inventoryCheckNoteId = ?", inventoryCheckNoteId)

	dbTemp, errPaging := common.HandlePaging(db, paging)
	if errPaging != nil {
		return nil, errPaging
	}
	db = dbTemp

	if err := db.
		Preload("Product").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
