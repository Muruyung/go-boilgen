package models

import (
	"sort"
	"time"

	"github.com/Muruyung/go-utilities/converter"
)

// NgetesModels ngetes models struct
type NgetesModels struct {
	ID        string     `dbq:"id" json:"id"`
	Name      string     `dbq:"name" json:"name"`
	CreatedAt time.Time  `dbq:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time  `dbq:"updated_at,omitempty" json:"updated_at,omitempty"`
	DeletedAt *time.Time `dbq:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

// GetTableName get table name of ngetes models
func (models NgetesModels) GetTableName() string {
	return "ngetes"
}

// GetModels get models of ngetes models
func (models NgetesModels) GetModels() interface{} {
	return models
}

// GetModelsMap get models map of ngetes models
func (models NgetesModels) GetModelsMap() map[string]interface{} {
	dataMap, _ := converter.ConvertInterfaceToMap(models)
	return dataMap
}

// GetColumns get columns of ngetes models
func (models NgetesModels) GetColumns() []string {
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

// GetValStruct get value struct of ngetes models
func (models NgetesModels) GetValStruct(arrColumn []string) []interface{} {
	var (
		modelsMap = models.GetModelsMap()
		arrValue  = make([]interface{}, 0)
	)

	for _, column := range arrColumn {
		arrValue = append(arrValue, modelsMap[column])
	}

	return []interface{}{arrValue}
}
