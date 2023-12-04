package rolefeaturestore

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/module/rolefeature/rolefeaturemodel"
)

func (s *sqlStore) CreateListImportNoteDetail(
	ctx context.Context,
	data []rolefeaturemodel.RoleFeature) error {
	db := s.db

	if err := db.CreateInBatches(data, len(data)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
