package userstore

import (
	"context"
	"errors"
	"qtk-store-api/common"
	"qtk-store-api/constants"
	"qtk-store-api/module/user/usermodel"

	"gorm.io/gorm"
)

func (s *sqlStore) FindUser(
	ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string) (*usermodel.User, error) {
	db := s.db.Table(constants.TblUser)

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user usermodel.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrRecordNotFound()
		}

		return nil, common.ErrDB(err)
	}

	return &user, nil
}
