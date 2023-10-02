package repository

// MapperCommon template for common mapper models
type MapperCommon interface {
	MapDomainToModels() ModelsCommon
	MapModelsToDomain(entityStruct interface{})
}

// ModelsCommon template for common models repository
type ModelsCommon interface {
	GetTableName() string
	GetModels() interface{}
	GetModelsMap() map[string]interface{}
	GetColumns() []string
	GetValStruct(arrColumn []string) []interface{}
}
