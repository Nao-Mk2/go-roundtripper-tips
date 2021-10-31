package main

import (
	"net/http"

	"github.com/Nao-Mk2/go-roundtripper-tips/logging"
)

func main() {
	// Logging
	lc := &http.Client{
		Transport: &logging.LoggingTransport{},
	}
	lc.Get("https://example.com")
	// 2021/11/13 00:00:00 GET https://example.com 200 OK
}
