package model

type Item struct {
	UserID      int     `bson:"user_id"`
	Title       string  `bson:"title"`
	Seller      string  `bson:"seller"`
	Price       float64 `bson:"price"`
	Currency    string  `bson:"currency"`
	Picture     string  `bson:"picture"`
	Description string  `bson:"description"`
	State       string  `bson:"state"`
	City        string  `bson:"city"`
	Street      string  `bson:"street"`
	Number      int     `bson:"number"`
}

type Items []Item
