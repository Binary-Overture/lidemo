package main

import (
	"context"
	"firstproject/groc/service"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	//建立到grpc服务器的连接，就下面的pbfile吧，WithTransportCredentials是可选项选择，
	//insecure.NewCredentials()函数是创建一个不安全的TLS配置，也就是不进行加密和身份验证.
	dial, err := grpc.Dial(":8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("服务器连接出错", err)
	}
	//关闭建立的连接
	defer func(dial *grpc.ClientConn) {
		err := dial.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(dial)

	proClient := service.NewProdServiceClient(dial)

	request := &service.ProductRequest{
		ProId: 123,
	}
	stockResponse, err := proClient.GetProductStock(context.Background(), request)
	if err != nil {
		log.Fatal("查询库存出错", stockResponse)
	}
	fmt.Println("查询成功，库存为：", stockResponse)
}
