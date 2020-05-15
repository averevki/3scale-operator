package helper

import (
	"log"
	"net/http"
	"net/http/httputil"
)

// Transport implements http.RoundTripper. When set as Transport of http.Client, it executes HTTP requests with logging.
// No field is mandatory.
type Transport struct {
	Transport http.RoundTripper
}

// DefaultTransport is the default logging transport that wraps http.DefaultTransport.
var DefaultTransport = &Transport{
	Transport: http.DefaultTransport,
}

// RoundTrip is the core part of this module and implements http.RoundTripper.
// Executes HTTP request with request/response logging.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.logRequest(req)

	resp, err := t.transport().RoundTrip(req)
	if err != nil {
		return resp, err
	}

	t.logResponse(resp)

	return resp, err
}

func (t *Transport) logRequest(req *http.Request) {
	dump, _ := httputil.DumpRequestOut(req, true)
	log.Println(string(dump))
}

func (t *Transport) logResponse(resp *http.Response) {
	dump, _ := httputil.DumpResponse(resp, true)
	log.Println(string(dump))
}

func (t *Transport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}

	return http.DefaultTransport
}
