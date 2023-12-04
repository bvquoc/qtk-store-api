package rolefeaturemodel

import (
	"errors"
	"qtk-store-api/common"
	"qtk-store-api/constants"
)

type RoleFeature struct {
	RoleId    string `json:"roleId" gorm:"column:roleId;"`
	FeatureId string `json:"featureId" gorm:"column:featureId;"`
}

func (*RoleFeature) TableName() string {
	return constants.TblRoleFeature
}

var (
	ErrRoleFeatureIdFeatureInvalid = common.NewCustomError(
		errors.New("id of feature is invalid"),
		"id of feature is invalid",
		"ErrRoleFeatureIdFeatureInvalid",
	)
)
