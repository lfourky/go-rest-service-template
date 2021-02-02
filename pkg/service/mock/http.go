package mock

import (
	"net/http"

	"github.com/lfourky/go-rest-service-template/pkg/service/http/option"
	"github.com/stretchr/testify/mock"
)

type HTTPClient struct {
	mock.Mock
}

func (m *HTTPClient) Do(method, url string, body interface{}, options ...option.RequestOption) (*http.Response, error) {
	args := m.Called(method, url, body, options)

	if args.Get(0) != nil {
		return args.Get(0).(*http.Response), args.Error(1)
	}

	return nil, args.Error(1)
}
