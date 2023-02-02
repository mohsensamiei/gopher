package query

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	skipKey = "skip"
	takeKey = "take"
)

func (q Query) SkipCount() int {
	val := q.get(skipKey)
	if len(val) == 0 {
		return 0
	}
	i, _ := strconv.Atoi(strings.TrimSpace(val[0]))
	if i < 0 {
		return 0
	}
	return i
}

func (q Query) TakeCount() int {
	val := q.get(takeKey)
	if len(val) == 0 {
		return 10
	}
	i, _ := strconv.Atoi(strings.TrimSpace(val[0]))
	if i < 1 {
		return 10
	} else if i > 100 {
		return 100
	}
	return i
}

func Skip(count int) Query {
	return make(Query).Skip(count)
}

func (q Query) Skip(count int) Query {
	q.set(skipKey, fmt.Sprint(count))
	return q
}

func Take(count int) Query {
	return make(Query).Take(count)
}

func (q Query) Take(count int) Query {
	q.set(takeKey, fmt.Sprint(count))
	return q
}
