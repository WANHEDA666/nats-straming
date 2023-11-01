package memory

import (
	"nats-service/model"
	"sync"
)

type Cache struct {
	data  map[uint64]model.ClientResponse
	mutex sync.RWMutex
}

func NewCache(db []model.ClientResponse) *Cache {
	newMap := make(map[uint64]model.ClientResponse)
	for _, response := range db {
		newMap[response.Id] = response
	}
	return &Cache{data: newMap}
}

func (cache *Cache) Write(response model.ClientResponse) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	if _, ok := cache.data[response.Id]; !ok {
		cache.data[response.Id] = response
	}
}

func (cache *Cache) Read(id uint64) model.ClientResponse {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()
	if value, ok := cache.data[id]; ok {
		return value
	}
	return model.ClientResponse{}
}
