package gormext

import (
	"context"
	"github.com/mohsensamiei/gopher/v3/pkg/di"
	"gorm.io/gorm"
)

func Transaction(ctx context.Context, f func(ctx context.Context) error) (err error) {
	tx := di.Provide[*gorm.DB](ctx).Begin().WithContext(ctx)
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()
	if err = f(di.Register[*gorm.DB](ctx, tx)); err != nil {
		return err
	}
	return tx.Commit().Error
}
