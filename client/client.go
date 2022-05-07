package main

import (
	"context"
	"log"

	pb "github.com/fearlessfe/tag-service/proto"
	"google.golang.org/grpc"
)

func main()  {
	ctx := context.Background()
	clientConn, _ := GetClientConn(ctx, "localhost:8004", nil)
	defer clientConn.Close()

	targetServiceClient := pb.NewTagServiceClient(clientConn)
	resp, _ := targetServiceClient.GetTagList(ctx, &pb.GetTagListRequest{Name: "Go"})

	log.Printf("resp: %v", resp)
}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}
