package _case

import (
	"fmt"
	"sync"
)

func MutexCase() {
	//singleRoutine()
	//multiRoutine()
	//multiSafeRoutine()
	multipleSafeRoutineByRWMutex()
}

// 单协程操作
func singleRoutine() {
	mp := make(map[string]int, 0)
	list := []string{"A", "B", "C", "D"}
	for i := 0; i < 100; i++ {
		for _, v := range list {
			if _, ok := mp[v]; !ok {
				mp[v] = 0
			}
			mp[v] += 1
		}
	}
	fmt.Println(mp)
}

// 多协程 非线程安全
func multiRoutine() {
	mp := make(map[string]int, 0)
	list := []string{"A", "B", "C", "D"}
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, v := range list {
				if _, ok := mp[v]; !ok {
					mp[v] = 0
				}
				mp[v] += 1
			}
		}()
	}
	wg.Wait()
	fmt.Println(mp)
}

// 互斥锁协程安全
func multiSafeRoutine() {
	type safeMap struct {
		data map[string]int
		sync.Mutex
	}
	mp := safeMap{
		data:  make(map[string]int, 0),
		Mutex: sync.Mutex{},
	}
	list := []string{"A", "B", "C", "D"}
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			mp.Lock()
			defer mp.Unlock()
			for _, v := range list {
				if _, ok := mp.data[v]; !ok {
					mp.data[v] = 0
				}
				mp.data[v] += 1
			}
		}()
	}
	wg.Wait()
	fmt.Println(mp.data)
}
