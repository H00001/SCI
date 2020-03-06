package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
)

func UpdateSystem() {
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
