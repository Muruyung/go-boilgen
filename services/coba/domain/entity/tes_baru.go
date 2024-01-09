package entity

import (
	"github.com/Muruyung/go-boilgen/pkg/logger"
	"time"
)

// TesBaru tes baru entity
type TesBaru struct {
	id        string
	apacing   int
	test      bool
	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

// DTOTesBaru dto for tes baru entity
type DTOTesBaru struct {
	ID      string
	Apacing int
	Test    bool
}

// NewTesBaru build new entity tes baru
func NewTesBaru(dto DTOTesBaru) (*TesBaru, error) {
	tesBaru := &TesBaru{
		id:      dto.ID,
		apacing: dto.Apacing,
		test:    dto.Test,
	}

	err := tesBaru.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return tesBaru, nil
}

func (data *TesBaru) validate() error {
	return nil
}

// GetID get id value
func (data *TesBaru) GetID() string {
	return data.id
}

// SetID set id value
func (data *TesBaru) SetID(id string) (*TesBaru, error) {
	data.id = id
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetApacing get apacing value
func (data *TesBaru) GetApacing() int {
	return data.apacing
}

// SetApacing set apacing value
func (data *TesBaru) SetApacing(apacing int) (*TesBaru, error) {
	data.apacing = apacing
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetTest get test value
func (data *TesBaru) GetTest() bool {
	return data.test
}

// SetTest set test value
func (data *TesBaru) SetTest(test bool) (*TesBaru, error) {
	data.test = test
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetCreatedAt get createdAt value
func (data *TesBaru) GetCreatedAt() time.Time {
	return data.createdAt
}

// SetCreatedAt set createdAt value
func (data *TesBaru) SetCreatedAt(date time.Time) (*TesBaru, error) {
	data.createdAt = date
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetUpdatedAt get updatedAt value
func (data *TesBaru) GetUpdatedAt() time.Time {
	return data.updatedAt
}

// SetUpdatedAt set updatedAt value
func (data *TesBaru) SetUpdatedAt(date time.Time) (*TesBaru, error) {
	data.updatedAt = date
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetDeletedAt get deletedAt value
func (data *TesBaru) GetDeletedAt() *time.Time {
	return data.deletedAt
}

// SetDeletedAt set deletedAt value
func (data *TesBaru) SetDeletedAt(date *time.Time) (*TesBaru, error) {
	data.deletedAt = date
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}
