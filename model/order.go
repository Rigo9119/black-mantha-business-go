package model

type Order struct {
  Id int
  Created_at int
  OrderStatus order_status
  business int
  business_location int
  client_location string
  client int
}
