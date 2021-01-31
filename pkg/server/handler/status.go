package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Status manages status related endpoints.
type Status struct {
	version     string
	buildDate   string
	description string
	commitHash  string
}

// NewStatus creates a new status handler.
func NewStatus(version, buildDate, description, commitHash string) *Status {
	return &Status{
		version:     version,
		buildDate:   buildDate,
		description: description,
		commitHash:  commitHash,
	}
}

// Version provides the application version information.
func (h *Status) Version(ctx echo.Context) error {
	return ctx.String(http.StatusOK, fmt.Sprintf(`version="%s" buildDate="%s" serviceName="%s" commit="%s"`,
		h.version,
		h.buildDate,
		h.description,
		h.commitHash,
	))
}
