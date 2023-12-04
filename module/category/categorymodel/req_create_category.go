package categorymodel

import "qtk-store-api/constants"

type ReqCreateCategory struct {
	Name string `json:"name" json:"column:name;" example:"Trinh thám"`
}

func (*ReqCreateCategory) TableName() string {
	return constants.TblCategory
}
