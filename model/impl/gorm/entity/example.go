package entity

import (
	"context"

	"github.com/MayCMF/core/src/common/entity"
	"github.com/MayCMF/example/schema"
	"github.com/jinzhu/gorm"
)

// GetExampleDB - Get the Example store
func GetExampleDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return entity.GetDBWithModel(ctx, defDB, Example{})
}

// SchemaExample - Example object
type SchemaExample schema.Example

// ToExample - Convert to Example entity
func (a SchemaExample) ToExample() *Example {
	item := &Example{
		UUID:    a.UUID,
		Code:    &a.Code,
		Name:    &a.Name,
		Memo:    &a.Memo,
		Status:  &a.Status,
		Creator: &a.Creator,
	}
	return item
}

// Example - Example entity
type Example struct {
	entity.Model
	UUID    string  `gorm:"column:uuid;size:36;index;"`  // UUID code
	Code    *string `gorm:"column:code;size:50;index;"`  // Number
	Name    *string `gorm:"column:name;size:100;index;"` // Name
	Memo    *string `gorm:"column:memo;size:200;"`       // Remarks
	Status  *int    `gorm:"column:status;index;"`        // Status (1: Enable 2: Disable)
	Creator *string `gorm:"column:creator;size:36;"`     // Creator
}

func (a Example) String() string {
	return entity.ToString(a)
}

// TableName - Table Name
func (a Example) TableName() string {
	return a.Model.TableName("example")
}

// ToSchemaExample - Convert to Example object
func (a Example) ToSchemaExample() *schema.Example {
	item := &schema.Example{
		UUID:      a.UUID,
		Code:      *a.Code,
		Name:      *a.Name,
		Memo:      *a.Memo,
		Status:    *a.Status,
		Creator:   *a.Creator,
		CreatedAt: a.CreatedAt,
	}
	return item
}

// Examples - Example list
type Examples []*Example

// ToSchemaExamples - Convert to Example object list
func (a Examples) ToSchemaExamples() []*schema.Example {
	list := make([]*schema.Example, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaExample()
	}
	return list
}
