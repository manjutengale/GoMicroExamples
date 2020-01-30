package main

import (
	"context"
	"fmt"
	micro "github.com/micro/go-micro"
	proto "GoMicroExamples/proto"
	"github.com/micro/go-micro/registry"
	rm "github.com/micro/go-plugins/registry/consul"
	grpc "github.com/micro/go-micro/service/grpc"
)
func main(){
	sd := rm.NewRegistry(
		registry.Addrs("http://127.0.0.1:8"),


	)
	fmt.Println("Im client")
	service := grpc.NewService(
		micro.Version("one"),
		micro.Registry(sd),
	)

	clientpackage := service.Client()
    callserver:=proto.NewBasisService("server",clientpackage)

    serverresult,err:=callserver.Hello(context.Background(), &proto.Request{Name: "Amit"})
	if err != nil {
		fmt.Println("this is the error", err)

	} else {
		fmt.Println(serverresult) // get result
	}
}
