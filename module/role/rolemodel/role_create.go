package rolemodel

import (
	"qtk-store-api/common"
	"qtk-store-api/constants"
)

type RoleCreate struct {
	Id       string   `json:"-" gorm:"column:id;"`
	Name     string   `json:"name" gorm:"column:name;"`
	Features []string `json:"features" gorm:"-"`
}

func (*RoleCreate) TableName() string {
	return constants.TblRole
}

func (data *RoleCreate) Validate() *common.AppError {
	if common.ValidateEmptyString(data.Name) {
		return ErrRoleNameEmpty
	}
	if data.Features == nil || len(data.Features) == 0 {
		return ErrRoleFeaturesEmpty
	}
	for _, v := range data.Features {
		if !common.ValidateNotNilId(&v) {
			return ErrRoleFeatureInvalid
		}
	}
	return nil
}
