package validation

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	unmarshalTypeError = "Unmarshal type error: "
	syntaxError        = "Syntax error: "
	invalidJSONFormat  = "Invalid JSON format"
	unexpectedEOF      = "unexpected EOF"
)

// Binder is used to create custom binder instance.
type Binder struct{}

// Bind tries to bind data and if it doesn't succeed returns error.
func (b *Binder) Bind(i interface{}, c echo.Context) (err error) {
	db := echo.DefaultBinder{}
	if err = db.Bind(i, c); err != nil {
		var (
			errMsg    interface{}
			errMsgStr string
		)

		if he, ok := err.(*echo.HTTPError); ok {
			errMsg = he.Message
			errMsgStr = errMsg.(string)
		}

		switch {
		case strings.Contains(errMsgStr, unmarshalTypeError):
			errMsgStr = b.formatErrorStr(errMsgStr)

			return &Error{Errors: []string{errMsgStr}}
		case strings.Contains(errMsgStr, syntaxError) || strings.Contains(errMsgStr, unexpectedEOF):
			return &Error{Errors: []string{invalidJSONFormat}}
		case errMsgStr != "":
			return &Error{Errors: []string{errMsgStr}}
		default:
			return &Error{Errors: []string{err.Error()}}
		}
	}

	return nil
}

func (b *Binder) formatErrorStr(errorStr string) string {
	unmarshalErrFormat := unmarshalTypeError + "expected=%s got=%s field=%s offset=%s"

	var expected, got, field, offset string

	fmt.Sscanf(errorStr, unmarshalErrFormat, &expected, &got, &field, &offset)

	expected, got, field = strings.TrimRight(expected, ","), strings.TrimRight(got, ","), strings.TrimRight(field, ",")
	if field == "" {
		return invalidJSONFormat
	}

	return fmt.Sprintf(
		"Wrong type for %s, got %s but expects %s",
		field,
		got,
		expected,
	)
}
