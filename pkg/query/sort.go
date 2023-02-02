package query

import (
	"fmt"
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

type SortClause struct {
	Field    string
	Function SortFunction
}

func (c SortClause) String() string {
	return fmt.Sprintf("%v(%v)", c.Function, c.Field)
}

var (
	sortTermExp = fmt.Sprintf(`(%v)\(([^\)]*)\)`,
		strings.Join([]string{
			string(ASC),
			string(DESC),
		}, "|"))
	sortTermRegex = regexp.MustCompile(sortTermExp)
)

func (q Query) SortClauses() []*SortClause {
	var clauses []*SortClause
	for _, sort := range q.get(sortKey) {
		for _, raw := range sortTermRegex.FindAllStringSubmatch(sort, -1) {
			clauses = append(clauses, &SortClause{
				Field:    raw[2],
				Function: SortFunction(raw[1]),
			})
		}
	}
	return clauses
}

func Sort(field string, function SortFunction) Query {
	return make(Query).Sort(field, function)
}

func (q Query) Sort(field string, function SortFunction) Query {
	clause := SortClause{
		Field:    field,
		Function: function,
	}
	q.add(sortKey, clause.String())
	return q
}
