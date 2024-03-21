package query

import (
	"net/url"
)

func New() *Query {
	return new(Query)
}

type Query struct {
	TakeClause     TakeClause     `json:"take"`    // Count of return records, empty equal 10
	SkipClause     SkipClause     `json:"skip"`    // Count of skip records, empty equal 0
	IncludeClauses IncludeClauses `json:"include"` // Loads embedded fields
	FilterClauses  FilterClauses  `json:"filter"`  // Filters records, such as 'id:eq("5200efe3-3842-4c3e-929d-0a05b3bda793")'
	SortClauses    SortClauses    `json:"sort"`    // Sorts records. such as desc(arranged_at)
	SearchClause   SearchClause   `json:"search"`  // full text search between records
	CountClause    CountClause    `json:"count"`   // Count only
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
