package importnotestore

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/module/importnote/importnotemodel"
)

func (s *sqlStore) UpdateImportNote(
	ctx context.Context,
	id string,
	data *importnotemodel.ReqUpdateImportNote) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
