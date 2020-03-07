package main

import (
	"context"
	"fmt"
	"github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"gunplan.top/SCI/common"
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
	show("Get data from:"+common.ServerList, green)
	resp, err := http.Get(common.ServerList)
	body, err := ioutil.ReadAll(resp.Body)
	s := string(body)
	conn, err := grpc.Dial(strings.Split(s, "\n")[0], grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cnnection fail: %v", err)
	}
	defer conn.Close()
	t := protocol.NewVerifyClient(conn)
	show("protocol server ... ", blue)
	tr, err := t.AcquireMagic(context.Background(), &protocol.MagicInbound{Uuid: uuid.NewV4().String()})
	if err != nil {
		log.Fatalf("connect error reason: %v", err)
		return
	}
	show("verifying... UUID:"+tr.Magic, yellow)
	token, err := t.AcquireCalc(context.Background(), &protocol.Base64Result{B65: tr.Magic})
	if err != nil {
		log.Fatalf("could not check sha256: %v", err)
		return
	}
	show("connected!:token:"+token.Token,white)
	con := protocol.NewPageClient(conn)
	for {
		var d string
		_, _ = fmt.Scanln(&d)
		if len(strings.TrimSpace(d)) == 0 {
			continue
		}
		data, _ := con.DoDown(context.Background(), &protocol.PageInbound{Content: d})
		fmt.Printf("server reply:%s\n", data.Content)
	}
}
