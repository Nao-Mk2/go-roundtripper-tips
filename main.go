package main

import (
	"log"
	"net/http"

	"github.com/Nao-Mk2/go-roundtripper-tips/logging"
	"github.com/Nao-Mk2/go-roundtripper-tips/mocking"
	"github.com/Nao-Mk2/go-roundtripper-tips/retrying"
)

func main() {
	// Logging
	lc := &http.Client{
		Transport: &logging.LoggingTransport{},
	}
	lc.Get("https://example.com")
	// 2021/11/13 00:00:00 GET https://example.com 200 OK

	// Mocking
	mc := &http.Client{
		Transport: &mocking.MockingTransport{},
	}
	res, _ := mc.Get("https://example.com")
	log.Printf("%d %s", res.StatusCode, http.StatusText(res.StatusCode))
	// 2021/11/13 00:00:00 200 OK
	//   or
	// 2021/11/13 00:00:00 503 Service Unavailable

	// Retrying
	rc := &http.Client{
		Transport: &retrying.RetryingTransport{
			Transport: &logging.LoggingTransport{
				Transport: &mocking.MockingTransport{},
			},
		},
	}
	rc.Get("https://example.com")
	// 2021/11/13 00:00:00 503 Service Unavailable
	// 2021/11/13 00:00:01 200 OK
}
