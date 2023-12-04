package rolefeaturestore

import (
	"context"
	"errors"
	"qtk-store-api/common"
	"qtk-store-api/constants"
	"qtk-store-api/module/rolefeature/rolefeaturemodel"

	"gorm.io/gorm"
)

func (s *sqlStore) FindListFeatures(
	ctx context.Context,
	roleId string) ([]rolefeaturemodel.RoleFeature, error) {
	var data []rolefeaturemodel.RoleFeature
	db := s.db

	if err := db.
		Table(constants.TblRoleFeature).
		Where("roleId = ?", roleId).
		Find(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound()
		}
		return nil, common.ErrDB(err)
	}

	return data, nil
}
