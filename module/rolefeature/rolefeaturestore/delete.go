package rolefeaturestore

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/constants"
	"qtk-store-api/module/role/rolemodel"
)

func (s *sqlStore) DeleteRoleFeature(
	ctx context.Context,
	conditions map[string]interface{}) error {
	db := s.db

	if err := db.
		Table(constants.TblRoleFeature).
		Where(conditions).
		Delete(&rolemodel.Role{}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
