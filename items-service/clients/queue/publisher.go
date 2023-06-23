package queue

import (
	"context"
	"items-service/dtos"
)

type Publisher interface {
	Publish(ctx context.Context, item dtos.ItemDto) error
}