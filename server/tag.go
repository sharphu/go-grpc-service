package server

import (
	"context"
	"encoding/json"
	"go-grpc-service/pkg/bapi"
	"go-grpc-service/pkg/errcode"
	pb "go-grpc-service/proto"
)

type TagServer struct {}

func NewTagServer() *TagServer {
	return &TagServer{}
}

func (t *TagServer) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
	api := bapi.NewAPI("192.168.1.43:8686")
	body, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		return nil, errcode.TogRPCError(errcode.ErrorGetTagListFail)
	}

	tagList := pb.GetTagListReply{}
	err = json.Unmarshal(body, &tagList)
	if err != nil {
		return nil, err
	}

	return &tagList, nil
}
