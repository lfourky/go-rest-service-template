package log

// Config configures the logger.
type Config struct {
	Type  string `default:"json"`
	Level string `default:"debug"`

	DefaultFields
}

// DefaultFields hold the default fields every log statement will output.
type DefaultFields struct {
	Service  string
	Hostname string
	Version  string
	Build    string
}
