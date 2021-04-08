package tcpSocket

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
)

type TcpSocket struct {
	ChanMsg chan []byte
	Conn    net.Conn
}

func (item *TcpSocket) ReadMsg() {
	defer item.Conn.Close()
	for {
		bufferHead := make([]byte, 8)
		n, err := item.Conn.Read(bufferHead)
		if err != nil || n != 8 {
			fmt.Println("连接关闭", item.Conn)
			return
		}
		length := item.BytesToInt(bufferHead)
		bufferBody := make([]byte, length)
		n, err = item.Conn.Read(bufferBody)
		if err != nil || n != length {
			fmt.Println("连接关闭", item.Conn)
			return
		}
		//fmt.Println("收到：", string(bufferBody))
		item.ChanMsg <- bufferBody
	}
}
func (item *TcpSocket) WriteMsg(msg []byte) {
	length := len(msg)
	btsLen := item.IntToBytes(length)
	item.Conn.Write(btsLen)
	item.Conn.Write(msg)
}
func (item *TcpSocket) BytesToInt(bts []byte) int {
	byteBuffer := bytes.NewBuffer(bts)
	var data int64
	binary.Read(byteBuffer, binary.BigEndian, &data)
	return int(data)
}
func (item *TcpSocket) IntToBytes(n int) []byte {
	data := int64(n)
	byteBuffer := bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer, binary.BigEndian, data)
	return byteBuffer.Bytes()
}

func (item *TcpSocket) ExternalIP() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip, nil
		}
	}
	return nil, errors.New("connected to the network?")
}
func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}

	return ip
}
