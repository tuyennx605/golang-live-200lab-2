package rpc

import (
	"context"
	"fmt"
	"log"
	"todo-list/demogrpc/demo"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type rpcClient struct {
	client demo.ItemLikeServiceClient
}

// func NewClient(client demo.ItemLikeServiceClient) *rpcClient {
func NewClient() *rpcClient {

	// TODO: không nên khởi tạo connect ở đây, khi nào chuyển ra main và truyền vào nhé
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}

	client := demo.NewItemLikeServiceClient(cc)
	return &rpcClient{client: client}
}

func (c *rpcClient) GetItemLikes(ctx context.Context, ids []int) (map[int]int, error) {
	fmt.Println("gdfgdfgdfg")
	reqIds := make([]int32, len(ids))
	for i := range ids {
		reqIds[i] = int32(ids[i])
	}

	resp, err := c.client.GetItemLikes(ctx, &demo.GetItemLikeReq{Ids: reqIds})

	if err != nil {
		return nil, err
	}

	rs := make(map[int]int)

	for k, v := range resp.Result {
		rs[int(k)] = int(v)
	}

	return rs, nil
}
