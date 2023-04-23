package gormext

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/pinosell/gopher/pkg/query"
	"gorm.io/gorm"
	"strings"
)

var (
	queries = map[query.FilterFunction]string{
		query.Is:                 "%v IS ?",
		query.IsNot:              "%v IS NOT ?",
		query.Null:               "%v IS NULL",
		query.NotNull:            "%v IS NOT NULL",
		query.Equal:              "%v = ?",
		query.NotEqual:           "%v <> ?",
		query.In:                 "%v IN (?)",
		query.NotIn:              "%v NOT IN (?)",
		query.Contains:           "? IN %v",
		query.NotContains:        "? NOT IN %v",
		query.LessThan:           "%v < ?",
		query.LessThanOrEqual:    "%v <= ?",
		query.GreaterThan:        "%v > ?",
		query.GreaterThanOrEqual: "%v >= ?",
		query.Like:               "%v LIKE ?",
		query.NotLike:            "%v NOT LIKE ?",
	}
	orders = map[query.SortFunction]string{
		query.ASC:  "%v ASC",
		query.DESC: "%v DESC",
	}
)

func ApplyQuery[T any](db *gorm.DB, q query.Query) *gorm.DB {
	{
		if q.CountOnly() {
			db = db.Limit(0)
			return db
		}
	}
	{
		db = db.Limit(q.TakeCount()).Offset(q.SkipCount())
	}
	{
		for _, include := range q.IncludeItems() {
			db = db.Preload(toDelimited(include, strcase.ToCamel))
		}
	}
	{
		db = applyFilter(db, q.FilterClauses())
	}
	{
		db = applySearch[T](db, q.SearchQuery())
	}
	{
		for _, sort := range q.SortClauses() {
			db = db.Order(fmt.Sprintf(orders[sort.Function], sort.Field))
		}
	}
	return db
}

func ApplyCount[T any](db *gorm.DB, q query.Query) *gorm.DB {
	{
		db = applyFilter(db, q.FilterClauses())
	}
	{
		db = applySearch[T](db, q.SearchQuery())
	}
	return db
}

func applySearch[T any](db *gorm.DB, query string) *gorm.DB {
	if query == "" {
		return db
	}
	obj, ok := any(new(T)).(Search)
	if !ok {
		return db
	}
	db = db.Where(
		fmt.Sprintf("%v LIKE ?", toDelimited(obj.FullTextName(), strcase.ToSnake)),
		fmt.Sprint("%", query, "%"),
	)
	return db
}

func applyFilter(db *gorm.DB, filters []*query.FilterClause) *gorm.DB {
	for _, filter := range filters {
		exp, values, err := convertExpression(filter)
		if err != nil {
			continue
		}
		db = db.Where(exp, values)
	}
	return db
}

func convertExpression(filter *query.FilterClause) (string, []any, error) {
	switch filter.Function {
	case query.Null, query.NotNull:
		return fmt.Sprintf(queries[filter.Function], toDelimited(filter.Field, strcase.ToSnake)), nil, nil
	case query.Or:
		var (
			exps   []string
			values []any
		)
		for _, value := range filter.Values {
			q, err := query.Parse(value.(string))
			if err != nil {
				return "", nil, err
			}
			for _, clause := range q.FilterClauses() {
				var (
					e string
					v []any
				)
				e, v, err = convertExpression(clause)
				if err != nil {
					return "", nil, err
				}
				exps = append(exps, e)
				values = append(values, v)
			}
		}
		return strings.Join(exps, " OR "), values, nil
	case query.Like, query.NotLike:
		var values []any
		for _, value := range filter.Values {
			values = append(values, fmt.Sprint("%", value, "%"))
		}
		return fmt.Sprintf(queries[filter.Function], toDelimited(filter.Field, strcase.ToSnake)), values, nil
	default:
		return fmt.Sprintf(queries[filter.Function], toDelimited(filter.Field, strcase.ToSnake)), filter.Values, nil
	}
}

func toDelimited(s string, f func(s string) string) string {
	var res []string
	for _, v := range strings.Split(s, ".") {
		res = append(res, f(v))
	}
	return strings.Join(res, ".")
}
