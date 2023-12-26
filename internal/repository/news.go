package repository

import (
	"context"
	"gopkg.in/reform.v1"
	"test-zero-agency/internal/entity"
)

type NewsRepository struct {
	db *reform.DB
}

func NewNewsRepository(db *reform.DB) entity.INewsRepository {
	return &NewsRepository{db: db}
}

func (r NewsRepository) EditNews(ctx context.Context, pointerNews *entity.PointerNews) (*entity.News, error) {
	return nil, nil
}

func (r NewsRepository) GetNewsList(ctx context.Context, list, limit int) ([]entity.News, error) {
	return nil, nil
}
