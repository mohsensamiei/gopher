package query

import (
	"net/url"
)

func New() *Query {
	return new(Query)
}

type Params struct {
	Take     int      `json:"take"`    // Count of return records, empty equal 10
	Skip     int      `json:"skip"`    // Count of skip records, empty equal 0
	Includes []string `json:"include"` // Loads embedded fields
	Filters  []string `json:"filter"`  // Filters records, such as 'id:eq("5200efe3-3842-4c3e-929d-0a05b3bda793")'
	Sorts    []string `json:"sort"`    // Sorts records. such as desc(arranged_at)
	Search   string   `json:"search"`  // full text search between records
	Count    bool     `json:"count"`   // Count only
}

type Query struct {
	TakeClause     TakeClause
	SkipClause     SkipClause
	IncludeClauses IncludeClauses
	FilterClauses  FilterClauses
	SortClauses    SortClauses
	SearchClause   SearchClause
	CountClause    CountClause
}

func (q Query) String() string {
	values := make(url.Values)
	q.TakeClause.MarshalQuery(&values)
	q.SkipClause.MarshalQuery(&values)
	q.IncludeClauses.MarshalQuery(&values)
	q.FilterClauses.MarshalQuery(&values)
	q.SortClauses.MarshalQuery(&values)
	q.SearchClause.MarshalQuery(&values)
	q.CountClause.MarshalQuery(&values)
	return values.Encode()
}

func Parse(raw string) (*Query, error) {
	values, err := url.ParseQuery(raw)
	if err != nil {
		return nil, err
	}
	q := new(Query)
	q.TakeClause.UnmarshalQuery(values)
	q.SkipClause.UnmarshalQuery(values)
	q.IncludeClauses.UnmarshalQuery(values)
	q.FilterClauses.UnmarshalQuery(values)
	q.SortClauses.UnmarshalQuery(values)
	q.SearchClause.UnmarshalQuery(values)
	q.CountClause.UnmarshalQuery(values)
	return q, nil
}
