package services

import (
	solr "search/clients/searcher"
	"search/dtos"
	e "search/utils"
)

type ServiceImpl struct {
}

func NewServiceImpl() *ServiceImpl {
	return &ServiceImpl{}
}

func (serv *ServiceImpl) GetItemsByQuery(query string) (dtos.ItemsDto, e.ApiError) {
	var items dtos.ItemsDto

	// try to find it in localCache
	items, err := solr.GetItemsByQuery(query)
	if err != nil {
		return dtos.ItemsDto{}, e.NewBadRequestApiError("item not found")
	}

	return items, nil
}

func (serv *ServiceImpl) GetItems() (dtos.ItemsDto, e.ApiError) {
	var items dtos.ItemsDto

	// try to find it in localCache
	items, err := solr.GetItems()
	if err != nil {
		return dtos.ItemsDto{}, e.NewBadRequestApiError("item not found")
	}

	return items, nil
}

