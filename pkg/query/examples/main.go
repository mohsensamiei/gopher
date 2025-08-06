package main

import (
	"net/url"

	"github.com/mohsensamiei/gopher/v3/pkg/query"
	log "github.com/sirupsen/logrus"
)

func main() {
	Filter()
	Include()
	Page()
	Sort()
	Search()
	Count()
}

func Filter() {
	q := query.
		Filter("verified_at", query.NotNull).
		Filter("id", query.Equal, 123).
		Filter("status", query.In, "PAID", "CHECKOUT").
		Filter("name", query.Like, "text")
	log.Print(url.QueryUnescape(q.String()))

	pq, _ := query.Parse(q.String())
	for _, item := range pq.FilterClauses {
		log.Print(item.Field, " ", item.Function, " ", item.Values)
	}
}

func Include() {
	q := query.
		Include("a", "b").
		Include("c")
	log.Print(url.QueryUnescape(q.String()))

	pq, _ := query.Parse(q.String())
	for _, item := range pq.IncludeClauses {
		log.Print(item)
	}
}

func Page() {
	q := query.
		Skip(-1).
		Take(1000)
	log.Print(url.QueryUnescape(q.String()))

	pq, _ := query.Parse(q.String())
	log.Print(pq.SkipClause)
	log.Print(pq.TakeClause)
}

func Sort() {
	q := query.
		Sort("created_at", query.ASC).
		Sort("status", query.DESC)
	log.Print(url.QueryUnescape(q.String()))

	pq, _ := query.Parse(q.String())
	for _, item := range pq.SortClauses {
		log.Print(item.Field, " ", item.Function)
	}
}

func Search() {
	q := query.
		Search("text").
		Search("text1")
	log.Print(url.QueryUnescape(q.String()))

	pq, _ := query.Parse(q.String())
	log.Print(pq.SearchClause)
}

func Count() {
	q := query.
		Count(false).
		Count(true)
	log.Print(url.QueryUnescape(q.String()))

	pq, _ := query.Parse(q.String())
	log.Print(pq.CountClause)
}
