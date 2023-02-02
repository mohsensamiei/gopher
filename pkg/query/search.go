package query

const (
	searchKey = "search"
)

func (q Query) SearchQuery() string {
	val := q.get(searchKey)
	if len(val) == 0 {
		return ""
	}
	return val[0]
}

func Search(query string) Query {
	return make(Query).Search(query)
}

func (q Query) Search(query string) Query {
	q.set(searchKey, query)
	return q
}
