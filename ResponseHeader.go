package main

import (
	"net/http"
	"strings"
)

func ResponseHeader(response *http.Response) error {
	var newCookies string

	for _, cookie := range response.Cookies() {
		//log.Println("---", cookie)
		//cookie.Domain = "localhost"
		newCookies = (cookie.String() + ";")
		if strings.HasSuffix(newCookies, ";") {
			newCookies = newCookies[:len(newCookies)-2]
		}
		response.Header.Add("Set-Cookie", newCookies)
	}

	location, err := response.Location()
	if err != nil {
		if err == http.ErrNoLocation {
			return nil
		}
		return err
	}

	// Turn it into a relative URL
	location.Scheme = ""
	location.Host = ""
	response.Header.Set("Location", location.String())
	response.Header.Del("Content-Security-Policy")
	return nil
}
