package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":4545")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}

func print(s []byte) string {
	return strings.Replace(string(s), "\n", "\\n", -1)
}

func handleConnection(conn net.Conn) {
	fmt.Printf("received a conn at %v from %v\n", conn.LocalAddr(), conn.RemoteAddr())
	buffer := make([]byte, 100)
	conn.Write([]byte("POOP"))
	conn.Read(buffer)
	conn.Write([]byte("POOP"))
	fmt.Printf("◤%v◢\n", print(buffer))
}
