package netpkg

import (
	"fmt"
	"net"
)

func TCPServer() error {
	listenAddr := fmt.Sprintf("%s:%d", "127.0.0.1", 8080)
	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	conn, err := ln.Accept()
	fmt.Println(conn.RemoteAddr().String(), "come in")
	return nil
}
