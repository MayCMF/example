package implement

import (
	"context"

	"github.com/MayCMF/core/src/common/errors"
	commonschema "github.com/MayCMF/core/src/common/schema"
	"github.com/MayCMF/core/src/common/util"
	"github.com/MayCMF/example/model"
	"github.com/MayCMF/example/schema"
)

// NewExample - Create a Example
func NewExample(mExample model.IExample) *Example {
	return &Example{
		ExampleModel: mExample,
	}
}

// Example - Sample program
type Example struct {
	ExampleModel model.IExample
}

// Query - Query data
func (a *Example) Query(ctx context.Context, params schema.ExampleQueryParam, opts ...schema.ExampleQueryOptions) (*schema.ExampleQueryResult, error) {
	return a.ExampleModel.Query(ctx, params, opts...)
}

// Get - Get specified data
func (a *Example) Get(ctx context.Context, UUID string, opts ...schema.ExampleQueryOptions) (*schema.Example, error) {
	item, err := a.ExampleModel.Get(ctx, UUID, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *Example) checkCode(ctx context.Context, code string) error {
	result, err := a.ExampleModel.Query(ctx, schema.ExampleQueryParam{
		Code: code,
	}, schema.ExampleQueryOptions{
		PageParam: &commonschema.PaginationParam{PageSize: -1},
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return errors.New400Response("Number already exists")
	}
	return nil
}

func (a *Example) getUpdate(ctx context.Context, UUID string) (*schema.Example, error) {
	return a.Get(ctx, UUID)
}

// Create - Create Example data
func (a *Example) Create(ctx context.Context, item schema.Example) (*schema.Example, error) {
	err := a.checkCode(ctx, item.Code)
	if err != nil {
		return nil, err
	}

	item.UUID = util.MustUUID()
	err = a.ExampleModel.Create(ctx, item)
	if err != nil {
		return nil, err
	}
	return a.getUpdate(ctx, item.UUID)
}

// Update - Update Example data
func (a *Example) Update(ctx context.Context, UUID string, item schema.Example) (*schema.Example, error) {
	oldItem, err := a.ExampleModel.Get(ctx, UUID)
	if err != nil {
		return nil, err
	} else if oldItem == nil {
		return nil, errors.ErrNotFound
	} else if oldItem.Code != item.Code {
		err := a.checkCode(ctx, item.Code)
		if err != nil {
			return nil, err
		}
	}

	err = a.ExampleModel.Update(ctx, UUID, item)
	if err != nil {
		return nil, err
	}
	return a.getUpdate(ctx, UUID)
}

// Delete - Delete data
func (a *Example) Delete(ctx context.Context, UUID string) error {
	oldItem, err := a.ExampleModel.Get(ctx, UUID)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.ExampleModel.Delete(ctx, UUID)
}

// UpdateStatus - Update status
func (a *Example) UpdateStatus(ctx context.Context, UUID string, status int) error {
	oldItem, err := a.ExampleModel.Get(ctx, UUID)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.ExampleModel.UpdateStatus(ctx, UUID, status)
}
