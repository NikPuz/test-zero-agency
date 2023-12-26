package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"

	"go.uber.org/zap"
)

type ctxKeyRequestID int

const RequestIDKey ctxKeyRequestID = 0

type Middleware struct {
	logger *zap.Logger
}

func NewMiddleware(logger *zap.Logger) *Middleware {
	return &Middleware{logger: logger}
}

func (m *Middleware) ContentTypeJSON(f *fiber.Ctx) error {
	f.Set("Content-Type", "application/json; charset=UTF-8")
	return f.Next()
}

func (m *Middleware) DebugLogger(f *fiber.Ctx) error {
	timeStart := time.Now()

	err := f.Next()

	m.logger.Debug("Request Logger",
		zap.String("LeadTime", fmt.Sprintf("%.3f", time.Duration(time.Now().UnixNano()-timeStart.UnixNano()).Seconds())),
		zap.String("RequestMethod", f.Method()),
		zap.String("URL", f.OriginalURL()),
		zap.String("IP", f.IP()),
	)

	return err
}
