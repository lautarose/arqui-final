package repositories

import (
	"context"
	"fmt"
	"items/dtos"
	e "items/utils/errors/errors"
	"time"

	"github.com/karlseguin/ccache/v2"
)

type RepositoryCCache struct {
	Client     *ccache.Cache
	DefaultTTL time.Duration
}

func NewCCache(maxSize int64, itemsToPrune uint32, defaultTTL time.Duration) *RepositoryCCache {
	client := ccache.New(ccache.Configure().MaxSize(maxSize).ItemsToPrune(itemsToPrune))
	fmt.Println("[CCache] Initialized")
	return &RepositoryCCache{
		Client:     client,
		DefaultTTL: defaultTTL,
	}
}

func (repo *RepositoryCCache) GetItemById(ctx context.Context, id string) (dtos.ItemDto, e.ApiError) {
	item := repo.Client.Get(id)
	if item == nil {
		return dtos.ItemDto{}, e.NewNotFoundApiError(fmt.Sprintf("item %s not found", id))
	}
	if item.Expired() {
		return dtos.ItemDto{}, e.NewNotFoundApiError(fmt.Sprintf("item %s not found", id))
	}
	return item.Value().(dtos.ItemDto), nil
}

func (repo *RepositoryCCache) InsertItems(ctx context.Context, items dtos.ItemsDto) (dtos.ItemsDto, e.ApiError) {
	for _, item := range items {
		repo.Client.Set(item.Id, item, repo.DefaultTTL)
	}

	return items, nil
}

func (repo *RepositoryCCache) UpdateItem(ctx context.Context, item dtos.ItemDto) (dtos.ItemDto, e.ApiError) {
	existingItem := repo.Client.Get(item.Id)
	if existingItem == nil {
		var items dtos.ItemsDto
		items = append(items, item)
		returnItems, err := repo.InsertItems(ctx, items)
		item = returnItems[0]
		return item, err
	}

	repo.Client.Set(item.Id, item, repo.DefaultTTL)

	return item, nil
}

func (repo *RepositoryCCache) DeleteItem(ctx context.Context, id string) e.ApiError {
	repo.Client.Delete(id)
	return nil
}

func (repo *RepositoryCCache) GetItemsIdByUserId(ctx context.Context, userId string) ([]string, e.ApiError) {
	return nil, nil
}
