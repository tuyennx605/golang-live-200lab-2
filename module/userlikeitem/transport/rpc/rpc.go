package rpc

import (
	"context"
	"todo-list/demogrpc/demo"
)

type ItemStorage interface {
	GetItemLikes(ctx context.Context, ids []int) (map[int]int, error)
}

type rpc struct {
	store ItemStorage
}

func NewRPCService(store ItemStorage) demo.ItemLikeServiceServer {
	return &rpc{store: store}
}

func (s *rpc) GetItemLikes(ctx context.Context, req *demo.GetItemLikeReq) (*demo.ItemLikesResp, error) {
	ids := make([]int, len(req.Ids))

	for i := range ids {
		ids[i] = int(req.Ids[i]) // ep kieu sang int
	}

	result, err := s.store.GetItemLikes(ctx, ids)
	if err != nil {
		return nil, err
	}

	rs := make(map[int32]int32)

	for k, v := range result { // eps lai kieu int32
		rs[int32(k)] = int32(v)
	}

	return &demo.ItemLikesResp{Result: rs}, nil
}
