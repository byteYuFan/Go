package _case

import (
	"fmt"
	"sync"
)

func MapCase() {
	mp := sync.Map{}
	mp.Store("name", "wyf")
	mp.Store("email", "854978151@qq.com")
	fmt.Println(mp.Load("name"))
	fmt.Println(mp.Load("email"))

	_, ok := mp.LoadOrStore("age", 18)
	fmt.Println(ok)
	fmt.Println(mp.LoadAndDelete("age"))

	mp.Range(visit)
}

func visit(key, value any) bool {
	fmt.Println(key, "-", value)
	return true
}
