package model

type Client struct {
  Id int
  Created_at int
  Name string
  Last_name string
  Orders int[]
  Favorites Favourites
  Email string
  Phone_number string
  Client_id string
}

type Favourites struct {
	Business []*business
	Items []*Menu_items
}
