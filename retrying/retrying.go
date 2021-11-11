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

func (t *RetryingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	res, err := t.transport().RoundTrip(req)

	if err == nil {
		if res.StatusCode == 503 {
			time.Sleep(1 * time.Second)
			// retry once.
			res, err = t.transport().RoundTrip(req)
		}
	}

	return res, err
}
