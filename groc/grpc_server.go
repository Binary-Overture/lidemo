package main

import (
	"firstproject/groc/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	rpcServer := grpc.NewServer()
	service.RegisterProdServiceServer(rpcServer, service.ProductService)

	listen, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatal("启动监听", err)
	}
	err = rpcServer.Serve(listen)
	if err != nil {
		log.Fatal("启动grpc失败", err)
	}
}
