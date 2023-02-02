package gormext

import (
	"github.com/google/uuid"
)

type IncrementalModel struct {
	ID uint32
}

func (m IncrementalModel) GetID() uint32 {
	return m.ID
}

type UniversalModel struct {
	ID uuid.UUID `gorm:"default:uuid_generate_v4()"`
}

func (m UniversalModel) GetID() uuid.UUID {
	return m.ID
}
