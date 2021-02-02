package option

import (
	"github.com/go-resty/resty/v2"
)

type RequestOption interface {
	Apply(*resty.Request)
}

func WithAuthorizationHeader(value string) RequestOption {
	return requestOptionHeaderAuthorization(value)
}

type requestOptionHeaderAuthorization string

func (ro requestOptionHeaderAuthorization) Apply(r *resty.Request) {
	r.Header.Set("Authorization", string(ro))
}
