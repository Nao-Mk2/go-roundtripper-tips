package mocking

import (
	"math/rand"
	"net/http"
	"time"
)

type MockingTransport struct {
	Transport http.RoundTripper
}

func (t *MockingTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

func (t *MockingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	res := &http.Response{
		StatusCode: 200,
		Status:     http.StatusText(200),
	}

	rand.Seed(time.Now().UnixNano())
	if rand.Uint32()%2 != 0 {
		res.StatusCode = 503
		res.Status = http.StatusText(503)
	}

	return res, nil
}
