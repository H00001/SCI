package main

import (
	"context"
	"fmt"
	"github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"gunplan.top/SCI/protocol"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "--upgrade" {
		UpdateSystem()
		return
	}
	url := "https://p.gunplan.top/tools/serverlist.html"
	show("Get data from:"+url, green)
	resp, err := http.Get(url)
	body, err := ioutil.ReadAll(resp.Body)
	s := string(body)
	conn, err := grpc.Dial(strings.Split(s, "\n")[0], grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cnnection fail: %v", err)
	}
	defer conn.Close()
	t := connection.NewVerifyClient(conn)
	show("connection server ... ", blue)
	tr, err := t.AcquireMagic(context.Background(), &connection.MagicInbound{Uuid: uuid.NewV4().String()})
	if err != nil {
		log.Fatalf("connect error reason: %v", err)
	}
	show("verifying... UUID:"+tr.Magic, yellow)
	token, err := t.AcquireCalc(context.Background(), &connection.Base64Result{B65: tr.Magic})
	if err != nil {
		log.Fatalf("could not check sha256: %v", err)
	}
	fmt.Printf("connected!:token:%s\n", token.Token)
}
