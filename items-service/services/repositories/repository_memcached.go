package repositories

import (
	"context"
	"fmt"
	"items/dtos"
	e "items/utils/errors"

	"github.com/bradfitz/gomemcache/memcache"
	json "github.com/json-iterator/go"
)

type RepositoryMemcached struct {
	Client *memcache.Client
}

func NewMemcached(host string, port int) *RepositoryMemcached {
	client := memcache.New(fmt.Sprintf("%s:%d", host, port))
	fmt.Println("[Memcached] Initialized connection")
	return &RepositoryMemcached{
		Client: client,
	}
}

func (repo *RepositoryMemcached) GetItemById(ctx context.Context, id string) (dtos.ItemDto, e.ApiError) {
	item, err := repo.Client.Get(id)
	if err != nil {
		if err == memcache.ErrCacheMiss {
			return dtos.ItemDto{}, e.NewNotFoundApiError(fmt.Sprintf("item %s not found", id))
		}
		return dtos.ItemDto{}, e.NewInternalServerApiError(fmt.Sprintf("error getting item %s", id), err)
	}

	var itemDTO dtos.ItemDto
	if err := json.Unmarshal(item.Value, &itemDTO); err != nil {
		return dtos.ItemDto{}, e.NewInternalServerApiError(fmt.Sprintf("error getting item %s", id), err)
	}

	return itemDTO, nil
}

func (repo *RepositoryMemcached) InsertItems(ctx context.Context, items dtos.ItemsDto) (dtos.ItemsDto, e.ApiError) {
	for _, item := range items {
		bytes, err := json.Marshal(item)
		if err != nil {
			return dtos.ItemsDto{}, e.NewBadRequestApiError(err.Error())
		}

		if err := repo.Client.Set(&memcache.Item{
			Key:   item.Id,
			Value: bytes,
		}); err != nil {
			return dtos.ItemsDto{}, e.NewInternalServerApiError(fmt.Sprintf("error inserting book %s", item.Id), err)
		}
	}

	return items, nil
}
