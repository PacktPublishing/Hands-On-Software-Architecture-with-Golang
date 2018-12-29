package main

import (
	"log"
	"net/http"
	"net/rpc"
	"net"

)

type Args struct {
	A, B int
}

type MuliplyService struct{}

func (t *MuliplyService) Do(args *Args, reply *int) error {
	log.Println("inside MuliplyService")
	*reply = args.A * args.B
	return nil
}

func main() {
	service := new(MuliplyService)
	rpc.Register(service)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
