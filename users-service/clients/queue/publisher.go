package queue

import (
	"context"
	dto "user/dtos/user"
)

type Publisher interface {
	Publish(ctx context.Context, msg dto.UserMessageDto) error
}