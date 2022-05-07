package server

import (
	"context"
	"encoding/json"

	"github.com/fearlessfe/tag-service/pkg/errcode"
	pb "github.com/fearlessfe/tag-service/proto"

	"github.com/fearlessfe/tag-service/pkg/bapi"
)

type TagService struct {
	pb.UnimplementedTagServiceServer
}


func NewTagService() *TagService {
	return &TagService{}
}

func (t *TagService) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
	api := bapi.NewAPI("http://127.0.0.1:8000")
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
