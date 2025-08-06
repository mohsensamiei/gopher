package gormext

import (
	"gorm.io/gorm"
)

func NewAutoPreload() gorm.Plugin {
	return new(AutoPreload)
}

type AutoPreload struct {
}

func (p AutoPreload) Name() string {
	return "gorm:auto_preload"
}

func (p AutoPreload) Initialize(db *gorm.DB) error {
	return db.Callback().Query().Before("gorm:query").Register(p.Name(), p.function)
}

func (p AutoPreload) function(db *gorm.DB) {
	if db.Statement.Schema == nil {
		return
	}
	for _, field := range db.Statement.Schema.Fields {
		if _, ok := field.TagSettings["AUTO_PRELOAD"]; !ok {
			continue
		}
		*db = *db.Preload(field.Name, func(db *gorm.DB) *gorm.DB {
			return db.Unscoped()
		})
	}
}
