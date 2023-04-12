package _case

import (
	"fmt"
	"sync"
)

type cache struct {
	data map[string]string
	sync.RWMutex
}

func newCache() *cache {
	return &cache{
		data:    make(map[string]string, 0),
		RWMutex: sync.RWMutex{},
	}
}
func (c *cache) Get(key string) string {
	c.RLock()
	defer c.RUnlock()
	if value, ok := c.data[key]; ok {
		return value
	}
	return ""
}

func (c *cache) Set(key, value string) {
	c.Lock()
	defer c.Unlock()
	c.data[key] = value
}
func multipleSafeRoutineByRWMutex() {
	c := newCache()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.Set("name", "wyf")
	}()
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(c.Get("name"))
		}()
	}
	wg.Wait()
}
