package gormext

import (
	"context"
	"fmt"
	"github.com/mohsensamiei/gopher/v3/pkg/di"
	"github.com/mohsensamiei/gopher/v3/pkg/errors"
	"github.com/mohsensamiei/gopher/v3/pkg/query"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
	"strings"
)

type Crud[M Model] interface {
	Create(ctx context.Context, model *M) error
	ReturnByPK(ctx context.Context, qs query.Encode, pk ...any) (*M, error)
	Update(ctx context.Context, model *M) error
	Save(ctx context.Context, model *M) error
	List(ctx context.Context, qs query.Encode) ([]*M, int64, error)
	DeleteByPK(ctx context.Context, pk ...any) error
	Delete(ctx context.Context, model *M) error
}

func NewCrudRepository[M Model]() *CrudRepository[M] {
	return new(CrudRepository[M])
}

type CrudRepository[M Model] struct {
}

func (r CrudRepository[M]) ReturnByPK(ctx context.Context, qe query.Encode, pk ...any) (*M, error) {
	q, err := qe.Parse()
	if err != nil {
		return nil, err
	}
	db := di.Provide[*gorm.DB](ctx)

	model := new(M)
	table := TableName(db, model)
	primaryKeys := any(model).(Model).PrimaryKeys()

	var fields []string
	db = ApplyQuery[M](db, q)
	for i, name := range primaryKeys {
		db = db.Where(fmt.Sprintf("%v.%v = ?", table, name), pk[i])
		fields = append(fields, fmt.Sprintf("%v = '%v'", name, pk[i]))
	}

	if err = db.First(model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrap(err, codes.NotFound).
				WithDetailF("%v with %v", table, strings.Join(fields, ", "))
		}
		return nil, err
	}
	return model, nil
}

func (r CrudRepository[M]) Create(ctx context.Context, model *M) error {
	return di.Provide[*gorm.DB](ctx).Create(model).Error
}

func (r CrudRepository[M]) Update(ctx context.Context, model *M) error {
	return di.Provide[*gorm.DB](ctx).Updates(model).Error
}

func (r CrudRepository[M]) Save(ctx context.Context, model *M) error {
	return di.Provide[*gorm.DB](ctx).Save(model).Error
}

func (r CrudRepository[M]) List(ctx context.Context, qe query.Encode) ([]*M, int64, error) {
	q, err := qe.Parse()
	if err != nil {
		return nil, 0, err
	}
	var (
		list  []*M
		count int64
		db    = di.Provide[*gorm.DB](ctx)
	)
	if err = ApplyQuery[M](db, q).
		Find(&list).Error; err != nil {
		return nil, 0, err
	}
	if err = ApplyCount[M](db, q).Model(new(M)).
		Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (r CrudRepository[M]) DeleteByPK(ctx context.Context, pk ...any) error {
	model, err := r.ReturnByPK(ctx, query.Empty, pk...)
	if err != nil {
		return err
	}
	if err = r.Delete(ctx, model); err != nil {
		return err
	}
	return nil
}

func (r CrudRepository[M]) Delete(ctx context.Context, model *M) error {
	if err := di.Provide[*gorm.DB](ctx).Delete(model).Error; err != nil {
		return err
	}
	return nil
}
