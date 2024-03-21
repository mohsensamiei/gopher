package query

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

const (
	sortKey = "sort"
)

type SortFunction string

const (
	ASC  SortFunction = "asc"
	DESC SortFunction = "desc"
)

var (
	sortTermExp = fmt.Sprintf(`(%v)\(([^\)]*)\)`,
		strings.Join([]string{
			string(ASC),
			string(DESC),
		}, "|"))
	sortTermRegex = regexp.MustCompile(sortTermExp)
)

type SortClause struct {
	Field    string
	Function SortFunction
}

type SortClauses []SortClause

func (c *SortClauses) UnmarshalQuery(values url.Values) {
	for _, sort := range values[sortKey] {
		for _, raw := range sortTermRegex.FindAllStringSubmatch(sort, -1) {
			*c = append(*c, SortClause{
				Field:    raw[2],
				Function: SortFunction(raw[1]),
			})
		}
	}
}

func (c SortClauses) MarshalQuery(values *url.Values) {
	if len(c) <= 0 {
		return
	}
	for _, f := range c {
		values.Add(sortKey, fmt.Sprintf("%v(%v)", f.Function, f.Field))
	}
}

func Sort(field string, function SortFunction) *Query {
	return new(Query).Sort(field, function)
}

func (q *Query) Sort(field string, function SortFunction) *Query {
	q.SortClauses = append(q.SortClauses, SortClause{
		Field:    field,
		Function: function,
	})
	return q
}
