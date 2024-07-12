package main

import "sync"

// Cache структура для хранения данных в оперативной памяти.
type Cache struct {
	mu    sync.RWMutex
	store map[string]interface{}
}

// New создает новый экземпляр Cache.
func New() *Cache {
	return &Cache{
		store: make(map[string]interface{}),
	}
}

// Set добавляет значение в кэш по заданному ключу.
func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = value
}

// Get возвращает значение из кэша по заданному ключу.
func (c *Cache) Get(key string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.store[key]
}

// Delete удаляет значение из кэша по заданному ключу.
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.store, key)
}
func main() {
	New()
}
