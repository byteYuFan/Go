package tool

import (
	"bufio"
	"github.com/tedcy/fdfs_client"
	"log"
	"os"
	"strings"
)

// UploadFile 上传文件
func UploadFile(filename string) string {
	//新建一个客户端对象
	client, err := fdfs_client.NewClientWithConfig("./config/fastdfs.conf")
	if err != nil {
		log.Fatal(err.Error())
		return ""
	}
	defer client.Destory()
	fileID, err := client.UploadByFilename(filename)
	if err != nil {
		log.Fatal(err.Error())
		return ""
	}
	return fileID
}

func FileServerAddr() string {
	file, err := os.Open("./config/fastdfs.config")
	if err != nil {
		return ""
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		str := strings.SplitN(line, "=", 2)
		switch str[0] {
		case "http_server":
			return str[1]
		}
		if err != nil {
			return ""
		}
	}
}
