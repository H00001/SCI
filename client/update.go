package main

import (
	"encoding/json"
	"fmt"
	"gunplan.top/SCI/client/common"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
)

func UpdateSystem(self string) {
	data, _ := http.Get("https://p.gunplan.top/tools/version.json")
	dt, _ := ioutil.ReadAll(data.Body)
	v := &Version{}
	json.Unmarshal(dt, v)
	md5 := common.Md5File(self)
	if md5 == v.Md5 {
		show("it is the lastest version", cyan)
		return
	}
	fmt.Printf("updating... :system:%s\n", runtime.GOOS)
	res, err := http.Get("https://p.gunplan.top/tools/SCI_client_" + runtime.GOOS)
	if err != nil {
		log.Fatal(err)
		return
	}
	suffix := ""
	if runtime.GOOS == "windows" {
		suffix = ".exe"
	}
	f, _ := os.Create("SCI_client" + suffix)
	io.Copy(f, res.Body)
	show("[SUCCEED]", green)
	return
}
