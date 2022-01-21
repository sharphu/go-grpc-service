package main

import (
	"context"
	pb "go-grpc-service/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	ctx := context.Background()
	clientConn, _ := GetClientConn(ctx, "192.168.1.43:8686", nil)
	defer clientConn.Close()

	tagServiceClient := pb.NewTagServiceClient(clientConn)
	resp, _ := tagServiceClient.GetTagList(
		ctx,
		&pb.GetTagListRequest{Name: "Go"},
		)
	log.Printf("resp: %v", resp)
}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}
