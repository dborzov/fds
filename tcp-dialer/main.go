package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "0.0.0.0:4545")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(conn, "HELLO BODY\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Printf("and the status is: %v", status)
}
