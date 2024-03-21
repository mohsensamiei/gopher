package query

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mohsensamiei/gopher/v2/pkg/stringsext"
	"net/url"
	"regexp"
	"strings"
	"time"
)

const (
	filterKey = "filter"
)

type FilterFunction string

const (
	Null                     FilterFunction = "null"
	NotNull                  FilterFunction = "nnull"
	Equal                    FilterFunction = "eq"
	EqualOrNull              FilterFunction = "eqn"
	NotEqual                 FilterFunction = "neq"
	GreaterThan              FilterFunction = "gt"
	GreaterThanOrEqual       FilterFunction = "gte"
	LessThan                 FilterFunction = "lt"
	LessThanOrEqual          FilterFunction = "lte"
	GreaterThanOrNull        FilterFunction = "gtn"
	GreaterThanOrEqualOrNull FilterFunction = "gten"
	LessThanOrNull           FilterFunction = "ltn"
	LessThanOrEqualOrNull    FilterFunction = "lten"
	In                       FilterFunction = "in"
	NotIn                    FilterFunction = "nin"
	Contains                 FilterFunction = "cnt"
	NotContains              FilterFunction = "ncnt"
	Like                     FilterFunction = "like"
	NotLike                  FilterFunction = "nlike"
)

var (
	filterValueExp = `\d+|\"[^\"]*\"|true|false`
	filterTermExp  = fmt.Sprintf(`(%v):(%v)(\((,|\s+|%v)*\))`,
		`[a-zA-Z][a-zA-Z0-9_.]*`,
		strings.Join([]string{
			string(Like),
			string(NotLike),
			string(Null),
			string(NotNull),
			string(Equal),
			string(EqualOrNull),
			string(NotEqual),
			string(GreaterThan),
			string(GreaterThanOrEqual),
			string(LessThan),
			string(LessThanOrEqual),
			string(GreaterThanOrNull),
			string(GreaterThanOrEqualOrNull),
			string(LessThanOrNull),
			string(LessThanOrEqualOrNull),
			string(In),
			string(NotIn),
			string(Contains),
			string(NotContains),
		}, "|"),
		filterValueExp,
	)
	filterTermRegex  = regexp.MustCompile(filterTermExp)
	filterValueRegex = regexp.MustCompile(filterValueExp)
)

type FilterClause struct {
	Field    string
	Function FilterFunction
	Values   []any
}

type FilterClauses []FilterClause

func (c *FilterClauses) UnmarshalQuery(values url.Values) {
	for _, filter := range values[filterKey] {
		for _, raw := range filterTermRegex.FindAllStringSubmatch(filter, -1) {
			*c = append(*c, FilterClause{
				Field:    raw[1],
				Function: FilterFunction(raw[2]),
				Values:   toInterfaceSlice(filterValueRegex.FindAllString(raw[3], -1)),
			})
		}
	}
}

func (c FilterClauses) MarshalQuery(values *url.Values) {
	if len(c) <= 0 {
		return
	}
	for _, f := range c {
		values.Add(filterKey, fmt.Sprintf("%v:%v(%v)",
			f.Field,
			f.Function,
			strings.Join(toStringSlice(f.Values), ","),
		))
	}
}

func Filter(field string, function FilterFunction, values ...any) *Query {
	return new(Query).Filter(field, function, values...)
}

func (q *Query) Filter(field string, function FilterFunction, values ...any) *Query {
	q.FilterClauses = append(q.FilterClauses, FilterClause{
		Field:    field,
		Function: function,
		Values:   values,
	})
	return q
}

func toStringSlice(values []any) []string {
	var res []string
	for _, value := range values {
		switch v := value.(type) {
		case time.Time:
			res = append(res, fmt.Sprintf("%q", v.UTC().Format(time.RFC3339Nano)))
		case uuid.UUID, string, stringsext.Stringify:
			res = append(res, fmt.Sprintf("%q", v))
		default:
			res = append(res, fmt.Sprint(v))
		}
	}
	return res
}

func toInterfaceSlice(values []string) []any {
	var res []any
	for _, v := range values {
		res = append(res, strings.Trim(v, "\""))
	}
	return res
}
