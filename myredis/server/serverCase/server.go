package serverCase

//  RedisServer 定义服务端结构体,里面存储着服务端的所有信息

type RedisServer struct {
	ConfigFile string //配置文件的位置
	Port       uint16
	IP         string
}
