package _case

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AtomicCase() {
	var count int64 = 5
	atomic.StoreInt64(&count, 10)
	fmt.Println("获取变量的值", atomic.LoadInt64(&count))
	fmt.Println("在原有的基础上累加10:", atomic.AddInt64(&count, 10))
	fmt.Println("交换并返回原有的值:", atomic.SwapInt64(&count, 50))
	fmt.Println("获取变量的值", atomic.LoadInt64(&count))
}

type atomicCounter struct {
	count int64
}

func (c *atomicCounter) Inc() {
	atomic.AddInt64(&c.count, 1)
}
func (c *atomicCounter) Load() int64 {
	return atomic.LoadInt64(&c.count)
}

func AtomicCase1() {
	var count int64
	wg := sync.WaitGroup{}
	locker := sync.Mutex{}
	for i := 1; i <= 200; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			locker.Lock()
			count += 1
			locker.Unlock()
		}()
	}
	wg.Wait()
	fmt.Printf("%d ", count)
}
func AtomicCase2() {
	var count = &atomicCounter{}
	wg := sync.WaitGroup{}
	for i := 1; i <= 200; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count.Inc()
		}()
	}
	wg.Wait()
	fmt.Printf("%d ", count.Load())
}
func AtomicCase3() {
	list := []string{"A", "B", "C", "D"}
	var atomicMp atomic.Value
	mp := map[string]int{}
	atomicMp.Store(&mp)
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
		label:
			m := atomicMp.Load().(*map[string]int)

			m1 := map[string]int{}
			for k, v := range *m {
				m1[k] = v
			}
			for _, item := range list {
				if _, ok := m1[item]; !ok {
					m1[item] = 0
				}
				m1[item] += 1
			}
			swap := atomicMp.CompareAndSwap(m, &m1)
			if !swap {
				goto label
			}
		}()
	}

	wg.Wait()
	fmt.Println(atomicMp.Load())
}
