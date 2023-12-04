package usermodel

import "qtk-store-api/constants"

type ReqLoginUser struct {
	Email    string `json:"email" gorm:"column:email;" example:"b@gmail.com"`
	Password string `json:"password" gorm:"-" example:"app123"`
}

func (*ReqLoginUser) TableName() string {
	return constants.TblUser
}
