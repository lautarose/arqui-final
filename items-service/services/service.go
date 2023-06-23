package services

import (
	"context"
	"items/dtos"
	e "items/utils/errors"
)

type Service interface {
	GetItemById(ctx context.Context, id string) (dtos.ItemDto, e.ApiError)
	InsertItems(ctx context.Context, items dtos.ItemsDto) (dtos.ItemsDto, e.ApiError)
}
