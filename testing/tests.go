package testing

import (
	"github.com/nats-io/stan.go"
	"github.com/stretchr/testify/assert"
	"nats-service/db"
	"nats-service/memory"
	"nats-service/model"
	"os"
	"sync"
	"testing"
)

var cache *memory.Cache

func main() {
}

func TestAll(t *testing.T) {
	t.Run("TestCreateCache", TestCreateCache)
	t.Run("TestSaveData", TestSaveData)
	t.Run("TestConnectToNats", TestConnectToNats)
	t.Run("TestBlock", TestBlock)
}

func TestCreateCache(t *testing.T) {
	db.Connect()
	defer db.CloseConnect()
	response := db.GetDBData()
	cache = memory.NewCache(response)

	assert.NotNil(t, cache, "Cache should not be nil")
}

func TestSaveData(t *testing.T) {
	response := model.ClientResponse{
		Id:                0,
		OrderUid:          "",
		TrackNumber:       "",
		Entry:             "",
		Delivery:          model.Delivery{},
		Payment:           model.Payment{},
		Items:             nil,
		Locale:            "",
		InternalSignature: "",
		CustomerId:        "",
		DeliveryService:   "",
		Shardkey:          "",
		SmId:              0,
		DateCreated:       "",
		OofShard:          "",
	}

	err := SaveData(response)

	assert.NoError(t, err, "SaveData should not return an error")
}

func TestConnectToNats(t *testing.T) {
	connect, _ := stan.Connect("test-cluster", "test-subscriber")
	defer connect.Close()

	// Здесь нужно отправить сообщение в NATS и убедиться, что оно обработано
}

func TestBlock(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		Block()
		wg.Done()
	}()
	wg.Wait()
}

func TestMain(m *testing.M) {
	db.Connect()
	defer db.CloseConnect()
	code := m.Run()
	os.Exit(code)
}

func SaveData(response model.ClientResponse) error {
	err := db.Insert(response)
	if err != nil {
		return err
	}
	cache.Write(response)
	return err
}

func Block() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	waitGroup.Wait()
}
