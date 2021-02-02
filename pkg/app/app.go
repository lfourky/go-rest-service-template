package app

import (
	"github.com/lfourky/go-rest-service-template/pkg/repository/postgres"
	"github.com/lfourky/go-rest-service-template/pkg/server"
	"github.com/lfourky/go-rest-service-template/pkg/server/handler"
	"github.com/lfourky/go-rest-service-template/pkg/service/http"
	"github.com/lfourky/go-rest-service-template/pkg/service/log"
	"github.com/lfourky/go-rest-service-template/pkg/service/mail"
	"github.com/lfourky/go-rest-service-template/pkg/service/time"
	"github.com/lfourky/go-rest-service-template/pkg/service/validation"
	"github.com/lfourky/go-rest-service-template/pkg/usecase"
)

func Run(cfg Config) {
	logger, err := log.New(cfg.Logger)
	if err != nil {
		panic(err)
	}

	// Setup repository.
	db, err := postgres.New(cfg.PostgresConfig, logger)
	// db, err := mysql.New(cfg.MySQLConfig, logger)
	if err != nil {
		logger.WithError(err).Fatal("unable to connect to the database")
	}

	// Setup services.
	clock := &time.Clock{}

	smtpSender := mail.NewSMTPSender(cfg.Mail)

	validator, err := validation.NewValidator()
	if err != nil {
		logger.WithError(err).Fatal("unable to create validator")
	}

	httpClient := http.NewClient(logger, http.DefaultClientOptions...)
	_ = httpClient

	// Setup server.
	restServer := server.New(cfg.Server, clock, logger)

	// Setup validation.
	restServer.SetValidation(validator, &validation.Binder{})
	restServer.SetErrorHandler(server.ErrorHandler(logger))

	// Setup usecases.
	usecase := usecase.New(db, smtpSender, logger)

	// Setup handlers.
	userHandler := handler.NewUser(usecase)
	itemHandler := handler.NewItem(usecase)
	statusHandler := handler.NewStatus(cfg.Info.Version, cfg.Info.BuildDate, cfg.Info.ServiceName, cfg.Info.CommitHash)

	// Setup routes.
	routes := restServer.Routes()
	routes.GET("/version", statusHandler.Version)

	v1 := routes.Group("v1")

	items := v1.Group("/item")
	items.POST("", itemHandler.CreateItem)
	items.GET("", itemHandler.FindAll)

	users := v1.Group("/user")
	users.GET("", userHandler.FindAll)
	users.POST("", userHandler.RegisterUser)

	if err := restServer.Run(); err != nil {
		logger.WithError(err).Error("server shutdown")
	}
}
