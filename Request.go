package main

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"strings"
)

func Request(request *http.Request, u *url.URL) (*http.Request, error) {

	target := request.URL
	target.Scheme = u.Scheme
	target.Host = u.Host

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: false}

	req, err := http.NewRequest(request.Method, target.String(), request.Body)
	if err != nil {
		return nil, err
	}
	for key := range request.Header {
		req.Header.Set(key, request.Header.Get(key))
	}

	if request.Referer() != "" {
		req.Header.Set("Referer", strings.Replace(request.Referer(), request.Host, u.Host, -1))
	}

	req.Header.Del("Accept-Encoding")
	return req, nil
}
