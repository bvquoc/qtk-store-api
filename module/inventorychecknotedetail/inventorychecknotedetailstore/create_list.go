package inventorychecknotedetailstore

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/module/inventorychecknotedetail/inventorychecknotedetailmodel"
)

func (s *sqlStore) CreateListInventoryCheckNoteDetail(
	ctx context.Context,
	data []inventorychecknotedetailmodel.InventoryCheckNoteDetailCreate) error {
	db := s.db

	if err := db.CreateInBatches(data, len(data)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
