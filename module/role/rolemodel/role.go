package rolemodel

import (
	"errors"
	"qtk-store-api/common"
	"qtk-store-api/constants"
	"qtk-store-api/module/rolefeature/rolefeaturemodel"
)

type Role struct {
	Id           string           `json:"id" gorm:"column:id;"`
	Name         string           `json:"name" gorm:"column:name;"`
	RoleFeatures ListRoleFeatures `json:"features"`
}

func (*Role) TableName() string {
	return constants.TblRole
}

type ListRoleFeatures []rolefeaturemodel.RoleFeature

func (*ListRoleFeatures) TableName() string {
	return constants.TblRoleFeature
}

var (
	ErrRoleNameEmpty = common.NewCustomError(
		errors.New("name of role is empty"),
		"name of role is empty",
		"ErrRoleNameEmpty",
	)
	ErrRoleFeaturesEmpty = common.NewCustomError(
		errors.New("features of role is empty"),
		"features of role is empty",
		"ErrRoleFeaturesEmpty",
	)
	ErrRoleFeatureInvalid = common.NewCustomError(
		errors.New("features of role is invalid"),
		"features of role is invalid",
		"ErrRoleFeatureInvalid",
	)
	ErrRoleCreateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to create role"),
	)
	ErrRoleUpdateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to update role"),
	)
	ErrRoleViewNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to view role"),
	)
)
