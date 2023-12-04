package productstore

import (
	"context"
	"errors"
	"qtk-store-api/common"
	"qtk-store-api/module/product/productmodel"

	"gorm.io/gorm"
)

func (s *sqlStore) FindProduct(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*productmodel.Product, error) {
	var data productmodel.Product
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
