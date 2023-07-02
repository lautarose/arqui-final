package repositories

import (
	"context"
	"items/dtos"
	e "items/utils/errors/errors"
)

type Repository interface {
	GetItemById(ctx context.Context, id string) (dtos.ItemDto, e.ApiError)
	GetItemsIdByUserId(ctx context.Context, userId string) ([]string, e.ApiError)
	InsertItems(ctx context.Context, items dtos.ItemsDto) (dtos.ItemsDto, e.ApiError)
	UpdateItem(ctx context.Context, item dtos.ItemDto) (dtos.ItemDto, e.ApiError)
	DeleteItem(ctx context.Context, id string) e.ApiError
}
