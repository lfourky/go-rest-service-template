package http

import (
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/lfourky/go-rest-service-template/pkg/service/log"
)

type (
	clientOption interface {
		Apply(*Client)
	}

	requestOption interface {
		Apply(*resty.Request)
	}
)

type Client struct {
	retryer *resty.Client

	logger *log.Logger
}

func NewClient(logger *log.Logger, options ...clientOption) *Client {
	client := &Client{
		retryer: resty.New(),
		logger:  logger,
	}

	for _, opt := range options {
		opt.Apply(client)
	}

	return client
}

func (c *Client) Do(method, url string, body interface{}, options ...requestOption) (*http.Response, error) {
	request := c.retryer.NewRequest()

	for _, opt := range options {
		opt.Apply(request)
	}

	resp, err := request.Execute(method, url)
	if err != nil {
		return nil, err
	}

	return resp.RawResponse, nil
}

var DefaultClientOptions = []clientOption{
	ClientOptionTraceEnabled(true),
	ClientOptionRetryConditions{
		http.StatusRequestTimeout,
		http.StatusInternalServerError,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout,
	},
	ClientOptionRetryConfiguration{
		MaxRetries:  2,
		WaitTime:    time.Second,
		MaxWaitTime: time.Second * 2,
		Timeout:     time.Second * 3,
	},
}

type ClientOptionRetryConditions []int

func (co ClientOptionRetryConditions) Apply(c *Client) {
	c.retryer.AddRetryCondition(func(r *resty.Response, e error) bool {
		for _, status := range co {
			if r.StatusCode() == status {
				return true
			}
		}

		return false
	})
}

type ClientOptionTraceEnabled bool

func (co ClientOptionTraceEnabled) Apply(c *Client) {
	if co {
		c.retryer.EnableTrace()
	} else {
		c.retryer.DisableTrace()
	}
}

type ClientOptionRetryConfiguration struct {
	MaxRetries  uint
	WaitTime    time.Duration
	MaxWaitTime time.Duration
	Timeout     time.Duration
}

func (co ClientOptionRetryConfiguration) Apply(c *Client) {
	c.retryer.
		SetRetryWaitTime(co.WaitTime).
		SetRetryCount(int(co.MaxRetries)).
		SetRetryMaxWaitTime(co.MaxWaitTime).
		SetTimeout(co.Timeout)
}

type ClientOptionBasicAuth struct {
	Username string
	Password string
}

func (co ClientOptionBasicAuth) Apply(c *Client) {
	c.retryer.SetBasicAuth(co.Username, co.Password)
}

type ClientOptionAuthToken string

func (co ClientOptionAuthToken) Apply(c *Client) {
	c.retryer.SetAuthToken(string(co))
}
