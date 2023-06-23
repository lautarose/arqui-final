package queue

import (
	"context"
	"items-service/dtos"
)

type PublisherMock struct{}

func (PublisherMock) Publish(ctx context.Context, item dtos.ItemDto) error {
	//TODO implement me
	panic("implement me")
}
