package clients

import (
	"encoding/json"
	io "io/ioutil"
	"log"
	"net/http"
	dtos "search/dtos"
)

func GetItemsByQuery(query string) (dtos.ItemsDto, error) {
	url := "http://solr:8983/solr/items/query?q=description:*" + query + "*"
	r, err := http.Get(url)

	if err != nil {
		log.Panic(err)
		return dtos.ItemsDto{}, err
	}

	bytes, err := io.ReadAll(r.Body)

	if err != nil {
		log.Panic(err)
		return dtos.ItemsDto{}, err
	}

	bodyDto, err := ParseBody(bytes)

	if err != nil {
		log.Println(err)
		return dtos.ItemsDto{}, err
	}

	docs := bodyDto.Response.Docs

	var items dtos.ItemsDto

	for _, doc := range docs {
		var item dtos.ItemDto
		item.Id = doc.Id
		item.Title = doc.Title
		item.Seller = doc.Seller
		item.Price = doc.Price
		item.Currency = doc.Currency
		item.Picture = doc.Picture
		item.Description = doc.Description
		item.State = doc.State
		item.City = doc.City
		item.Street = doc.Street
		item.Number = doc.Number

		items = append(items, item)
	}

	return items, nil
}

func ParseBody(bytes []byte) (dtos.BodyDto, error) {
	var body dtos.BodyDto
	if err := json.Unmarshal(bytes, &body); err != nil {
		log.Fatal(err)
		return dtos.BodyDto{}, err
	}
	return body, nil
}
