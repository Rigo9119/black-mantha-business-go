package model

type Client struct {
	Id              int
	Created_at      int
	Name            string
	Last_name       string
	Orders          []Order
	Favorites       Favourites
	Email           string
	Phone_number    string
	Client_id       string
	Client_location []string
}

type Favourites struct {
	Business []*Business
	Items    []*MenuItem
}
