package model

type Business struct {
	Id int
  Created_at string
  Business_name string
  Food_type string[]
  Description string
  Locations []*Location
  Menu []*Menu_item
  Orders []*Order
}
