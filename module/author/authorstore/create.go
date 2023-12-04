package authorstore

import (
	"context"
	"qtk-store-api/common"
	"qtk-store-api/module/author/authormodel"
)

func (s *sqlStore) CreateAuthor(ctx context.Context, data *authormodel.Author) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		if gormErr := common.GetGormErr(err); gormErr != nil {
			switch key := gormErr.GetDuplicateErrorKey("PRIMARY"); key {
			case "PRIMARY":
				return authormodel.ErrAuthorIdDuplicate
			}
		}
		return common.ErrDB(err)
	}

	return nil
}
