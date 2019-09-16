package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func ErrorFunc(message string) {
	log.Println(message)
	os.Exit(-1)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}
func restart(request *http.Request) []byte {
	request.ParseForm() //解析参数，默认是不会解析的

	if len(request.Form["flag"]) > 0 {
		condition := request.Form["flag"][0]
		command := ""
		switch condition {
		//case "start":
		//command = "sudo nohup ./judas &"
		case "stop":
			command = "sudo pkill judas"
		case "restart":
			command = "sudo pkill judas && sudo nohup ./judas > judas.log 2>&1 &"
		default:
			command = ""
		}

		if command != "" {
			cmd := exec.Command("bash", "-c", command)
			bytes, err := cmd.Output()
			if err != nil {
				log.Println(err)
			}
			return bytes
			//log.Println(resp)
			//w.Write([]byte(resp))

		} else {
			// w.Write([]byte("command:为空"))
			return []byte("command:为空")
		}
	} else {
		// w.Write([]byte("flag：为空"))
		return []byte("flag：为空")
	}

}
