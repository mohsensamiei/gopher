package gormext

import (
	"context"
	"github.com/mohsensamiei/gopher/v2/pkg/di"
	"gorm.io/gorm"
)

func Transaction(ctx context.Context, f func(ctx context.Context) error) (err error) {
	db := di.Provide[*gorm.DB](ctx, Name)
	tx := db.Begin().WithContext(ctx)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()
	sub := context.WithValue(ctx, Name, tx)
	if err = f(sub); err != nil {
		return err
	}
	return tx.Commit().Error
}
