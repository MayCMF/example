package model

import (
	"context"

	"github.com/MayCMF/core/src/common/errors"
	"github.com/MayCMF/core/src/common/model"
	"github.com/MayCMF/example/model/impl/gorm/entity"
	"github.com/MayCMF/example/schema"
	"github.com/jinzhu/gorm"
)

// NewExample - Create a Example storage instance
func NewExample(db *gorm.DB) *Example {
	return &Example{db}
}

// Example - Example storage
type Example struct {
	db *gorm.DB
}

func (a *Example) getQueryOption(opts ...schema.ExampleQueryOptions) schema.ExampleQueryOptions {
	var opt schema.ExampleQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query - Query data
func (a *Example) Query(ctx context.Context, params schema.ExampleQueryParam, opts ...schema.ExampleQueryOptions) (*schema.ExampleQueryResult, error) {
	db := entity.GetExampleDB(ctx, a.db)
	if v := params.Code; v != "" {
		db = db.Where("code=?", v)
	}
	if v := params.LikeCode; v != "" {
		db = db.Where("code LIKE ?", "%"+v+"%")
	}
	if v := params.LikeName; v != "" {
		db = db.Where("name LIKE ?", "%"+v+"%")
	}
	if v := params.Status; v > 0 {
		db = db.Where("status=?", v)
	}
	db = db.Order("id DESC")

	opt := a.getQueryOption(opts...)
	var list entity.Examples
	pr, err := model.WrapPageQuery(ctx, db, opt.PageParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	qr := &schema.ExampleQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaExamples(),
	}

	return qr, nil
}

// Get - Query specified data
func (a *Example) Get(ctx context.Context, UUID string, opts ...schema.ExampleQueryOptions) (*schema.Example, error) {
	db := entity.GetExampleDB(ctx, a.db).Where("uuid=?", UUID)
	var item entity.Example
	ok, err := model.FindOne(ctx, db, &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaExample(), nil
}

// Create - Create data
func (a *Example) Create(ctx context.Context, item schema.Example) error {
	example := entity.SchemaExample(item).ToExample()
	result := entity.GetExampleDB(ctx, a.db).Create(example)
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Update - Update data
func (a *Example) Update(ctx context.Context, UUID string, item schema.Example) error {
	example := entity.SchemaExample(item).ToExample()
	result := entity.GetExampleDB(ctx, a.db).Where("uuid=?", UUID).Omit("uuid", "creator").Updates(example)
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Delete - delete data
func (a *Example) Delete(ctx context.Context, UUID string) error {
	result := entity.GetExampleDB(ctx, a.db).Where("uuid=?", UUID).Delete(entity.Example{})
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// UpdateStatus - update status
func (a *Example) UpdateStatus(ctx context.Context, UUID string, status int) error {
	result := entity.GetExampleDB(ctx, a.db).Where("uuid=?", UUID).Update("status", status)
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
