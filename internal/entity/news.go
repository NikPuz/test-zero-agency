package entity

import (
	"context"
)

type INewsService interface {
	EditNews(ctx context.Context, news *News) (*News, error)
	GetNewsList(ctx context.Context, list, limit int) ([]News, error)
}

type INewsRepository interface {
	EditNews(ctx context.Context, news *News) (*News, error)
	GetNewsList(ctx context.Context, list, limit int) ([]News, error)
	GetNews(ctx context.Context, id int) (*News, error)
	DeleteNewsCategory(ctx context.Context, newsId int) error
	InsertNewsCategory(ctx context.Context, newsCategories *NewsCategories) error
}

//go:generate reform

type (
	//reform:news
	News struct {
		Id         int    `json:"Id" reform:"id,pk"`
		Title      string `json:"Title" reform:"title"`
		Context    string `json:"Context" reform:"content"`
		Categories []int  `json:"Categories" reform:"-"`
	}
)

//go:generate reform

type (
	//reform:newscategories
	NewsCategories struct {
		NewsId     int `reform:"newsid,pk"`
		CategoryId int `reform:"categoryid"`
	}
)
