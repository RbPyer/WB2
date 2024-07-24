package patterns

import (
	"fmt"
	"sync"
)

type DataFacade struct {
	db    *Database
	cache *INMEMCache
}

func NewDataFacade() *DataFacade {
	return &DataFacade{
		db: &Database{},
		cache: &INMEMCache{
			m:    new(sync.RWMutex),
			data: make(map[string]string),
		},
	}
}

func (f *DataFacade) Work() {
	f.db.Select(1)
	f.db.Delete(2)
	f.cache.Set("hello", "world")
	f.cache.Get("hello")
}

type INMEMCache struct {
	m    *sync.RWMutex
	data map[string]string
}

type Database struct{}

func (d *Database) Select(id int) string {
	return fmt.Sprintf("some information from database with id %d", id)
}

func (d *Database) Insert(id int, data string) {
	fmt.Printf("Inserted data into database with id %d with data %s\n", id, data)
}

func (d *Database) Update(id int, data string) {
	fmt.Printf("Updated data into database with id %d with data %s\n", id, data)
}

func (d *Database) Delete(id int) {
	fmt.Printf("Deleted data into database with id %d\n", id)
}

func (c *INMEMCache) Get(key string) string {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.data[key]
}

func (c *INMEMCache) Set(key string, value string) {
	c.m.Lock()
	c.data[key] = value
	c.m.Unlock()
}

func (c *INMEMCache) Del(key string) {
	c.m.Lock()
	delete(c.data, key)
	c.m.Unlock()
}
