package main

import (
	"bufio"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Handler(conn net.Conn, client *http.Client, u *url.URL, conf []Config) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	request, err := http.ReadRequest(reader)
	if err != nil {
		log.Println("Error reader:", err.Error())
		return
	}

	response, err := Router(request)
	if err != nil {
		log.Println("Error request:", err.Error())
		return
	}

	if len(response) != 0 {
		conn.Write(response)
		return
	}

	req, err := Request(request, u)
	if err != nil {
		log.Println("Error Request:", err.Error())
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Proxy error:", err.Error())
		return
	}

	if ResponseHeader(resp) != nil {
		return
	}

	if ResponseBody(resp, conf) != nil {
		return
	}

	modifiedResponse, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Println("Error modifiedResponses:", err.Error())
		return
	}

	_, err = conn.Write(modifiedResponse)
	if err != nil {
		log.Println("Error conn.Write:", err.Error())
		return
	}
}
