package query

import (
	"net/url"
	"regexp"
)

const (
	includeKey = "include"
)

var (
	includeTermExp   = `[a-zA-Z][a-zA-Z0-9_.]*`
	includeTermRegex = regexp.MustCompile(includeTermExp)
)

type IncludeClauses []string

func (c *IncludeClauses) UnmarshalQuery(values url.Values) {
	for _, raw := range values[includeKey] {
		*c = append(*c, includeTermRegex.FindAllString(raw, -1)...)
	}
}

func (c IncludeClauses) MarshalQuery(values *url.Values) {
	if len(c) <= 0 {
		return
	}
	for _, f := range c {
		values.Add(includeKey, f)
	}
}

func Include(includes ...string) *Query {
	return new(Query).Include(includes...)
}

func (q *Query) Include(includes ...string) *Query {
	for _, include := range includes {
		q.IncludeClauses = append(q.IncludeClauses, include)
	}
	return q
}
