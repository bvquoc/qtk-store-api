package importnotedetailstore

import (
	"context"
	"errors"
	"qtk-store-api/common"
	"qtk-store-api/constants"
	"qtk-store-api/module/importnotedetail/importnotedetailmodel"

	"gorm.io/gorm"
)

func (s *sqlStore) FindListImportNoteDetail(ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) ([]importnotedetailmodel.ImportNoteDetail, error) {
	var data []importnotedetailmodel.ImportNoteDetail
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.
		Table(constants.TblImportNoteDetail).
		Where(conditions).
		Find(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound()
		}
		return nil, common.ErrDB(err)
	}

	return data, nil
}
