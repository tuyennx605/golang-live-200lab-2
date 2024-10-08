package resapi

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

// https://github.com/go-resty/resty
// go get github.com/go-resty/resty/v2

type itemService struct {
	client     *resty.Client
	serviceURL string
}

func New(serviceURL string) *itemService {
	return &itemService{
		client:     resty.New(),
		serviceURL: serviceURL,
	}
}

func (s *itemService) GetItemLikes(ctx context.Context, ids []int) (map[int]int, error) {
	type requestBody struct {
		Ids []int `json:"ids"`
	}

	var response struct {
		Data map[int]int `json:"data"`
	}

	resp, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(requestBody{Ids: ids}).
		SetResult(&response).
		Post(fmt.Sprintf("%s/%s", s.serviceURL, "v1/rpc/get_item_likes"))

	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		log.Println(resp.RawResponse)
		return nil, errors.New("cannot call api get item likes")
	}

	return response.Data, nil
}
