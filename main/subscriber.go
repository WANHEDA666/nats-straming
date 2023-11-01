package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"nats-service/db"
	"nats-service/memory"
	"nats-service/model"
	"sync"
)

var cache *memory.Cache

func main() {
	CreateCache()
	ConnectToNats()
}

func CreateCache() {
	db.Connect()
	response := db.GetDBData()
	cache = memory.NewCache(response)
}

func ConnectToNats() {
	connect, _ := stan.Connect("prod", "sub-1")
	defer panicHandler()
	defer connect.Close()
	connect.Subscribe("model", func(m *stan.Msg) {
		var response model.ClientResponse
		err := json.Unmarshal(m.Data, &response)
		if err != nil {
			fmt.Printf("Wrong model came")
			return
		}
		fmt.Printf("%v", response)
		SaveData(response)
	})

	Block()
}

func SaveData(response model.ClientResponse) {
	err := db.Insert(response)
	if err != nil {
		return
	}
	cache.Write(response)
}

func Block() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	waitGroup.Wait()
}

func panicHandler() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}
