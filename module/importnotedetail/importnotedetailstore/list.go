package importnotedetailstore

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/constants"
	"qtk-store-api/module/importnotedetail/importnotedetailmodel"
)

func (s *sqlStore) ListImportNoteDetail(
	ctx context.Context,
	importNoteId string,
	paging *common.Paging) ([]importnotedetailmodel.ImportNoteDetail, error) {
	var result []importnotedetailmodel.ImportNoteDetail
	db := s.db

	db = db.Table(constants.TblImportNoteDetail)

	db = db.Where("importNoteId = ?", importNoteId)

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
