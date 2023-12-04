package importnotemodel

import (
	"qtk-store-api/common"
)

type ResSeeDetailImportNote struct {
	// Data contains the detailed information about import note details.
	Data ResDetailImportNote `json:"data"`
	// Paging provides information about pagination.
	Paging common.Paging `json:"paging,omitempty"`
}
