package categorymodel

import "qtk-store-api/common"

type ResListCategory struct {
	Data   []Category    `json:"data"`
	Paging common.Paging `json:"paging"`
	Filter Filter        `json:"filter"`
}
