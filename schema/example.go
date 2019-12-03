package schema

import (
	"github.com/MayCMF/core/src/common/schema"
	"time"
)

// Example - Example object
type Example struct {
	UUID      string    `json:"uuid"`                                  // Record ID
	Code      string    `json:"code" binding:"required"`               // Code
	Name      string    `json:"name" binding:"required"`               // Name
	Memo      string    `json:"memo"`                                  // Remarks
	Status    int       `json:"status" binding:"required,max=2,min=1"` // Status (1: Enable 2: Disable)
	Creator   string    `json:"creator"`                               // Creator
	CreatedAt time.Time `json:"created_at"`                            // Creation time
}

// ExampleQueryParam - Query conditions
type ExampleQueryParam struct {
	Code     string // CODE
	Status   int    // Status (1: Enable 2: Disable)
	LikeCode string // Number (fuzzy query)
	LikeName string // Name (fuzzy query)
}

// ExampleQueryOptions - Example object query optional parameter item
type ExampleQueryOptions struct {
	PageParam *schema.PaginationParam // Paging parameter
}

// ExampleQueryResult - Example object query result
type ExampleQueryResult struct {
	Data       []*Example
	PageResult *schema.PaginationResult
}
