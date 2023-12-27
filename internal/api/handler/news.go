package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"go.uber.org/zap"
	"net/http"
	"test-zero-agency/internal/api/dto"
	routerMiddleware "test-zero-agency/internal/api/middleware"
	"test-zero-agency/internal/entity"
	"time"
)

type newsHandler struct {
	NewsService entity.INewsService
	Logger      *zap.Logger
}

func RegisterNewsHandlers(f *fiber.App, service entity.INewsService, logger *zap.Logger, routerMiddleware *routerMiddleware.Middleware) {
	NewsHandler := new(newsHandler)
	NewsHandler.NewsService = service
	NewsHandler.Logger = logger

	f.Route("/news", func(r fiber.Router) {
		r.Use(recover.New())
		r.Use(requestid.New())
		r.Use(routerMiddleware.ContentTypeJSON)
		r.Use(routerMiddleware.DebugLogger)

		r.Post("/edit/:id<int>", timeout.NewWithContext(NewsHandler.EditNews, 10*time.Second))
		r.Get("/list", timeout.NewWithContext(NewsHandler.GetNewsList, 10*time.Second))
	})
}

func (h newsHandler) EditNews(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	news := new(entity.News)
	err = c.BodyParser(news)
	if err != nil {
		return fiber.ErrBadRequest
	}
	news.Id = id

	news, err = h.NewsService.EditNews(c.Context(), news)
	if err != nil {
		return err
	}

	c.SendStatus(http.StatusOK)
	return c.JSON(news)
}

func (h newsHandler) GetNewsList(c *fiber.Ctx) error {
	list := c.QueryInt("list", 1)
	limit := c.QueryInt("limit", 10)

	news, err := h.NewsService.GetNewsList(c.Context(), list, limit)
	if err != nil {
		return err
	}

	newsList := dto.NewsList{Success: true, News: news}

	c.SendStatus(http.StatusOK)
	return c.JSON(newsList)
}
