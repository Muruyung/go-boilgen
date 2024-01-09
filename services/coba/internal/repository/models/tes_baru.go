package models

import (
	"sort"
	"time"

	"github.com/Muruyung/go-utilities/converter"
)

// TesBaruModels tes baru models struct
type TesBaruModels struct {
	ID        string     `dbq:"id" json:"id"`
	Apacing   int        `dbq:"apacing" json:"apacing"`
	Test      bool       `dbq:"test" json:"test"`
	CreatedAt time.Time  `dbq:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time  `dbq:"updated_at,omitempty" json:"updated_at,omitempty"`
	DeletedAt *time.Time `dbq:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

// GetTableName get table name of tes baru models
func (models TesBaruModels) GetTableName() string {
	return "tes_baru"
}

// GetModels get models of tes baru models
func (models TesBaruModels) GetModels() interface{} {
	return models
}

// GetModelsMap get models map of tes baru models
func (models TesBaruModels) GetModelsMap() map[string]interface{} {
	dataMap, _ := converter.ConvertInterfaceToMap(models)
	return dataMap
}

// GetColumns get columns of tes baru models
func (models TesBaruModels) GetColumns() []string {
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

// GetValStruct get value struct of tes baru models
func (models TesBaruModels) GetValStruct(arrColumn []string) []interface{} {
	var (
		modelsMap = models.GetModelsMap()
		arrValue  = make([]interface{}, 0)
	)

	for _, column := range arrColumn {
		arrValue = append(arrValue, modelsMap[column])
	}

	return []interface{}{arrValue}
}
