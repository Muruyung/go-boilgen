package entity

import (
	"github.com/Muruyung/go-boilgen/pkg/logger"
	"time"
)

// Ngetes ngetes entity
type Ngetes struct {
	id        string
	name      string
	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

// DTONgetes dto for ngetes entity
type DTONgetes struct {
	ID   string
	Name string
}

// NewNgetes build new entity ngetes
func NewNgetes(dto DTONgetes) (*Ngetes, error) {
	ngetes := &Ngetes{
		id:   dto.ID,
		name: dto.Name,
	}

	err := ngetes.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return ngetes, nil
}

func (data *Ngetes) validate() error {
	return nil
}

// GetID get id value
func (data *Ngetes) GetID() string {
	return data.id
}

// SetID set id value
func (data *Ngetes) SetID(id string) (*Ngetes, error) {
	data.id = id
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetName get name value
func (data *Ngetes) GetName() string {
	return data.name
}

// SetName set name value
func (data *Ngetes) SetName(name string) (*Ngetes, error) {
	data.name = name
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetCreatedAt get createdAt value
func (data *Ngetes) GetCreatedAt() time.Time {
	return data.createdAt
}

// SetCreatedAt set createdAt value
func (data *Ngetes) SetCreatedAt(date time.Time) (*Ngetes, error) {
	data.createdAt = date
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetUpdatedAt get updatedAt value
func (data *Ngetes) GetUpdatedAt() time.Time {
	return data.updatedAt
}

// SetUpdatedAt set updatedAt value
func (data *Ngetes) SetUpdatedAt(date time.Time) (*Ngetes, error) {
	data.updatedAt = date
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}

// GetDeletedAt get deletedAt value
func (data *Ngetes) GetDeletedAt() *time.Time {
	return data.deletedAt
}

// SetDeletedAt set deletedAt value
func (data *Ngetes) SetDeletedAt(date *time.Time) (*Ngetes, error) {
	data.deletedAt = date
	err := data.validate()
	if err != nil {
		logger.Logger.Error(err)
		return nil, err
	}
	return data, nil
}
