package controllers

import (
	"context"

	"github.com/MayCMF/example/schema"
)

// IExample - Example business logic interface
type IExample interface {
	// Query data
	Query(ctx context.Context, params schema.ExampleQueryParam, opts ...schema.ExampleQueryOptions) (*schema.ExampleQueryResult, error)
	// Get specified data
	Get(ctx context.Context, UUID string, opts ...schema.ExampleQueryOptions) (*schema.Example, error)
	// Create data
	Create(ctx context.Context, item schema.Example) (*schema.Example, error)
	// Update data
	Update(ctx context.Context, UUID string, item schema.Example) (*schema.Example, error)
	// Delete data
	Delete(ctx context.Context, UUID string) error
	// Update status
	UpdateStatus(ctx context.Context, UUID string, status int) error
}
