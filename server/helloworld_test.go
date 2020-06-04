package server

import (
	"context"
	"fmt"
	pb "github.com/lumeng689/go-micro/proto/demo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"testing"
)

func connectSecureGreeterServer() pb.GreeterClient {
	const SERVER_ADDR string = "127.0.0.1:50053"
	dcreds, err := credentials.NewClientTLSFromFile("../conf/certs/server.pem", CertServerName)
	if err != nil {
		fmt.Printf("get cert failed!!")
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}
	//opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(SERVER_ADDR, opts...)
	if err != nil {
		fmt.Printf("Cannot connect %v", err)
	}
	return pb.NewGreeterClient(conn)
}

func connectUnSecureGreeterServer() pb.GreeterClient {
	const SERVER_ADDR string = "127.0.0.1:50053"
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(SERVER_ADDR, opts...)

	if err != nil {
		fmt.Printf("Cannot connect %v", err)
	}
	return pb.NewGreeterClient(conn)
}

func Test_Hello1(t *testing.T) {

	//echoClient := connectSecureGreeterServer()
	echoClient := connectUnSecureGreeterServer()

	helloRequest := &pb.HelloRequest{
		Name: "张三",
	}
	res, err := echoClient.SayHello(context.Background(), helloRequest)

	fmt.Printf("err: %v \n", err)
	fmt.Printf("res: %v \n", res)
}
