package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8081")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("listening on 0.0.0.0:8081")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		if err := handle(conn); err != nil {
			log.Println(err)
		}
	}
}

func handle(conn net.Conn) error {
	defer conn.Close()

	requestLine, _ := bufio.NewReader(conn).ReadString('\n')
	fields := strings.Fields(requestLine)

	status, body := "404 NOT FOUND", "not found"
	if len(fields) >= 2 {
		switch fields[1] {
		case "/ping":
			status, body = "200 OK", "pong"
		case "/":
			status, body = "200 OK", "hello world from Go"
		}
	}

	response := fmt.Sprintf(
		"HTTP/1.1 %s\r\nContent-Length: %d\r\nContent-Type: text/plain\r\n\r\n%s",
		status, len(body), body,
	)
	_, err := conn.Write([]byte(response))
	return err
}
