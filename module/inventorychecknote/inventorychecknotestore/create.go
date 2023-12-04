package inventorychecknotestore

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/module/inventorychecknote/inventorychecknotemodel"
)

func (s *sqlStore) CreateInventoryCheckNote(
	ctx context.Context,
	data *inventorychecknotemodel.ReqCreateInventoryCheckNote) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
