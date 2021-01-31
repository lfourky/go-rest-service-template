package log

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

const (
	logTypeJSON = "json"
	logTypeText = "text"

	logLevelDebug = "DEBUG"
)

// Logger is the application logger.
type Logger struct {
	engine *log.Logger
	*log.Entry
}

// New creates a new logger.
func New(cfg Config) (*Logger, error) {
	engine := log.New()

	switch cfg.Type {
	case logTypeJSON:
		engine.Formatter = &log.JSONFormatter{}
	case logTypeText:
		engine.Formatter = &log.TextFormatter{
			TimestampFormat: "2006-01-02T15:04:05.000",
			FullTimestamp:   true,
		}
	default:
		return nil, fmt.Errorf("bad log type provided; supported log types are: [json,text]; got: %s", cfg.Type)
	}

	lvl, err := log.ParseLevel(cfg.Level)
	if err != nil {
		return nil, fmt.Errorf("bad log level provided: %w", err)
	}

	engine.SetLevel(lvl)

	if cfg.Level == logLevelDebug {
		engine.SetReportCaller(true)
	}

	contextLogger := engine.WithFields(log.Fields{
		"service":  cfg.DefaultFields.Service,
		"hostname": cfg.DefaultFields.Hostname,
		"version":  cfg.DefaultFields.Version,
		"build":    cfg.DefaultFields.Build,
	})

	return &Logger{
		engine: engine,
		Entry:  contextLogger,
	}, nil
}

// TestLogger is used in unit & integration tests.
func TestLogger() *Logger {
	l, _ := New(Config{
		Level: logLevelDebug,
		Type:  logTypeText,
	})

	return l
}
