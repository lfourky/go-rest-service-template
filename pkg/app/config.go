package app

import (
	"github.com/lfourky/go-rest-service-template/pkg/repository/mysql"
	"github.com/lfourky/go-rest-service-template/pkg/repository/postgres"
	"github.com/lfourky/go-rest-service-template/pkg/server"
	"github.com/lfourky/go-rest-service-template/pkg/service/log"
	"github.com/lfourky/go-rest-service-template/pkg/service/mail"
)

type Config struct {
	Logger         log.Config
	Server         server.Config
	Info           InfoConfig
	PostgresConfig postgres.Config
	MySQLConfig    mysql.Config
	Mail           mail.Config
}

type InfoConfig struct {
	Version     string
	BuildDate   string
	ServiceName string
	CommitHash  string
}
