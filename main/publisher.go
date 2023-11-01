package main

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"nats-service/model"
	"time"
)

func main() {

	firstItem := model.Item{
		ChrtId:      9934930,
		TrackNumber: "WBILMTESTTRACK",
		Price:       453,
		Rid:         "ab4219087a764ae0btest",
		Name:        "Mascaras",
		Sale:        30,
		Size:        "0",
		TotalPrice:  317,
		NmId:        2389212,
		Brand:       "Vivienne Sabo",
		Status:      202,
	}

	secondItem := model.Item{
		ChrtId:      1111111,
		TrackNumber: "MMMMMMMMMMMMMM",
		Price:       444,
		Rid:         "atestatestatestatest",
		Name:        "Sascaras",
		Sale:        55,
		Size:        "7",
		TotalPrice:  217,
		NmId:        3333333,
		Brand:       "Sabo Vivienne",
		Status:      203,
	}

	response := model.ClientResponse{
		OrderUid:    "b563feb7b2b84b6test",
		TrackNumber: "WBILMTESTTRACK",
		Entry:       "WBIL",
		Delivery: model.Delivery{
			Name:    "Test Testov",
			Phone:   "+9720000000",
			Zip:     "2639809",
			City:    "Kiryat Mozkin",
			Address: "Ploshad Mira 15",
			Region:  "Kraiot",
			Email:   "test@gmail.com",
		},
		Payment: model.Payment{
			Transaction:  "b563feb7b2b84b6test",
			RequestId:    "",
			Currency:     "USD",
			Provider:     "wbpay",
			Amount:       1817,
			PaymentDt:    1637907727,
			Bank:         "alpha",
			DeliveryCost: 1500,
			GoodsTotal:   317,
			CustomFee:    0,
		},
		Items:             []model.Item{firstItem, secondItem},
		Locale:            "en",
		InternalSignature: "",
		CustomerId:        "test",
		DeliveryService:   "meest",
		Shardkey:          "9",
		SmId:              99,
		DateCreated:       "2021-11-26T06:22:19Z",
		OofShard:          "1",
	}

	var bytes, _ = json.Marshal(response)

	connect, _ := stan.Connect("prod", "simple-pub")
	defer connect.Close()

	for i := 1; ; i++ {
		connect.Publish("model", bytes)
		time.Sleep(2 * time.Second)
	}
}
