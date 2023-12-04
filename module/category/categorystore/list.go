package categorystore

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/constants"
	"qtk-store-api/module/category/categorymodel"

	"gorm.io/gorm"
)

func (s *sqlStore) ListCategory(ctx context.Context, filter *categorymodel.Filter, propertiesContainSearchKey []string, paging *common.Paging) ([]categorymodel.Category, error) {
	var result []categorymodel.Category
	db := s.db

	db = db.Table(constants.TblCategory)

	handleFilter(db, filter, propertiesContainSearchKey)

	dbTemp, errPaging := common.HandlePaging(db, paging)
	if errPaging != nil {
		return nil, errPaging
	}
	db = dbTemp

	if err := db.
		Limit(int(paging.Limit)).
		Order("name").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}

func handleFilter(
	db *gorm.DB,
	filter *categorymodel.Filter,
	propertiesContainSearchKey []string) {
	if filter != nil {
		if filter.SearchKey != "" {
			db = common.GetWhereClause(db, filter.SearchKey, propertiesContainSearchKey)
		}
	}
}
