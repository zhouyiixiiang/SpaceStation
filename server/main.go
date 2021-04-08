package main

import (
	"fmt"
	"net"
	"strings"
	"tcpsocket"
)

func main() {
	var socket tcpSocket.TcpSocket
	ip, err := socket.ExternalIP()
	if err != nil {
		fmt.Println("conn err")
	}
	//addr := ip.String() + ":8848"
	ip.String()
	addr := "127.0.0.1:8848"
	server, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer server.Close()
	for {
		var tmpsocket tcpSocket.TcpSocket
		conn, err := server.Accept()
		if err != nil {
			panic(err)
		}
		tmpsocket.Conn = conn
		tmpsocket.ChanMsg = make(chan []byte, 100)
		tcpSocket.SocketList=append(tcpSocket.SocketList,tmpsocket )
		go tmpsocket.ReadMsg()
		go readMsg(tmpsocket)
	}
}

func readMsg(socket tcpSocket.TcpSocket) {
	name:="some one"
	for {
		msg := <-socket.ChanMsg
		fmt.Println("read: ", string(msg))
		msgString:=string(msg)
		if strings.Contains(msgString,"name:"){
			idx:=strings.Index(msgString,"name:")
			name=msgString[idx+5:]
		}
		msgOut:=name+": "+string(msg)
		for _,item:=range tcpSocket.SocketList{
			if item!=socket{
				item.WriteMsg([]byte(msgOut))
			}
		}
	}
}
