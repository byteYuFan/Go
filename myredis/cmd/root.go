package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var cfgFile string

// package cmd  存储着myRedis的启动命令
// root 为整个启动命令的跟节点
var version = "1.0.0"
var rootCmd = &cobra.Command{
	Use:   "myredis",
	Short: "A simple Redis clone written in Go",
	Long: `myredis is a simple Redis clone written in Go. 
        It provides basic functionality of Redis data structures, such as strings, lists, hashes, sets, and sorted sets.`,
	Version: version,
	Run:     rootRun,
}

func rootRun(cmd *cobra.Command, args []string) {
	fmt.Printf("Configuration file %s loaded successfully\n", cfgFile)

	// TODO: 实现启动 Redis 服务的代码0
	fmt.Println("Redis server started")

	// 进入 REPL 模式
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("127.0.0.1:6379>>> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		input = strings.TrimSpace(input)

		// 解析用户输入的命令，并执行相应的操作
		switch input {
		case "get":
			fmt.Println("Executing get command")
			// TODO: 实现 get 命令的代码
		case "set":
			fmt.Println("Executing set command")
			// TODO: 实现 set 命令的代码
		case "quit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Unknown command")
		}
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "filename", "f", "/usr/local/redis.conf", "The  path for the Redis configuration file")
}

func Exec() {
	rootCmd.Execute()
}
