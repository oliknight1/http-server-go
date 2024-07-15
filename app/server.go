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

	fmt.Println("Server Started")
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

		split_req := strings.Split(reqStr, "\r\n")
		request_line := split_req[0]

		headers := make(map[string]string)
		for _, l := range split_req[1 : len(split_req)-1] {
			if len(l) > 0 {
				header_key := strings.Split(l, " ")[0]
				// Remove the colon
				header_key = header_key[:len(header_key)-1]

				header_val := strings.Split(l, " ")[1]
				fmt.Println(header_val)
				headers[header_key] = header_val

			}
		}

		path := strings.Split(request_line, " ")[1]
		path_split := strings.Split(path, "/")

		return_val := ""

		if path == "/" {
			conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))

		} else if path_split[1] == "echo" {
			echo_val := strings.Split(path, "/")[2]
			return_val = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(echo_val), echo_val)

		} else if path_split[1] == "user-agent" {

			return_val = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(headers["User-Agent"]), headers["User-Agent"])

		} else {
			return_val = "HTTP/1.1 404 Not Found\r\n\r\n"
		}
		conn.Write([]byte(return_val))

		conn.Close()

	}

}
