package suppliermodel

import (
	"qtk-store-api/common"
	"qtk-store-api/module/supplier/suppliermodel/filter"
)

type ResSeeImportNoteSupplier struct {
	// Data contains the detailed information about supplier's import notes.
	Data ResImportNoteSupplier `json:"data"`
	// Paging provides information about pagination.
	Paging common.Paging `json:"paging,omitempty"`
	// Filter contains the filter parameters used to retrieve import notes.
	Filter filter.SupplierImportFilter `json:"filter,omitempty"`
}
