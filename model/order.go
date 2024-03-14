package model

import appStatusTpes "black-mantha-business-go/utils"

type Order struct {
	Id                int
	Created_at        int
	OrderStatus       appStatusTpes.OrderStatus
	Business          int
	Business_location int
	Client_location   string
	Client            int
}
