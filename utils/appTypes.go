package appStatusTpes

type DeliveryPersonStatus string
type LocationStatus string
type OrderStatus string

const (
	INACTIVE   DeliveryPersonStatus = "inactive"
	BUSSY      DeliveryPersonStatus = "bussy"
	ON_THE_WAY DeliveryPersonStatus = "on_the_way"
)

const (
	OPEN   LocationStatus = "open"
	CLOSED LocationStatus = "closed"
)

const (
	ACCEPTED        OrderStatus = "accepted"
	DELIVERED       OrderStatus = "delivered"
	CANCELED        OrderStatus = "canceled"
	PREPARING       OrderStatus = "preparing"
	READY_TO_PICKUP OrderStatus = "ready_to_pickup"
	DELAYED         OrderStatus = "delayed"
)
