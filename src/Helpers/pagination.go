package Helpers

import (
	"key-rate-api/src/client"
	"math"
)

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
	totalPages := int(math.Ceil(totalPagesFloat))

	result := make([]client.KeyRates, 0)
	startIndex := p.CurrentPage * p.PerPage
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
