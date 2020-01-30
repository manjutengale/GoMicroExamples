package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	micro "github.com/micro/go-micro"
	proto "GoMicroExamples/proto"
	"github.com/micro/go-micro/registry"
	rm "github.com/micro/go-plugins/registry/consul"
	grpc "github.com/micro/go-micro/service/grpc"
	"log"
)

func main(){
	fmt.Println("Im retry client ")
	sd := rm.NewRegistry(
		registry.Addrs("http://127.0.0.1:8"),


	)
	service := grpc.NewService(
		micro.Version("one"),
		micro.Registry(sd),
	)
	clientpackage := service.Client()
	clientpackage.Init(

		client.Retries(4),
		client.Retry(func(ctx context.Context, req client.Request, retryCount int, err error) (bool, error) {
			log.Print(req.Method(), " ", "the retry count is  ", retryCount, " I re-tried  ")
			return true, nil
		}),
	)
	callserver:=proto.NewBasisService("server",clientpackage)

	serverresult,err:=callserver.Hello(context.Background(), &proto.Request{Name: "Amit"})
	if err != nil {
		fmt.Println("this is the error", err)

	} else {
		fmt.Println(serverresult) // get result
	}
}
