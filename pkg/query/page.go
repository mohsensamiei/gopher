package query

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

const (
	skipKey = "skip"
	takeKey = "take"
)

type SkipClause int

func (c *SkipClause) UnmarshalQuery(values url.Values) {
	*c = 0
	val := values[skipKey]
	if len(val) == 0 {
		return
	}
	i, _ := strconv.Atoi(strings.TrimSpace(val[0]))
	if i < 0 {
		return
	}
	*c = SkipClause(i)
}

func (c SkipClause) MarshalQuery(values *url.Values) {
	if c <= 0 {
		return
	}
	values.Add(skipKey, fmt.Sprint(c))
}

type TakeClause int

func (c *TakeClause) UnmarshalQuery(values url.Values) {
	*c = 10
	val := values[takeKey]
	if len(val) == 0 {
		return
	}
	i, _ := strconv.ParseInt(strings.TrimSpace(val[0]), 10, 64)
	if i < 1 {
		return
	}
	*c = TakeClause(i)
}

func (c TakeClause) MarshalQuery(values *url.Values) {
	if c <= 0 {
		return
	}
	values.Add(takeKey, fmt.Sprint(c))
}

func Skip(count int) *Query {
	return new(Query).Skip(count)
}

func (q *Query) Skip(count int) *Query {
	q.SkipClause = SkipClause(count)
	return q
}

func Take(count int) *Query {
	return new(Query).Take(count)
}

func (q *Query) Take(count int) *Query {
	q.TakeClause = TakeClause(count)
	return q
}
