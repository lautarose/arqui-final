package repositories

import (
	"context"
	"items/dtos"
	"items/utils/errors"
)

type Repository interface {
	GetItemById(ctx context.Context, id string) (dtos.ItemDto, errors.ApiError)
	InsertItems(ctx context.Context, item dtos.ItemsDto) (dtos.ItemsDto, errors.ApiError)
}