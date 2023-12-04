package rolestore

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/module/role/rolemodel"
)

func (s *sqlStore) CreateRole(
	ctx context.Context,
	data *rolemodel.RoleCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
