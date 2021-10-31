package retrying

import (
	"net/http"
	"time"
)

type RetryingTransport struct {
	Transport http.RoundTripper
}

func (t *RetryingTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

func (t *RetryingTransport) RoundTrip(req *http.Request) (res *http.Response, err error) {
	res, err = t.transport().RoundTrip(req)

	if err == nil {
		if res.StatusCode/100 == 5 {
			time.Sleep(1 * time.Second)
			// retry simply once.
			res, err = t.transport().RoundTrip(req)
		}
	}

	return res, err
}
