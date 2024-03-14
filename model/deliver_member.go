package model

import appStatusTpes "black-mantha-business-go/utils"

type DeliveryMember struct {
	Id                     int
	Create_at              int
	Name                   string
	Last_name              string
	Phone_number           string
	Busines                int
	Busines_location       int
	Delivery_member_status appStatusTpes.DeliveryPersonStatus
}
