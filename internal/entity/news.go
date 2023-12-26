package entity

import (
	"context"
)

type INewsService interface {
	EditNews(ctx context.Context, pointerNews *PointerNews) (*News, error)
	GetNewsList(ctx context.Context, list, limit int) ([]News, error)
}

type INewsRepository interface {
	EditNews(ctx context.Context, pointerNews *PointerNews) (*News, error)
	GetNewsList(ctx context.Context, list, limit int) ([]News, error)
}

type News struct {
	Id         int    `json:"Id" reform:"Id,pk"`
	Title      string `json:"Title" reform:"Title"`
	Context    string `json:"Context" reform:"Context"`
	Categories []int  `json:"Categories"`
}

type PointerNews struct {
	Id         *int    `json:"Id" reform:"Id,pk"`
	Title      *string `json:"Title" reform:"Title"`
	Context    *string `json:"Context" reform:"Context"`
	Categories []int   `json:"Categories" reform:"-"`
}
