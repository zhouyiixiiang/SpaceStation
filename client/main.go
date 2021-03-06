package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"tcpsocket"
	"time"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp4", "104.168.174.201:8848")
	//addr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8848")
	if err != nil {
		panic(err)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		panic(err)
	}
	var socket tcpSocket.TcpSocket
	socket.Conn = conn
	socket.ChanMsg = make(chan []byte, 100)
	go socket.ReadMsg() //处理接收消息
	go readMsg(socket)
	var msg string
	var inputReader *bufio.Reader
	for {
		time.Sleep(50)
		inputReader=bufio.NewReader(os.Stdin)
		msg, err = inputReader.ReadString('\n')
		socket.WriteMsg([]byte(msg))
	}
	time.Sleep(10 * time.Second)
}
func readMsg(socket tcpSocket.TcpSocket) {
	for {
		msg := <-socket.ChanMsg
		fmt.Println( string(msg))
	}
}
