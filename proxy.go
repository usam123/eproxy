package main

import (
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"net/url"
	"time"

	"golang.org/x/net/proxy"
)

const (
	DefaultTimeout = 20 * time.Second
)

func main() {
	conf, target_url, socks5, err := Init()
	if err != nil {
		ErrorFunc(err.Error())
	}

	u, err := url.Parse(target_url)
	if err != nil {
		ErrorFunc(err.Error())
	}

	client := &http.Client{
		Timeout: DefaultTimeout,
	}

	if socks5 != "" {
		dialer, err := proxy.SOCKS5("tcp", socks5, nil, proxy.Direct)
		if err != nil {
			ErrorFunc(err.Error())
		}
		httpTransport := &http.Transport{}
		httpTransport.Dial = dialer.Dial
		client.Transport = httpTransport

		log.Println("socks5: ", socks5)
	}

	var server net.Listener
	server, err = net.Listen("tcp", ":80")
	log.Println("Listening on:80")
	if err != nil {
		ErrorFunc(err.Error())
	}

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Println("Error server.Accept:", err.Error())
			continue
		}
		go Handler(conn, client, u, conf)
	}
}
