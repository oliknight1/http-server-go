package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		req := make([]byte, 1024)
		conn.Read(req)

		reqStr := string(req)
		if strings.Contains(reqStr, "GET / HTTP/1.1") {
			conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))

		} else {

			conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
		}

	}

}
