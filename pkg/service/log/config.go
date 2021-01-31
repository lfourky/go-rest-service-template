package log

// Config configures the logger.
type Config struct {
	Type  string `default:"json"`
	Level string `default:"debug"`
}
