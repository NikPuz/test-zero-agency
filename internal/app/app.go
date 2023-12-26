package app

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
	"test-zero-agency/internal/api/handler"
	"test-zero-agency/internal/api/middleware"
	"test-zero-agency/internal/app/config"
	"test-zero-agency/internal/entity"
	"test-zero-agency/internal/repository"
	"test-zero-agency/internal/service"
)

func Run(ctx context.Context, cfg *config.Config) {
	closer := newCloser()
	logger := newLogger()
	mw := middleware.NewMiddleware(logger)
	fiberApp := newFiberApp(cfg, logger)
	db := newDataBase(logger, cfg)
	reformDB := reform.NewDB(db, postgresql.Dialect, nil)

	closer.Add(fiberApp.ShutdownWithContext)
	defer db.Close()

	// Repository
	newsRepository := repository.NewNewsRepository(reformDB)

	// Service
	newsService := service.NewNewsService(cfg, newsRepository)

	// API
	handler.RegisterNewsHandlers(fiberApp, newsService, logger, mw)

	go func() {
		logger.DPanic("ListenAndServe", zap.Any("Error", fiberApp.Listen(fmt.Sprintf(":%s", cfg.Port))))
	}()

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()

	if err := closer.Close(shutdownCtx); err != nil {
		logger.Error("Close err", zap.Error(err))
	}
}

func newFiberApp(cfg *config.Config, logger *zap.Logger) *fiber.App {
	fiberApp := fiber.New(fiber.Config{
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
		ErrorHandler: func(f *fiber.Ctx, err error) error {
			response, code := entity.HandleError(f.Context(), logger, err)
			f.Send(response)
			f.SendStatus(code)
			return nil
		},
	})

	fiberApp.Use(compress.New(compress.Config{Level: compress.LevelBestSpeed}))

	return fiberApp
}

func newDataBase(logger *zap.Logger, cfg *config.Config) *sql.DB {

	conn := "postgres" + "://" + cfg.Username + ":" + cfg.Password + "@" + cfg.Address + "/" + cfg.DBName + cfg.Params

	logger.Info("start establishing a connection to the database", zap.String("Connect", conn))

	db, err := sql.Open("pgx", conn)
	if err != nil {
		panic(err.Error())
	}

	db.SetConnMaxLifetime(cfg.MaxConnLifetime)
	db.SetConnMaxIdleTime(cfg.MaxConnIdleTime)
	db.SetMaxOpenConns(cfg.DataBase.MaxOpenCons)
	db.SetMaxIdleConns(cfg.DataBase.MaxIdleCons)

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	logger.Info("successful connection to the database")

	return db
}

func newLogger() *zap.Logger {
	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	logger.Info("Zap Logger", zap.String("Level", logger.Level().String()))

	return logger
}

func newCloser() *entity.Closer {
	return &entity.Closer{}
}
