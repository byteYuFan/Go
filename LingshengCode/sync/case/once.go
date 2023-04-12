package _case

import (
	"fmt"
	"sync"
)

type onceMap struct {
	sync.Once
	data map[string]int
}

func (o *onceMap) LoadData() {
	o.Do(func() {
		list := []string{"A", "B", "C", "D"}
		for _, item := range list {
			if _, ok := o.data[item]; !ok {
				o.data[item] = 0
			}
			o.data[item] += 1
		}
	})
}

func OnceCase() {
	o := &onceMap{
		data: make(map[string]int),
	}
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			o.LoadData()
		}()
	}
	wg.Wait()
	fmt.Println(o.data)
}
