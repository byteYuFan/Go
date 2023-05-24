package main

import (
	"context"
	"fmt"
	"time"
)

// Cancel the function that will be testing the Context.WithCancel
func Cancel() {
	// 创建一个带有取消功能的上下文
	ctx, cancel := context.WithCancel(context.Background())

	// 启动一个 goroutine 来执行定时任务
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("定时任务被取消")
				return
			default:
				// 模拟定时任务的工作
				fmt.Println("执行定时任务...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// 模拟等待一段时间后取消定时任务
	time.Sleep(5 * time.Second)
	cancel()

	// 等待一段时间以观察任务的状态
	time.Sleep(2 * time.Second)
}

// Dead the function that will be testing the context.WithDeadline
func Dead() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))

	// 启动一个 goroutine 来执行任务
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("任务被取消：", ctx.Err())
				return
			default:
				// 模拟任务的工作
				fmt.Println("执行任务...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// 等待一段时间以观察任务的状态
	time.Sleep(8 * time.Second)

	// 取消任务
	cancel()

	// 等待一段时间以观察任务的状态
	time.Sleep(2 * time.Second)
}

// TimeOut the function that will be testing the context.WithTimeout
func TimeOut() {
	// 设置一个超时时间为5秒的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// 启动一个 goroutine 来执行任务
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("任务被取消：", ctx.Err())
				return
			default:
				// 模拟任务的工作
				fmt.Println("执行任务...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// 等待一段时间以观察任务的状态
	time.Sleep(8 * time.Second)

	// 取消任务
	cancel()

	// 等待一段时间以观察任务的状态
	time.Sleep(2 * time.Second)
}

// Value the function that will be testing the context.WithValue
func Value() {
	parent := context.Background()

	// 使用 WithValue 创建一个带有键值对数据的子级上下文
	ctx := context.WithValue(parent, "userID", 123)

	// 在子级上下文中获取键值对数据
	userID := ctx.Value("userID")

	fmt.Println("用户ID:", userID)
}

// User the simple struct that using for test context.WithValue
type User struct {
	ID   int
	Name string
}

func Test() {
	// 创建一个父级上下文
	parent := context.Background()

	// 使用 WithValue 创建一个带有用户身份信息的子级上下文
	user := User{ID: 123, Name: "Alice"}
	ctx := context.WithValue(parent, "user", user)

	// 在不同的函数中获取用户身份信息
	processRequest(ctx)
}

// processRequest a function that get information from ctx
func processRequest(ctx context.Context) {
	// 从上下文中获取用户身份信息
	user, ok := ctx.Value("user").(User)
	if !ok {
		fmt.Println("无法获取用户身份信息")
		return
	}

	// 使用用户身份信息执行请求处理
	fmt.Printf("处理请求，用户ID: %d, 用户名: %s\n", user.ID, user.Name)

	// 调用其他函数传递上下文
	otherFunction(ctx)
}

// otherFunction another function that get information form ctx
func otherFunction(ctx context.Context) {
	// 从上下文中获取用户身份信息
	user, ok := ctx.Value("user").(User)
	if !ok {
		fmt.Println("无法获取用户身份信息")
		return
	}

	// 使用用户身份信息执行其他操作
	fmt.Printf("执行其他操作，用户ID: %d, 用户名: %s\n", user.ID, user.Name)
}
