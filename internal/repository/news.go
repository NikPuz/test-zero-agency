package repository

import (
	"context"
	"database/sql"
	"errors"
	"gopkg.in/reform.v1"
	"test-zero-agency/internal/entity"
)

type NewsRepository struct {
	db *reform.DB
}

func NewNewsRepository(db *reform.DB) entity.INewsRepository {
	return &NewsRepository{db: db}
}

func (r NewsRepository) EditNews(ctx context.Context, news *entity.News) (*entity.News, error) {

	err := r.db.WithContext(ctx).Update(news)
	if err != nil {
		return nil, err
	}

	return news, nil
}

func (r NewsRepository) GetNewsList(ctx context.Context, list, limit int) ([]entity.News, error) {

	rows, err := r.db.QueryContext(ctx, `SELECT n.id, n.title, n.content, nc.categoryid FROM 
                                                   (SELECT id, title, content FROM news LIMIT $1 OFFSET $2) AS n
                                                       LEFT JOIN newscategories AS nc ON nc.newsid = n.id;`, limit, limit*(list-1))
	if err != nil {
		return nil, err
	}

	newsMap := make(map[int]entity.News)
	for rows.Next() {
		news := entity.News{Categories: []int{}}
		var categoryId *int
		rows.Scan(
			&news.Id,
			&news.Title,
			&news.Context,
			&categoryId,
		)

		if _, ok := newsMap[news.Id]; !ok {
			newsMap[news.Id] = news
		}

		if categoryId != nil {
			n := newsMap[news.Id]
			n.Categories = append(n.Categories, *categoryId)
			newsMap[news.Id] = n
		}
	}

	var news []entity.News
	for _, v := range newsMap {
		news = append(news, v)
	}

	return news, nil
}

func (r NewsRepository) GetNews(ctx context.Context, id int) (*entity.News, error) {

	rows, err := r.db.QueryContext(ctx, `SELECT id, title, content, nc.categoryid FROM news
                                                       LEFT JOIN newscategories AS nc ON nc.newsid = id WHERE id = $1;`, id)
	if err != nil {
		return nil, err
	}

	news := entity.News{Categories: []int{}}
	for rows.Next() {
		var categoryId *int
		rows.Scan(
			&news.Id,
			&news.Title,
			&news.Context,
			&categoryId,
		)

		if categoryId != nil {
			news.Categories = append(news.Categories, *categoryId)
		}
	}

	return &news, nil
}

func (r NewsRepository) DeleteNewsCategory(ctx context.Context, newsId int) error {

	_, err := r.db.ExecContext(ctx, `DELETE FROM newscategories WHERE newsid = $1;`, newsId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	return nil
}

func (r NewsRepository) InsertNewsCategory(ctx context.Context, newsCategories *entity.NewsCategories) error {

	err := r.db.WithContext(ctx).Insert(newsCategories)
	if err != nil {
		return err
	}

	return nil
}
