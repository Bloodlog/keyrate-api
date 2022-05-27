package paginator

import (
	"key-rate-api/internal/pkg/keyrate/client"
	"math"
)

type Pager interface {
	Paginate(p *Pages) Pagineted
}

type Pages struct {
	Items       []client.KeyRates
	Total       int
	PerPage     int
	CurrentPage int
}

type Pagineted struct {
	Data        []client.KeyRates
	Total       int
	CurrentPage int
	PerPage     int
	TotalPages  int
}

func Paginate(p *Pages) Pagineted {
	total := p.Total
	totalPagesFloat := float64(p.Total / p.PerPage)
	totalPages := max(int(math.Ceil(totalPagesFloat)), 1)
	p.CurrentPage = max(p.CurrentPage, 1)
	if p.CurrentPage > totalPages {
		p.CurrentPage = totalPages
	}

	result := make([]client.KeyRates, 0)
	startIndex := (p.CurrentPage - 1) * p.PerPage
	endIndex := startIndex + p.PerPage

	for i := startIndex; i < endIndex; i++ {
		if p.CurrentPage <= totalPages && i <= p.Total-1 {
			result = append(result, p.Items[i])
		}
	}

	return Pagineted{
		Data:        result,
		Total:       total,
		CurrentPage: p.CurrentPage,
		PerPage:     p.PerPage,
		TotalPages:  totalPages,
	}
}

func max(x int, y int) int {
	if x < y {
		return y
	}

	return x
}
