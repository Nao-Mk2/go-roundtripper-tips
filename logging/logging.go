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

func (t *LoggingTransport) RoundTrip(req *http.Request) (res *http.Response, err error) {
	res, err = t.transport().RoundTrip(req)

	if err == nil {
		log.Printf("%s %s %d %s", req.Method, req.URL.String(), res.StatusCode, http.StatusText(res.StatusCode))
	} else {
		log.Printf("%s %s %s", req.Method, req.URL.String(), err.Error())
	}

	return res, err
}
