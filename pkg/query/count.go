package query

import "fmt"

const (
	countKey = "count"
)

func (q Query) CountOnly() bool {
	val := q.get(countKey)
	if len(val) == 0 {
		return false
	}
	return val[0] == "true"
}

func Count(flag bool) Query {
	return make(Query).Count(flag)
}

func (q Query) Count(flag bool) Query {
	q.set(countKey, fmt.Sprint(flag))
	return q
}
