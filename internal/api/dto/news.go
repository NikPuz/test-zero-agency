package dto

import "test-zero-agency/internal/entity"

type NewsList struct {
	Success bool          `json:"Success"`
	News    []entity.News `json:"News"`
}

type NewsGetNewsListLimitParam int
type NewsListParam int
