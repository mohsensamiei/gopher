package query

import (
	"fmt"
	"net/url"
	"strings"
)

const (
	countKey = "count"
)

type CountClause bool

func (c *CountClause) UnmarshalQuery(values url.Values) {
	*c = strings.ToLower(strings.TrimSpace(values.Get(countKey))) == "true"
}

func (c CountClause) MarshalQuery(values *url.Values) {
	if c == false {
		return
	}
	values.Add(countKey, fmt.Sprint(c))
}

func Count(flag bool) *Query {
	return new(Query).Count(flag)
}

func (q *Query) Count(flag bool) *Query {
	q.CountClause = CountClause(flag)
	return q
}
