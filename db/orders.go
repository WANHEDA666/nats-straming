package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"nats-service/model"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "orders"
)

var db *sqlx.DB

func Connect() *sqlx.DB {
	psqlconn := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode = disable", host, port, user, password, dbname)

	newDB, err := sqlx.Open("postgres", psqlconn)
	if err != nil {
		fmt.Printf("%s\n", err.Error()+" Can't connect to db")
		return nil
	}
	db = newDB
	return db
}

func CloseConnect() {
	db.Close()
}

func Insert(response model.ClientResponse) error {
	deliveryId := insertDelivery(response.Delivery)
	paymentId := insertPayment(response.Payment)
	itemsIds := insertItems(response.Items)
	insertion := `insert into "orders" ("order_uid", "track_number", "entry", "delivery", "payment", "items", "locale", "internal_signature", 
        "customer_id", "delivery_service", "shardkey", "sm_id", "date_created", "oof_shard") values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, 
        $11, $12, $13, $14) RETURNING Id`
	_, err := db.Exec(insertion, response.OrderUid, response.TrackNumber, response.Entry, deliveryId, paymentId, pq.Array(itemsIds), response.Locale,
		response.InternalSignature, response.CustomerId, response.DeliveryService, response.Shardkey, response.SmId, response.DateCreated, response.OofShard)
	if err != nil {
		fmt.Printf("%s\n", err.Error()+" Can't insert to orders table")
	}
	return err
}

func insertDelivery(delivery model.Delivery) int {
	lastInsertId := 0
	err := db.QueryRow("insert into delivery (name, phone, zip, city, address, region, email) values($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		delivery.Name, delivery.Phone, delivery.Zip, delivery.City, delivery.Address, delivery.Region, delivery.Email).Scan(&lastInsertId)
	if err != nil {
		fmt.Printf("%s\n", err.Error()+" Can't insert to delivery table")
		return -1
	}
	return lastInsertId
}

func insertPayment(payment model.Payment) int {
	lastInsertId := 0
	err := db.QueryRow("insert into payment (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id",
		payment.Transaction, payment.RequestId, payment.Currency, payment.Provider, payment.Amount, payment.PaymentDt, payment.Bank, payment.DeliveryCost, payment.GoodsTotal, payment.CustomFee).Scan(&lastInsertId)
	if err != nil {
		fmt.Printf("%s\n", err.Error()+" Can't insert to db to payment table")
		return -1
	}
	return lastInsertId
}

func insertItems(items []model.Item) []int {
	var itemsIds []int
	for _, item := range items {
		lastInsertId := 0
		err := db.QueryRow("insert into item (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id",
			item.ChrtId, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmId, item.Brand, item.Status).Scan(&lastInsertId)
		if err != nil {
			fmt.Printf("%s\n", err.Error()+" Can't insert to item table")
			return nil
		}
		itemsIds = append(itemsIds, lastInsertId)
	}
	return itemsIds
}

func GetDBData() []model.ClientResponse {
	deliveriesData, paymentsData, itemsData, ordersData := retrieveData()
	var result []model.ClientResponse
	for _, order := range ordersData {
		var items []model.Item
		for _, id := range order.Items {
			items = append(items, itemsData[id])
		}
		result = append(result, model.ClientResponse{
			Id:                order.Id,
			OrderUid:          order.OrderUid,
			TrackNumber:       order.TrackNumber,
			Entry:             order.Entry,
			Delivery:          deliveriesData[order.Delivery],
			Payment:           paymentsData[order.Payment],
			Items:             items,
			Locale:            order.Locale,
			InternalSignature: order.InternalSignature,
			CustomerId:        order.CustomerId,
			DeliveryService:   order.DeliveryService,
			Shardkey:          order.Shardkey,
			SmId:              order.SmId,
			DateCreated:       order.DateCreated,
			OofShard:          order.OofShard,
		})
	}
	return result
}

func retrieveData() (map[uint64]model.Delivery, map[uint64]model.Payment, map[int64]model.Item, map[uint64]model.Orders) {
	var deliveriesData []model.Delivery
	err := db.Select(&deliveriesData, "select * from delivery")
	if err != nil {
		fmt.Printf("%v", err.Error()+" Can't get full data from delivery table")
	}
	deliveriesMap := make(map[uint64]model.Delivery)
	for _, item := range deliveriesData {
		deliveriesMap[item.Id] = item
	}
	var paymentsData []model.Payment
	err = db.Select(&paymentsData, "select * from payment")
	if err != nil {
		fmt.Printf("%v", err.Error()+" Can't get full data from payment table")
	}
	paymentsMap := make(map[uint64]model.Payment)
	for _, item := range paymentsData {
		paymentsMap[item.Id] = item
	}
	var itemsData []model.Item
	err = db.Select(&itemsData, "select * from item")
	if err != nil {
		fmt.Printf("%v", err.Error()+" Can't get full data from item table")
	}
	itemsMap := make(map[int64]model.Item)
	for _, item := range itemsData {
		itemsMap[item.Id] = item
	}
	var ordersData []model.Orders
	err = db.Select(&ordersData, "select * from orders")
	if err != nil {
		fmt.Printf("%v", err.Error()+" Can't get full data from orders table")
	}
	ordersMap := make(map[uint64]model.Orders)
	for _, item := range ordersData {
		ordersMap[item.Id] = item
	}
	return deliveriesMap, paymentsMap, itemsMap, ordersMap
}
