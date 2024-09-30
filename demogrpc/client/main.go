package main

import (
	"context"
	"log"
	"todo-list/demogrpc/demo"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}

	defer cc.Close()

	client := demo.NewItemLikeServiceClient(cc)

	resp, _ := client.GetItemLikes(context.Background(), &demo.GetItemLikeReq{Ids: []int32{1, 2, 3}})

	log.Println(resp.Result)
}
