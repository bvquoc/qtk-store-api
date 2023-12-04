package importnotedetailstore

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/module/importnotedetail/importnotedetailmodel"
)

func (s *sqlStore) CreateListImportNoteDetail(
	ctx context.Context,
	data []importnotedetailmodel.ImportNoteDetailCreate) error {
	db := s.db

	if err := db.CreateInBatches(data, len(data)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
