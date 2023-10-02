package entity

import (
	"github.com/Muruyung/go-utilities/logger"
	"time"
)

// ExampleName example name entity
type ExampleName struct {
	id        int
	name      string
	status    int
	isActive  bool
	testID    int64
	startDate time.Time
	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

// DTOExampleName dto for example name entity
type DTOExampleName struct {
	ID        int
	Name      string
	Status    int
	IsActive  bool
	TestID    int64
	StartDate time.Time
}

// NewExampleName build new entity example name
func NewExampleName(dto DTOExampleName) (*ExampleName, error) {
	exampleName := &ExampleName{
		id:        dto.ID,
		name:      dto.Name,
		status:    dto.Status,
		isActive:  dto.IsActive,
		testID:    dto.TestID,
		startDate: dto.StartDate,
	}

	err := exampleName.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return exampleName, nil
}

func (data *ExampleName) validate() error {
	return nil
}

// GetID get id value
func (data *ExampleName) GetID() int {
	return data.id
}

// SetID set id value
func (data *ExampleName) SetID(id int) (*ExampleName, error) {
	data.id = id
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetName get name value
func (data *ExampleName) GetName() string {
	return data.name
}

// SetName set name value
func (data *ExampleName) SetName(name string) (*ExampleName, error) {
	data.name = name
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetStatus get status value
func (data *ExampleName) GetStatus() int {
	return data.status
}

// SetStatus set status value
func (data *ExampleName) SetStatus(status int) (*ExampleName, error) {
	data.status = status
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetIsActive get isActive value
func (data *ExampleName) GetIsActive() bool {
	return data.isActive
}

// SetIsActive set isActive value
func (data *ExampleName) SetIsActive(isActive bool) (*ExampleName, error) {
	data.isActive = isActive
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetTestID get testID value
func (data *ExampleName) GetTestID() int64 {
	return data.testID
}

// SetTestID set testID value
func (data *ExampleName) SetTestID(testID int64) (*ExampleName, error) {
	data.testID = testID
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetStartDate get startDate value
func (data *ExampleName) GetStartDate() time.Time {
	return data.startDate
}

// SetStartDate set startDate value
func (data *ExampleName) SetStartDate(startDate time.Time) (*ExampleName, error) {
	data.startDate = startDate
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetCreatedAt get createdAt value
func (data *ExampleName) GetCreatedAt() time.Time {
	return data.createdAt
}

// SetCreatedAt set createdAt value
func (data *ExampleName) SetCreatedAt(date time.Time) (*ExampleName, error) {
	data.createdAt = date
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetUpdatedAt get updatedAt value
func (data *ExampleName) GetUpdatedAt() time.Time {
	return data.updatedAt
}

// SetUpdatedAt set updatedAt value
func (data *ExampleName) SetUpdatedAt(date time.Time) (*ExampleName, error) {
	data.updatedAt = date
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetDeletedAt get deletedAt value
func (data *ExampleName) GetDeletedAt() *time.Time {
	return data.deletedAt
}

// SetDeletedAt set deletedAt value
func (data *ExampleName) SetDeletedAt(date *time.Time) (*ExampleName, error) {
	data.deletedAt = date
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}
