package models

import (
	"sort"
	"time"

	"github.com/Muruyung/go-utilities/converter"
)

// ExampleNameModels example name models struct
type ExampleNameModels struct {
	ID        int        `dbq:"id" json:"id"`
	Name      string     `dbq:"name" json:"name"`
	Status    int        `dbq:"status" json:"status"`
	IsActive  bool       `dbq:"is_active" json:"is_active"`
	TestID    int64      `dbq:"test_id" json:"test_id"`
	StartDate time.Time  `dbq:"start_date" json:"start_date"`
	CreatedAt time.Time  `dbq:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time  `dbq:"updated_at,omitempty" json:"updated_at,omitempty"`
	DeletedAt *time.Time `dbq:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

// GetTableName get table name of example name models
func (models ExampleNameModels) GetTableName() string {
	return "example_name"
}

// GetModels get models of example name models
func (models ExampleNameModels) GetModels() interface{} {
	return models
}

// GetModelsMap get models map of example name models
func (models ExampleNameModels) GetModelsMap() map[string]interface{} {
	dataMap, _ := converter.ConvertInterfaceToMap(models)
	return dataMap
}

// GetColumns get columns of example name models
func (models ExampleNameModels) GetColumns() []string {
	var (
		modelsMap = models.GetModelsMap()
		arrColumn = make([]string, 0)
	)

	for column := range modelsMap {
		arrColumn = append(arrColumn, column)
	}
	sort.Strings(arrColumn)

	return arrColumn
}

// GetValStruct get value struct of example name models
func (models ExampleNameModels) GetValStruct(arrColumn []string) []interface{} {
	var (
		modelsMap = models.GetModelsMap()
		arrValue  = make([]interface{}, 0)
	)

	for _, column := range arrColumn {
		arrValue = append(arrValue, modelsMap[column])
	}

	return []interface{}{arrValue}
}
