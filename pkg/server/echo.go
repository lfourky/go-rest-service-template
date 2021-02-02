package server

import (
	"strconv"
	"time"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lfourky/go-rest-service-template/pkg/service"
	"github.com/lfourky/go-rest-service-template/pkg/service/log"
)

// REST represents a REST server.
type REST struct {
	address string
	engine  *echo.Echo
	logger  *log.Logger
}

// New creates a new REST web server.
func New(cfg Config, clock service.Clock, logger *log.Logger) *REST {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{}))
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: cfg.CORS.AllowCredentials,
		AllowHeaders:     cfg.CORS.Headers,
		AllowMethods:     cfg.CORS.Methods,
		AllowOrigins:     cfg.CORS.Origins,
	}))
	e.Use(loggerMiddleware(logger, clock))

	e.Server.ReadTimeout = cfg.ReadTimeout
	e.Server.WriteTimeout = cfg.WriteTimeout

	e.Debug = cfg.Debug
	e.HideBanner = true

	e.Server.Addr = cfg.Address

	server := &REST{
		address: cfg.Address,
		engine:  e,
		logger:  logger,
	}

	return server
}

// SetErrorHandler sets the error handler.
func (r *REST) SetErrorHandler(errorHandler echo.HTTPErrorHandler) {
	r.engine.HTTPErrorHandler = errorHandler
}

// SetValidation sets the validator and binder that validate incoming payload.
func (r *REST) SetValidation(validator echo.Validator, binder echo.Binder) {
	r.engine.Validator = validator
	r.engine.Binder = binder
}

func (r *REST) Routes() *echo.Group {
	return r.engine.Group("")
}

// Run runs the REST server.
func (r *REST) Run() error {
	r.logger.WithField("address", r.address).Info("starting server")

	return gracehttp.Serve(r.engine.Server)
}

// Since logrus & echo loggers are incompatible (interface-wise).
func loggerMiddleware(logger *log.Logger, clock service.Clock) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}

			bytesIn := "0"

			contentLength := req.Header.Get(echo.HeaderContentLength)
			if contentLength != "" {
				bytesIn = contentLength
			}

			// Proceed handling the request, but return afterwards to log what happened.
			var err error
			if err = next(c); err != nil {
				c.Error(err)
			}

			logFields := map[string]interface{}{
				"id":           id,
				"time_rfc3339": clock.Now().Format(time.RFC3339),
				"time_unix":    clock.Now().Unix(),
				"host":         req.Host,
				"remote_ip":    c.RealIP(),
				"protocol":     req.Proto,
				"user_agent":   req.UserAgent(),
				"status":       res.Status,
				"uri":          req.RequestURI,
				"method":       req.Method,
				"bytes_in":     bytesIn,
				"bytes_out":    strconv.FormatInt(res.Size, 10),
			}

			logger.WithFields(logFields).Info()

			return err
		}
	}
}
