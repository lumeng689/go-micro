package server

import (
	"context"
	pb "github.com/lumeng689/go-micro/proto/demo"
	log "github.com/sirupsen/logrus"
)

type HelloworldServer struct{}

func (s *HelloworldServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
    log.Info("receive SayHello invoke")
	res := &pb.HelloReply{
		Message: "hello, reply for : " + in.Name,
	}
	return res, nil
}
