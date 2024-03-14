package model

type MenuItem struct {
	Id               int
	Created_at       int
	Item_name        string
	Item_description string
	Item_price       []int
}

type Price struct {
	Id             int
	Created_at     string
	Price_value    float64
	Tax            float64
	Delivery_value float64
	Discount       float64
	Delivery_tip   float64
}
