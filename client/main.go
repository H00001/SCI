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

func main(){
	if os.Args[0] == "--upgrade" {
		fmt.Print("updating...")
		res, _ := http.Get("https://p.gunplan.top/tools/SCI_client_" + runtime.GOOS)
		suffix := ""
		if runtime.GOOS == "windows" {
			suffix = ".exe"
		}
		f, _ := os.Create("SCI_client" + suffix)
		io.Copy(f, res.Body)
		return
	}
	resp,err := http.Get("https://p.gunplan.top/tools/serverlist.html")
	body,err := ioutil.ReadAll(resp.Body)
	s := string(body)
	conn, err := grpc.Dial(strings.Split(s,"\n")[0], grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cnnection fail: %v", err)
	}
	defer conn.Close()
	t := connection.NewVerifyClient(conn)
	fmt.Print("connection server ... \n")
	tr, err := t.AcquireMagic(context.Background(), &connection.MagicInbound{ Uuid:uuid.NewV4().String()})
	fmt.Printf("verifying... UUID:%s\n",tr.Magic)
	t.AcquireCalc(context.Background(),&connection.Base64Result{B65:tr.Magic})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("connected!:%s\n",tr.Magic)
}
