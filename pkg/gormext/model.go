package gormext

import (
	"github.com/google/uuid"
)

type Model interface {
	PrimaryKeys() []string
}

type IncrementalModel struct {
	ID uint32 `gorm:"primaryKey"`
}

func (m IncrementalModel) PrimaryKeys() []string {
	return []string{"id"}
}

type UniversalModel struct {
	ID uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4()"`
}

func (m UniversalModel) PrimaryKeys() []string {
	return []string{"id"}
}
