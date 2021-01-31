package app

import (
	"github.com/lfourky/go-rest-service-template/pkg/handler"
	"github.com/lfourky/go-rest-service-template/pkg/service/log"
	"github.com/lfourky/go-rest-service-template/pkg/service/mail"
	"github.com/lfourky/go-rest-service-template/pkg/service/usecase"
	"github.com/sirupsen/logrus"
)

// Run runs the application.
// nolint: funlen
func Run(cfg Config) {
	logger, err := log.New(cfg.Logger)
	if err != nil {
		panic(err)
	}

	logger.WithFields(logrus.Fields{
		"version":   cfg.Info.Version,
		"buildDate": cfg.Info.BuildDate,
	}).Info(cfg.Info.Description)

	// mysqlRepo := mysql.NewRepository(db)
	smtpSender := mail.NewSMTPSender(1234, "", "", "", "", "")

	usecase := usecase.New(nil, smtpSender)

	userHandler := handler.NewUserHandler(usecase)
	itemHandler := handler.NewItemHandler(usecase)

	_ = userHandler
	_ = itemHandler
}
