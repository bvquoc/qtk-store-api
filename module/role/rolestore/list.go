package rolestore

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/constants"
	"qtk-store-api/module/role/rolemodel"
)

func (s *sqlStore) ListRole(
	ctx context.Context) ([]rolemodel.Role, error) {
	var result []rolemodel.Role
	db := s.db

	db = db.Table(constants.TblRole)

	if err := db.
		Preload("RoleFeatures").
		Find(&result).
		Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
