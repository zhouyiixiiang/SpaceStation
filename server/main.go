package main

import (
	"fmt"
	"net"
	"tcpsocket"
)

func main() {
	var socket tcpSocket.TcpSocket
	ip, err := socket.ExternalIP()
	if err != nil {
		fmt.Println("conn err")
	}
	addr := ip.String() + ":8848"
	server, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer server.Close()
	for {
		conn, err := server.Accept()
		if err != nil {
			panic(err)
		}
		socket.Conn = conn
		socket.ChanMsg = make(chan []byte, 100)
		go socket.ReadMsg()
		go readMsg(socket)
	}
}

func readMsg(socket tcpSocket.TcpSocket) {
	for {
		msg := <-socket.ChanMsg
		fmt.Println("read: ", string(msg))
		socket.Conn.Write(msg)
	}
}
