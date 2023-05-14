package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type RequestLoggingTransport struct {
	next http.RoundTripper
}

func NewRequestLoggingTransport(next http.RoundTripper) *RequestLoggingTransport {
	return &RequestLoggingTransport{next: next}
}

func (t *RequestLoggingTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Println("Request to", r.URL, "Header:", r.Header)
	if r.Body != nil {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		if len(body) > 0 {
			log.Println("Body:", string(body))
		}
	}
	return t.next.RoundTrip(r)
}
