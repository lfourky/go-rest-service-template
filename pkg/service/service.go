package service

import (
	"net/http"
	"time"

	"github.com/lfourky/go-rest-service-template/pkg/service/http/option"
)

type Clock interface {
	Now() time.Time
}

type MailSender interface {
	SendMail(recipient, subject, body string) error
}

type HTTPClient interface {
	Do(method, url string, body interface{}, options ...option.RequestOption) (*http.Response, error)
}
