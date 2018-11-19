package server

import (
	"errors"
	"net/http"
	"rpc"
)

type Args struct {
	A, B int
}

type MuliplyService struct{}

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func main() {
	service := new(MuliplyService)
	rpc.Register(MuliplyService)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}
