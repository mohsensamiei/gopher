package gormext

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/mohsensamiei/gopher/v2/pkg/logic"
	"github.com/mohsensamiei/gopher/v2/pkg/query"
	"github.com/mohsensamiei/gopher/v2/pkg/strcaseext"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

var (
	queries = map[query.FilterFunction]string{
		query.Null:                     "%v IS NULL",
		query.NotNull:                  "%v IS NOT NULL",
		query.Equal:                    "%v = ?",
		query.EqualOrNull:              "(%v = ? OR %v IS NULL)",
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
		query.Like:                     "%v ILIKE ?",
		query.NotLike:                  "%v NOT ILIKE ?",
	}
	orders = map[query.SortFunction]string{
		query.ASC:  "%v ASC",
		query.DESC: "%v DESC",
	}
)

func ApplyQuery[T Model](db *gorm.DB, q *query.Query) *gorm.DB {
	normalize[T](db, q)
	{
		if q.CountClause {
			db = db.Limit(0)
			return db
		}
	}
	{
		db = db.Limit(int(q.TakeClause)).Offset(int(q.SkipClause))
	}
	{
		db = applyFilter(db, q.FilterClauses)
	}
	{
		db = applySearch[T](db, q.SearchClause)
	}
	{
		for _, sort := range q.SortClauses {
			db = db.Order(fmt.Sprintf(orders[sort.Function], sort.Field))
		}
	}
	{
		table := TableName(db, new(T))

		for _, include := range q.IncludeClauses {
			l := strcaseext.Delimited(include, ".", strcase.ToCamel)

			var v reflect.Value
			for i, f := range strings.Split(l, ".") {
				if i == 0 {
					v = reflect.ValueOf(new(T)).Elem()
				} else if v.Kind() != reflect.Struct {
					t := v.Type()
					if t.Kind() == reflect.Slice {
						t = t.Elem()
					}
					if t.Kind() == reflect.Pointer {
						t = t.Elem()
					}
					v = reflect.New(t).Elem()
				}
				v = v.FieldByName(f)
			}

			if v.Kind() == reflect.Slice {
				db = db.Preload(l)
			} else {
				db = db.Preload(l, func(db *gorm.DB) *gorm.DB {
					return db.Unscoped()
				})
			}
		}

		for _, c := range q.FilterClauses {
			d := strings.Split(c.Field, ".")
			if len(d) > 1 && d[0] != table {
				db = db.InnerJoins(strcaseext.Delimited(strings.Join(d[:len(d)-1], "."), ".", strcase.ToCamel))
			}
		}

		fields := []string{fmt.Sprintf("%v.*", table)}
		for _, c := range q.SortClauses {
			d := strings.Split(c.Field, ".")
			if len(d) > 1 && d[0] != table {
				fields = append(fields, c.Field)
				db = db.InnerJoins(strcaseext.Delimited(strings.Join(d[:len(d)-1], "."), ".", strcase.ToCamel))
			}
		}
		db = db.Select(fields).Distinct()
	}
	return db
}

func ApplyCount[T Model](db *gorm.DB, q *query.Query) *gorm.DB {
	normalize[T](db, q)
	{
		db = applyFilter(db, q.FilterClauses)
	}
	{
		db = applySearch[T](db, q.SearchClause)
	}
	{
		model := new(T)
		table := TableName(db, new(T))
		primaryKeys := any(model).(Model).PrimaryKeys()

		for _, c := range q.FilterClauses {
			d := strings.Split(c.Field, ".")
			if len(d) > 1 && d[0] != table {
				db = db.InnerJoins(strcaseext.Delimited(strings.Join(d[:len(d)-1], "."), ".", strcase.ToCamel))
			}
		}
		db = db.Distinct(logic.IFor(func(v string) string {
			return fmt.Sprintf("%v.%v", table, v)
		}, primaryKeys...))
	}
	return db
}

func applySearch[T any](db *gorm.DB, c query.SearchClause) *gorm.DB {
	if c == "" {
		return db
	}
	obj, ok := any(new(T)).(Search)
	if !ok {
		return db
	}
	f := normalizeField[T](db, obj.FullTextName())
	if strings.Contains(f, ".") {
		d := strings.Split(f, ".")
		db = db.InnerJoins(strcaseext.Delimited(strings.Join(d[:len(d)-1], "."), ".", strcase.ToCamel))
	}
	db = db.Where(
		fmt.Sprintf("%v ILIKE ?", f),
		fmt.Sprint("%", c, "%"),
	)
	return db
}

func applyFilter(db *gorm.DB, c query.FilterClauses) *gorm.DB {
	for _, filter := range c {
		switch filter.Function {
		case query.Null, query.NotNull:
			db = db.Where(fmt.Sprintf(queries[filter.Function], filter.Field))
		case query.EqualOrNull, query.LessThanOrNull, query.LessThanOrEqualOrNull, query.GreaterThanOrNull, query.GreaterThanOrEqualOrNull:
			db = db.Where(fmt.Sprintf(queries[filter.Function], filter.Field, filter.Field), filter.Values)
		case query.Like, query.NotLike:
			var values []any
			for _, value := range filter.Values {
				values = append(values, fmt.Sprint("%", value, "%"))
			}
			db = db.Where(fmt.Sprintf(queries[filter.Function], filter.Field), values)
		default:
			db = db.Where(fmt.Sprintf(queries[filter.Function], filter.Field), filter.Values)
		}
	}
	return db
}

func normalize[T any](db *gorm.DB, q *query.Query) {
	for _, c := range q.FilterClauses {
		c.Field = normalizeField[T](db, c.Field)
	}
	for _, c := range q.SortClauses {
		c.Field = normalizeField[T](db, c.Field)
	}
}

func normalizeField[T any](db *gorm.DB, field string) string {
	if strings.Contains(field, ".") {
		dump := strings.Split(field, ".")
		nested := strcaseext.Delimited(strings.Join(dump[0:len(dump)-1], "."), ".", strcase.ToCamel)
		return fmt.Sprintf(`"%v"."%v"`, strings.Join(strings.Split(nested, "."), `"."`), strcase.ToSnake(dump[len(dump)-1]))
	} else {
		return fmt.Sprintf(`"%v"."%v"`, TableName(db, new(T)), strcase.ToSnake(field))
	}
}
