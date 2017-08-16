package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
)

func sender(conn net.Conn) {
	r := bufio.NewReader(os.Stdin) /*使用bufio缓冲器*/
	for{
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)
		conn.Write([]byte(line))
		fmt.Println("send over")
	}
}

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:1024")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("connect success")
	sender(conn)
}