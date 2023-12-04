package usermodel

import "qtk-store-api/constants"

type SimpleUser struct {
	Id   string `json:"id" gorm:"column:id;" example:"user id"`
	Name string `json:"name" gorm:"column:name;" example:"Nguyễn Văn A"`
}

func (*SimpleUser) TableName() string {
	return constants.TblUser
}
