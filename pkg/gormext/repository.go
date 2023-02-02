package gormext

import (
	"context"
	"fmt"
	"github.com/pinosell/gopher/pkg/di"
	"github.com/pinosell/gopher/pkg/errors"
	"github.com/pinosell/gopher/pkg/query"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

type Identity[T any] interface {
	GetID() T
}

type Crud[M Identity[K], K any] interface {
	Create(ctx context.Context, model *M) error
	ReturnByID(ctx context.Context, id K, qs query.Encode) (*M, error)
	Update(ctx context.Context, model *M) error
	Save(ctx context.Context, model *M) error
	List(ctx context.Context, qs query.Encode) ([]*M, int64, error)
	DeleteByID(ctx context.Context, id K) error
	Delete(ctx context.Context, model *M) error
}

func NewCrudRepository[M Identity[K], K any]() *CrudRepository[M, K] {
	return new(CrudRepository[M, K])
}

type CrudRepository[M Identity[K], K any] struct {
}

func (r CrudRepository[M, K]) DB(ctx context.Context) *gorm.DB {
	return di.Provide[*gorm.DB](ctx, Name)
}

func (r CrudRepository[M, K]) ReturnByID(ctx context.Context, id K, qe query.Encode) (*M, error) {
	q, err := qe.Parse()
	if err != nil {
		return nil, err
	}

	model := new(M)
	table := TableName(r.DB(ctx), model)

	if err = ApplyQuery[M](r.DB(ctx), q).
		Where(fmt.Sprintf("%v.id = ?", table), id).
		First(model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.Wrap(err, codes.NotFound).
				WithDetailF("%v with id = '%v'", table, id)
		}
		return nil, err
	}
	return model, nil
}

func (r CrudRepository[M, K]) Create(ctx context.Context, model *M) error {
	return r.DB(ctx).Create(model).Error
}

func (r CrudRepository[M, K]) Update(ctx context.Context, model *M) error {
	return r.DB(ctx).Updates(model).Error
}

func (r CrudRepository[M, K]) Save(ctx context.Context, model *M) error {
	return r.DB(ctx).Save(model).Error
}

func (r CrudRepository[M, K]) List(ctx context.Context, qe query.Encode) ([]*M, int64, error) {
	q, err := qe.Parse()
	if err != nil {
		return nil, 0, err
	}

	var list []*M
	if err := ApplyQuery[M](r.DB(ctx), q).
		Find(&list).Error; err != nil {
		return nil, 0, err
	}

	var count int64
	if err := ApplyCount[M](r.DB(ctx), q).Model(new(M)).
		Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func (r CrudRepository[M, K]) DeleteByID(ctx context.Context, id K) error {
	model, err := r.ReturnByID(ctx, id, query.Empty)
	if err != nil {
		return err
	}

	if err = r.DB(ctx).Delete(model).Error; err != nil {
		return err
	}
	return nil
}

func (r CrudRepository[M, K]) Delete(ctx context.Context, model *M) error {
	if err := r.DB(ctx).Delete(model).Error; err != nil {
		return err
	}
	return nil
}
