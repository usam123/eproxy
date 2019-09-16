package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func ResponseBody(response *http.Response, conf []Config) error {
	ctype := response.Header.Get("Content-Type")
	if !strings.Contains(ctype, "text/html") && !strings.Contains(ctype, "text/xml") {
		return nil
	}

	if len(conf) == 0 {
		return nil
	}

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil
	}
	err = response.Body.Close()
	if err != nil {
		return nil
	}

	for _, c := range conf {
		b = bytes.Replace(b, []byte(c.Str), []byte(c.Replace), -1)
	}

	body := ioutil.NopCloser(bytes.NewReader(b))
	response.Body = body
	response.ContentLength = int64(len(b))
	response.Header.Set("Content-Length", strconv.Itoa(len(b)))

	return nil
}
