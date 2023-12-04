package authormodel

import "qtk-store-api/common"

type ResListAuthor struct {
	Data   []Author      `json:"data"`
	Paging common.Paging `json:"paging"`
	Filter Filter        `json:"filter"`
}
