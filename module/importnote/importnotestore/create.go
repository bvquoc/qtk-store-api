package importnotestore

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/module/importnote/importnotemodel"
)

func (s *sqlStore) CreateImportNote(
	ctx context.Context,
	data *importnotemodel.ReqCreateImportNote) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
