package services

import (
	"search/dtos"
	e "search/utils"
)

type Service interface {
	GetItemsByQuery(query string) (dtos.ItemsDto, e.ApiError)
}
