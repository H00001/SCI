package main

import (
	"context"
	"fmt"
	"github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"gunplan.top/SCI/protocol"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "--upgrade" {
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
		fmt.Printf("\033[42m;[SUCCEED]\033[0m\n")
		return
	}
	resp, err := http.Get("https://p.gunplan.top/tools/serverlist.html")
	body, err := ioutil.ReadAll(resp.Body)
	s := string(body)
	conn, err := grpc.Dial(strings.Split(s, "\n")[0], grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cnnection fail: %v", err)
	}
	defer conn.Close()
	t := connection.NewVerifyClient(conn)
	fmt.Print("connection server ... \n")
	tr, err := t.AcquireMagic(context.Background(), &connection.MagicInbound{Uuid: uuid.NewV4().String()})
	if err != nil {
		log.Fatalf("connect error reason: %v", err)
	}
	fmt.Printf("verifying... UUID:%s\n", tr.Magic)
	token, err := t.AcquireCalc(context.Background(), &connection.Base64Result{B65: tr.Magic})
	if err != nil {
		log.Fatalf("could not check sha256: %v", err)
	}
	fmt.Printf("connected!:token:%s\n", token.Token)
}
