package queue

import (
	"context"
	"items/dtos"
)

type Publisher interface {
	Publish(ctx context.Context, item dtos.ItemDto) error
}
