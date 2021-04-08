package main

import (
	"net"
	"strconv"
	"tcpsocket"
	"time"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp4", "104.168.174.201:8848")
	if err != nil {
		panic(err)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		panic(err)
	}
	var socket tcpSocket.TcpSocket
	socket.Conn = conn
	go socket.ReadMsg() //处理接收消息
	for i := 0; i < 10; i++ {
		str := "message " + strconv.Itoa(i)
		socket.WriteMsg([]byte(str))
	}
	time.Sleep(10 * time.Second)
}
