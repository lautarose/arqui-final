package dtos

type ItemReponseDto struct {
	Id          string    `json:"id"`
	Title       string  `json:"title"`
	Seller      string  `json:"seller"`
	Price       float64 `json:"price"`
	Currency    string  `json:"currency"`
	Picture     string  `json:"picture"`
	Description string  `json:"description"`
	State       string  `json:"state"`
	City        string  `json:"city"`
	Street      string  `json:"street"`
	Number      int     `json:"number"`
}

type ItemsResponseDto []ItemDto
