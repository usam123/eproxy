package main

import (
	"net/http"
	"strings"

	"github.com/Unknwon/goconfig"
)

func Router(request *http.Request) ([]byte, error) {

	if strings.Contains(request.URL.String(), "/confHttp") == true {
		cfg, err := goconfig.LoadConfigFile("conf.ini")
		if err != nil {
			return []byte("读取配置文件出错"), nil
		}
		request.ParseForm() //解析参数，默认是不会解析的

		if request.Method == "POST" {
			var (
				res       string = request.PostFormValue("json")
				targetUrl string = request.PostFormValue("targetUrl")
				proxyUrl  string = request.PostFormValue("proxyUrl")
				socks5    string = request.PostFormValue("socks5")
			)

			if strings.Contains(targetUrl, "http") {
				cfg.SetValue("json", "targetUrl", targetUrl)
			}
			if strings.Contains(proxyUrl, "http") {
				cfg.SetValue("json", "proxyUrl", proxyUrl)
			}
			if strings.Contains(socks5, ":") {
				cfg.SetValue("json", "socks5", socks5)
			}

			cfg.SetValue("json", "data", res)
			goconfig.SaveConfigFile(cfg, "conf.ini")

		}

		return []byte("ok"), nil
	}

	if strings.Contains(request.URL.String(), "/restart") == true {
		return restart(request), nil
	}

	if strings.Contains(request.URL.String(), "/config") == true {
		html := html()
		return []byte(html), nil
	}

	if strings.Contains(request.URL.String(), "insert_js.js") == true {
		html, err := ReadAll("insert_js.js")
		if err != nil {
			return []byte("alert(\"请检查js文件\")"), err
		}
		return html, nil
	}

	return []byte(""), nil
}
