package app

import (
	"github.com/lfourky/go-rest-service-template/pkg/service/log"
)

// Config holds the general application configuration.
type Config struct {
	Logger log.Config
	Info   InfoConfig
}

// InfoConfig configures application information.
type InfoConfig struct {
	Version     string
	BuildDate   string
	Description string
	CommitHash  string
}
