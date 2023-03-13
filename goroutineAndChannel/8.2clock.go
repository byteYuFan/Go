package main

import (
	"io"
	"log"
	"net"
	"time"
)

//Clock1 is a TCP server that periodically writes the time

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8090")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("1504:05\n"))
		if err != nil {
			return
		}

		time.Sleep(1 * time.Second)
	}
}