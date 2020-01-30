package main

import (
	"context"
	"fmt"
	"log"
	"time"



	proto "GoMicroExamples/proto"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"

	rm "github.com/micro/go-plugins/registry/consul"
)



type basis struct {
}

func (b *basis) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {


	fmt.Println("this is the server msg recv ", req.Name)
	result := "result from server  " + req.Name

	rsp.Msg = result
	return nil
}
func main() {
	sd := rm.NewRegistry(
		registry.Addrs("http://127.0.0.1:8"),
	)

	srv := micro.NewService(
		micro.Name("server"),
		micro.Registry(sd),
		micro.RegisterInterval(time.Second*10),
		micro.RegisterTTL(time.Second*15),
	)
	srv.Init()
	proto.RegisterBasisHandler(srv.Server(), new(basis))
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}

}
