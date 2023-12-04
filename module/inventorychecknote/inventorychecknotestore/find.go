package inventorychecknotestore

import (
	"context"
	"errors"
	"qtk-store-api/common"
	"qtk-store-api/module/inventorychecknote/inventorychecknotemodel"

	"gorm.io/gorm"
)

func (s *sqlStore) FindInventoryCheckNote(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*inventorychecknotemodel.InventoryCheckNote, error) {
	var data inventorychecknotemodel.InventoryCheckNote
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(conditions).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound()
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
