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

func (s NewsService) EditNews(ctx context.Context, news *entity.News) (*entity.News, error) {

	repoNews, err := s.NewsRepo.GetNews(ctx, news.Id)
	if err != nil {
		return nil, err
	}

	if len(news.Title) == 0 {
		news.Title = repoNews.Title
	}
	if len(news.Context) == 0 {
		news.Context = repoNews.Context
	}

	news, err = s.NewsRepo.EditNews(ctx, news)
	if err != nil {
		return nil, err
	}

	if news.Categories != nil {
		err = s.NewsRepo.DeleteNewsCategory(ctx, news.Id)
		if err != nil {
			return nil, err
		}

		for _, categoryId := range news.Categories {
			err = s.NewsRepo.InsertNewsCategory(ctx, &entity.NewsCategories{
				NewsId:     news.Id,
				CategoryId: categoryId,
			})
			if err != nil {
				return nil, err
			}
		}
	} else {
		news.Categories = repoNews.Categories
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
