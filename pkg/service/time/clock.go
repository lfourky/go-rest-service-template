package time

import (
	"time"
)

// Clock is a wrapper for Go's standard time package.
type Clock struct{}

// Now wraps the time.Now function.
func (c *Clock) Now() time.Time {
	return time.Now().UTC()
}
