package log

import (
	"errors"
	"fmt"
)

// Package level errors.
var (
	errWrongLogType  = errors.New("service/log: wrong log type provided")
	errWrongLogLevel = errors.New("service/log: wrong log level provided")
)

func wrongLogLevelError(err error) error {
	return fmt.Errorf("%s: %w", errWrongLogLevel, err)
}
