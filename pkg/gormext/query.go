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
		query.Null:                     "%v IS NULL",
		query.NotNull:                  "%v IS NOT NULL",
		query.Equal:                    "%v = ?",
		query.NotEqual:                 "%v <> ?",
		query.In:                       "%v IN (?)",
		query.NotIn:                    "%v NOT IN (?)",
		query.Contains:                 "? IN %v",
		query.NotContains:              "? NOT IN %v",
		query.LessThan:                 "%v < ?",
		query.LessThanOrEqual:          "%v <= ?",
		query.GreaterThan:              "%v > ?",
		query.GreaterThanOrEqual:       "%v >= ?",
		query.LessThanOrNull:           "(%v < ? OR %v IS NULL)",
		query.LessThanOrEqualOrNull:    "(%v <= ? OR %v IS NULL)",
		query.GreaterThanOrNull:        "(%v > ? OR %v IS NULL)",
		query.GreaterThanOrEqualOrNull: "(%v >= ? OR %v IS NULL)",
		query.Like:                     "%v LIKE ?",
		query.NotLike:                  "%v NOT LIKE ?",
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
		field := toDelimited(filter.Field, strcase.ToSnake)
		switch filter.Function {
		case query.Null, query.NotNull:
			db = db.Where(fmt.Sprintf(queries[filter.Function], field))
		case query.LessThanOrNull, query.LessThanOrEqualOrNull, query.GreaterThanOrNull, query.GreaterThanOrEqualOrNull:
			db = db.Where(fmt.Sprintf(queries[filter.Function], field, field), filter.Values)
		case query.Like, query.NotLike:
			var values []any
			for _, value := range filter.Values {
				values = append(values, fmt.Sprint("%", value, "%"))
			}
			db = db.Where(fmt.Sprintf(queries[filter.Function], field), values)
		default:
			db = db.Where(fmt.Sprintf(queries[filter.Function], field), filter.Values)
		}
	}
	return db
}

func toDelimited(s string, f func(s string) string) string {
	var res []string
	for _, v := range strings.Split(s, ".") {
		res = append(res, f(v))
	}
	return strings.Join(res, ".")
}
