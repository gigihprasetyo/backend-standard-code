package paging

import (
	"github.com/gigihprasetyo/backend-standard-code/internal/core/domain"
	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
)

var DefaultPage int = 1
var DefaultLimit int = 10

func Resolver(q domain.PagingQuery) *paginator.Paginator {
	p := paginator.New()

	p.SetKeys("ID") // [default: "ID"] (supporting multiple keys, order of keys matters)

	if q.After != nil {
		p.SetAfterCursor(*q.After) // [default: nil]
	}

	if q.Before != nil {
		p.SetBeforeCursor(*q.Before) // [default: nil]
	}

	if q.Limit != nil {
		p.SetLimit(*q.Limit) // [default: 10]
	}

	if q.Order != nil && *q.Order == "asc" {
		p.SetOrder(paginator.ASC) // [default: paginator.DESC]
	}
	return p
}
