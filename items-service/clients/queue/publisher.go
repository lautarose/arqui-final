package queue

import (
	"context"
	"items/dtos"
)

type Publisher interface {
	PublishInsert(ctx context.Context, item dtos.ItemDto) error
	PublishDelete(ctx context.Context, id string) error
	PublishUpdate(ctx context.Context, item dtos.ItemDto) error
}
