package query

import (
	"net/url"
)

const (
	searchKey = "search"
)

type SearchClause string

func (c *SearchClause) UnmarshalQuery(values url.Values) {
	*c = SearchClause(values.Get(searchKey))
}

func (c SearchClause) MarshalQuery(values *url.Values) {
	if c == "" {
		return
	}
	values.Add(searchKey, string(c))
}

func Search(query string) *Query {
	return new(Query).Search(query)
}

func (q *Query) Search(query string) *Query {
	q.SearchClause = SearchClause(query)
	return q
}
