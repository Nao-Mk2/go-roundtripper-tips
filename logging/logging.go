package logging

import (
	"log"
	"net/http"
)

type LoggingTransport struct {
	Transport http.RoundTripper
}

func (t *LoggingTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

func (t *LoggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	res, err := t.transport().RoundTrip(req)

	if err == nil {
		// 2021/11/11 00:00:00 GET https://example.com 200 OK
		log.Printf(
			"%s %s %d %s",
			req.Method,
			req.URL.String(),
			res.StatusCode,
			http.StatusText(res.StatusCode),
		)
	}

	return res, err
}
