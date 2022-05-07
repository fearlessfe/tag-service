package main

import (
	"log"
	"net"

	pb "github.com/fearlessfe/tag-service/proto"
	"github.com/fearlessfe/tag-service/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main()  {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagService())
	reflection.Register(s)

	lis, err := net.Listen("tcp", ":8001")

	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("server.Serve err: %v", err)
	}

	reflection.Register(s)
}
