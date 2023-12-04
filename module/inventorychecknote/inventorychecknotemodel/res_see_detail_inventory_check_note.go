package inventorychecknotemodel

import (
	"qtk-store-api/common"
)

type ResSeeDetailInventoryCheckNote struct {
	// Data contains the detailed information about inventory check note details.
	Data ResDetailInventoryCheckNote `json:"data"`
	// Paging provides information about pagination.
	Paging common.Paging `json:"paging,omitempty"`
}
