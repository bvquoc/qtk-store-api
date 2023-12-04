package authormodel

import "qtk-store-api/constants"

type ReqCreateAuthor struct {
	Name string `json:"name" json:"column:name;" example:"Nguyễn Nhật Ánh"`
}

func (*ReqCreateAuthor) TableName() string {
	return constants.TblAuthor
}
