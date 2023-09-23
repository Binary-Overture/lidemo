package main

import (
	"context"
	"firstproject/groc/service"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

func main() {
	//读取openssl证书
	//creds, err2 := credentials.NewServerTLSFromFile("groc/cert/server.pem", "groc/cert/server.key")
	//if err2 != nil {
	//	log.Fatal("证书生成错误：", err2)
	//}
	//设置拦截器进行账号密码匹配
	var logInOut grpc.UnaryServerInterceptor
	logInOut = func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		err = Auth(ctx)
		if err != nil {
			return
		}
		//继续处理请求
		return handler(ctx, req)
	}
	rpcServer := grpc.NewServer(grpc.UnaryInterceptor(logInOut))
	service.RegisterProdServiceServer(rpcServer, service.ProductService)
	//fmt.Println(creds)
	listen, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatal("启动监听", err)
	}

	err = rpcServer.Serve(listen)
	if err != nil {
		log.Fatal("启动grpc失败", err)
	}
	log.Fatal("grpc启动成功~")
}

func Auth(ctx context.Context) error {
	//拿到传输的用户名和密码
	fmt.Println("111")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("找不到值")
	}
	var user string
	var password string
	if val, ok := md["user"]; ok {
		user = val[0]
	}
	if val, ok := md["password"]; ok {
		password = val[0]
	}

	//fmt.Println(password)
	if user == "aaa" && password == "123" {

		return nil
	} else {
		fmt.Println(codes.Unauthenticated)
		return status.Errorf(codes.Unauthenticated, "账号或密码错误")
	}

}
