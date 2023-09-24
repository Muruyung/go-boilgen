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

func (strc *ExampleName) validate() error {
	return nil
}

// GetID get id value
func (strc *ExampleName) GetID() int {
	return strc.id
}

// SetID set id value
func (strc *ExampleName) SetID(id int) (*ExampleName, error) {
	strc.id = id
	err := strc.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return strc, nil
}

// GetName get name value
func (strc *ExampleName) GetName() string {
	return strc.name
}

// SetName set name value
func (strc *ExampleName) SetName(name string) (*ExampleName, error) {
	strc.name = name
	err := strc.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return strc, nil
}

// GetStatus get status value
func (strc *ExampleName) GetStatus() int {
	return strc.status
}

// SetStatus set status value
func (strc *ExampleName) SetStatus(status int) (*ExampleName, error) {
	strc.status = status
	err := strc.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return strc, nil
}

// GetIsActive get isActive value
func (strc *ExampleName) GetIsActive() bool {
	return strc.isActive
}

// SetIsActive set isActive value
func (strc *ExampleName) SetIsActive(isActive bool) (*ExampleName, error) {
	strc.isActive = isActive
	err := strc.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return strc, nil
}

// GetTestID get testID value
func (strc *ExampleName) GetTestID() int64 {
	return strc.testID
}

// SetTestID set testID value
func (strc *ExampleName) SetTestID(testID int64) (*ExampleName, error) {
	strc.testID = testID
	err := strc.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return strc, nil
}

// GetStartDate get startDate value
func (strc *ExampleName) GetStartDate() time.Time {
	return strc.startDate
}

// SetStartDate set startDate value
func (strc *ExampleName) SetStartDate(startDate time.Time) (*ExampleName, error) {
	strc.startDate = startDate
	err := strc.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return strc, nil
}

// GetCreatedAt get createdAt value
func (strc *ExampleName) GetCreatedAt() time.Time {
	return strc.createdAt
}

// GetUpdatedAt get updatedAt value
func (strc *ExampleName) GetUpdatedAt() time.Time {
	return strc.updatedAt
}

// GetDeletedAt get deletedAt value
func (strc *ExampleName) GetDeletedAt() *time.Time {
	return strc.deletedAt
}
