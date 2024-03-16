package model

import appStatusTpes "black-mantha-business-go/utils"

type Location struct {
	Id                     int
	Create_at              int
	City                   string
	Location_address       string
	Location_status        appStatusTpes.LocationStatus
	Loation_name           string
	Location_schedule      string
	Location_delivery_team []int
	Business_id            int
}
