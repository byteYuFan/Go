package main

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib2(x int) int {
	if x < 2 {
		return x
	}
	return fib2(x-1) + fib2(x-2)
}
func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib2(n)
	fmt.Printf("\rF(%d)=%d\n", n, fibN)
}
