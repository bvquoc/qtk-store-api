package suppliermodel

import (
	"qtk-store-api/common"
	"qtk-store-api/module/supplier/suppliermodel/filter"
)

type ResListSupplier struct {
	// Data contains list of supplier.
	Data []Supplier `json:"data"`
	// Paging provides information about pagination.
	Paging common.Paging `json:"paging,omitempty"`
	// Filter contains the filter parameters used to retrieve supplier.
	Filter filter.SupplierImportFilter `json:"filter,omitempty"`
}
