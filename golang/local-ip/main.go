package main

import (
	"log"
	"net"
)

// GetLocalIP can be get local ip
func GetLocalIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = conn.Close(); err != nil {
			log.Println("Close conn err:", err)
		}
	}()
	return conn.LocalAddr().(*net.UDPAddr).IP, nil
}

func main() {
	GetLocalIP()
}
