package model

import "github.com/lib/pq"

type ClientResponse struct {
	Id                uint64 `db:"id"`
	OrderUid          string `db:"order_uid"`
	TrackNumber       string `db:"track_number"`
	Entry             string `db:"entry"`
	Delivery          Delivery
	Payment           Payment
	Items             []Item
	Locale            string `db:"locale"`
	InternalSignature string `db:"internal_signature"`
	CustomerId        string `db:"customer_id"`
	DeliveryService   string `db:"delivery_service"`
	Shardkey          string `db:"shardkey"`
	SmId              int    `db:"sm_id"`
	DateCreated       string `db:"date_created"`
	OofShard          string `db:"oof_shard"`
}

type DBResponse struct {
	Orders   Orders
	Delivery Delivery
	Payment  Payment
	Items    []Item
}

type Orders struct {
	Id                uint64        `db:"id"`
	OrderUid          string        `db:"order_uid"`
	TrackNumber       string        `db:"track_number"`
	Entry             string        `db:"entry"`
	Delivery          uint64        `db:"delivery"`
	Payment           uint64        `db:"payment"`
	Items             pq.Int64Array `db:"items"`
	Locale            string        `db:"locale"`
	InternalSignature string        `db:"internal_signature"`
	CustomerId        string        `db:"customer_id"`
	DeliveryService   string        `db:"delivery_service"`
	Shardkey          string        `db:"shardkey"`
	SmId              int           `db:"sm_id"`
	DateCreated       string        `db:"date_created"`
	OofShard          string        `db:"oof_shard"`
}

type Delivery struct {
	Id      uint64 `db:"id"`
	Name    string `db:"name"`
	Phone   string `db:"phone"`
	Zip     string `db:"zip"`
	City    string `db:"city"`
	Address string `db:"address"`
	Region  string `db:"region"`
	Email   string `db:"email"`
}

type Payment struct {
	Id           uint64 `db:"id"`
	Transaction  string `db:"transaction"`
	RequestId    string `db:"request_id"`
	Currency     string `db:"currency"`
	Provider     string `db:"provider"`
	Amount       int    `db:"amount"`
	PaymentDt    uint64 `db:"payment_dt"`
	Bank         string `db:"bank"`
	DeliveryCost int    `db:"delivery_cost"`
	GoodsTotal   int    `db:"goods_total"`
	CustomFee    int    `db:"custom_fee"`
}

type Item struct {
	Id          int64  `db:"id"`
	ChrtId      uint64 `db:"chrt_id"`
	TrackNumber string `db:"track_number"`
	Price       int    `db:"price"`
	Rid         string `db:"rid"`
	Name        string `db:"name"`
	Sale        int    `db:"sale"`
	Size        string `db:"size"`
	TotalPrice  int    `db:"total_price"`
	NmId        uint64 `db:"nm_id"`
	Brand       string `db:"brand"`
	Status      int    `db:"status"`
}
