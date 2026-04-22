package utils

import "net/http"

type pagination struct {
	Page       int64 `json:"page"`
	Limit      int64 `json:"limit"`
	TotalItems int64 `json:"totalItems"`
	TotalPages int64 `json:"totalPages"`
	HasNext    bool  `json:"hasNext"`
	HasPrev    bool  `json:"hasPrev"`
}

type paginatedData struct {
	Data       any        `json:"data"`
	Pagination pagination `json:"pagination"`
}

func SendPaginatedData(w http.ResponseWriter, data any, count, page, limit int64) {
	var totalPages int64
	if limit > 0 {
		totalPages = (count + limit - 1) / limit // ceiling division
	}
	
	hasNext := page < totalPages
	hasPrev := page > 1 && totalPages > 0

	p_data := paginatedData{
		Data: data,
		Pagination: pagination{
			Page:       page,
			Limit:      limit,
			TotalItems: count,
			TotalPages: totalPages,
			HasNext:    hasNext,
			HasPrev:    hasPrev,
		},
	}
	SendJSONData(w, p_data, 200)
}
