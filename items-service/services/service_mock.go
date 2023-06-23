package services

import (
	"items/dtos"
	e "items/utils/errors"
)

type ServiceMock struct{}

func NewServiceMock() ServiceMock {
	return ServiceMock{}
}

func (ServiceMock) GetItemById(id string) (dtos.ItemDto, e.ApiError) {
	return dtos.ItemDto{
		Id:          "1",
		Title:      "Casa",
		Seller:      "Particular",
		Price:       150000,
		Currency:    "usd",
		Picture:     "https://www.bbva.com/wp-content/uploads/2021/04/casas-ecolo%CC%81gicas_apertura-hogar-sostenibilidad-certificado--1024x629.jpg",
		Description: "Casa 2 plantas con pileta",
		State:       "Cordoba",
		City:        "Cordoba",
		Street:      "Belgrano",
		Number:      1115,
	}, nil
}

func (ServiceMock) InsertItem(item dtos.ItemDto) (dtos.ItemDto, e.ApiError) {
	return dtos.ItemDto{
		Id:          "1",
		Title:      item.Title,
		Seller:      item.Seller,
		Price:       item.Price,
		Currency:    item.Currency,
		Picture:     item.Picture,
		Description: item.Description,
		State:       item.State,
		City:        item.City,
		Street:      item.Street,
		Number:      item.Number,
	}, nil
}
