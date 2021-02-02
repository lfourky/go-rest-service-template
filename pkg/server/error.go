package server

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lfourky/go-rest-service-template/pkg/service/log"
	"github.com/lfourky/go-rest-service-template/pkg/service/validation"
	"github.com/lfourky/go-rest-service-template/pkg/usecase"
)

// Error represents a handler error.
type HTTPError struct {
	Message   string      `json:"message"`
	Details   interface{} `json:"details,omitempty"`
	Code      int         `json:"code,omitempty"` // Unique application error code.
	HTTPCode  int         `json:"-"`              // HTTP status code.
	Err       error       `json:"-"`
	RequestID string      `json:"request_id,omitempty"`
}

// Error returns the error message.
func (e *HTTPError) Error() string {
	return e.Message
}

func GenericInternalServerError() *HTTPError {
	return &HTTPError{
		HTTPCode: http.StatusInternalServerError,
		Message:  "internal server error",
	}
}

func NewHTTPError(err error, errorCases ...map[int][]error) error {
	httpErr := GenericInternalServerError()
	httpErr.Err = err

	for _, errCase := range errorCases {
		for httpCode := range errCase {
			for _, caseError := range errCase[httpCode] {
				if errors.Is(err, caseError) {
					// Provide custom error code if this error was associated with it.
					httpErr.HTTPCode = httpCode

					return httpErr
				}
			}
		}
	}

	return httpErr
}

func ErrCase(code int, errors ...error) map[int][]error {
	return map[int][]error{code: errors}
}

func ErrorHandler(logger *log.Logger) echo.HTTPErrorHandler {
	return func(err error, ctx echo.Context) {
		// Prevent double execution of the error handler.
		if ctx.Response().Committed {
			return
		}

		requestID := ctx.Response().Header().Get(echo.HeaderXRequestID)

		httpErr := GenericInternalServerError()
		httpErr.RequestID = ctx.Response().Header().Get(echo.HeaderXRequestID)

		if echoErr, ok := err.(*echo.HTTPError); ok {
			// Echo middleware errors (e.g. basic auth middleware).
			logger.WithFields(map[string]interface{}{
				"origin": "echo",
				"id":     requestID,
			}).WithError(echoErr).Info()

			httpErr.HTTPCode = echoErr.Code

			if msg, ok := echoErr.Message.(string); ok {
				httpErr.Message = msg
			} else {
				httpErr.Message = echoErr.Error()
			}
		} else if appErr, ok := err.(*HTTPError); ok {
			// Application specific errors.
			logger.WithFields(map[string]interface{}{
				"origin": "app",
				"id":     requestID,
			}).WithError(appErr.Err).Info()

			httpErr.HTTPCode = appErr.HTTPCode
			httpErr.Message = appErr.Message

			switch cErr := appErr.Err.(type) {
			case *usecase.Error:
				httpErr.Message = cErr.Message
				httpErr.Code = cErr.Code
				httpErr.HTTPCode = appErr.HTTPCode

			case *validation.Error:
				httpErr.Message = "validation error"
				httpErr.Details = cErr
				httpErr.HTTPCode = http.StatusBadRequest
			}
		}

		var responseErr error
		if ctx.Request().Method == http.MethodHead {
			responseErr = ctx.NoContent(httpErr.HTTPCode)
		} else {
			responseErr = ctx.JSON(httpErr.HTTPCode, httpErr)
		}

		if responseErr != nil {
			logger.WithError(err).Warn("unable to send response")
		}
	}
}
