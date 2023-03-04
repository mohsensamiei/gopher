package query

import (
	"regexp"
)

const (
	includeKey = "include"
)

var (
	includeTermExp   = `[a-zA-Z][a-zA-Z0-9_.]*`
	includeTermRegex = regexp.MustCompile(includeTermExp)
)

func (q Query) IncludeItems() []string {
	var includes []string
	for _, raw := range q.get(includeKey) {
		includes = append(includes, includeTermRegex.FindAllString(raw, -1)...)
	}
	return includes
}

func Include(includes ...string) Query {
	return make(Query).Include(includes...)
}

func (q Query) Include(includes ...string) Query {
	for _, include := range includes {
		q.add(includeKey, include)
	}
	return q
}
