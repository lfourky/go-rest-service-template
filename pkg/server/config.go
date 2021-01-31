package server

import (
	"time"
)

// Config configures the server.
type Config struct {
	Debug        bool          `default:"true"`
	Address      string        `default:":8080"`
	ReadTimeout  time.Duration `default:"1m"`
	WriteTimeout time.Duration `default:"1m"`
	CORS         CORSConfig
}

// CORSConfig configures CORS.
type CORSConfig struct {
	AllowCredentials bool `default:"true"`
	Headers          []string
	Methods          []string
	Origins          []string
}
