package dtos

type ItemUpdateDto struct {
	Id          string     `json:"id"`
	Title       FieldValue `json:"title"`
	Seller      FieldValue `json:"seller"`
	Price       FieldValue `json:"price"`
	Currency    FieldValue `json:"currency"`
	Picture     FieldValue `json:"picture"`
	Description FieldValue `json:"description"`
	State       FieldValue `json:"state"`
	City        FieldValue `json:"city"`
	Street      FieldValue `json:"street"`
	Number      FieldValue `json:"number"`
}

type ItemsUpdateDto []ItemUpdateDto
