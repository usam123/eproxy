package main

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/Unknwon/goconfig"
)

type Config struct {
	Str        string `json:"Str"`
	Replace    string `json:"Replace"`
	IsComplete bool   `json:"IsComplete"`
}

func Init() ([]Config, string, string, error) {
	filename := "conf.ini"
	exist, err := PathExists(filename)
	if err != nil {
		//fmt.Printf("get dir error![%v]\n", err)
		return nil, "", "", err
	}

	if !exist {
		fc, err := os.Create(filename)
		if err != nil {
			return nil, "", "", err
		} else {
			_, err = fc.Write([]byte("[json]\ndata =\ntargetUrl =\nproxyUrl =\nsocks5 ="))
		}
		fc.Close()

	}

	cfg, err := goconfig.LoadConfigFile(filename)
	if err != nil {
		// panic("错误")
		return nil, "", "", err
	}

	conf_json, err := cfg.GetValue("json", "data")
	target_url, err := cfg.GetValue("json", "targetUrl")
	proxyUrl, err := cfg.GetValue("json", "proxyUrl")
	socks5, err := cfg.GetValue("json", "socks5")

	if len(target_url) == 0 {
		target_url = "https://google.com"
	}
	var conf []Config
	if conf_json != "" {
		err = json.Unmarshal([]byte(conf_json), &conf)
		if err != nil {
			return nil, "", "", err
		}

		if strings.Contains(target_url, "http") && strings.Contains(proxyUrl, "http") {
			target_url_arr := strings.Split(target_url, "://")
			proxyUrl_arr := strings.Split(proxyUrl, "://")

			conf_tp := Config{
				Str:        target_url_arr[1],
				Replace:    proxyUrl_arr[1],
				IsComplete: false,
			}
			conf = append(conf, conf_tp)

			if target_url_arr[0] != proxyUrl_arr[0] {
				conf_tp = Config{
					Str:        target_url_arr[0],
					Replace:    proxyUrl_arr[0],
					IsComplete: false,
				}
				conf = append(conf, conf_tp)
			}

		}
	}

	return conf, target_url, socks5, err
}
