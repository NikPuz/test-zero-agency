package service

import (
	"context"
	"test-zero-agency/internal/app/config"
	"test-zero-agency/internal/entity"
)

type NewsService struct {
	NewsRepo entity.INewsRepository
	Cfg      *config.Config
}

func NewNewsService(cfg *config.Config, newsRepo entity.INewsRepository) entity.INewsService {
	return &NewsService{
		NewsRepo: newsRepo,
		Cfg:      cfg,
	}
}

func (s NewsService) EditNews(ctx context.Context, pointerNews *entity.PointerNews) (*entity.News, error) {

	news, err := s.NewsRepo.EditNews(ctx, pointerNews)
	if err != nil {
		return nil, err
	}

	return news, nil
}

func (s NewsService) GetNewsList(ctx context.Context, list, limit int) ([]entity.News, error) {

	news, err := s.NewsRepo.GetNewsList(ctx, list, limit)
	if err != nil {
		return nil, err
	}

	return news, nil
}
