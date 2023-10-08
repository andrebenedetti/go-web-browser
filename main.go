package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// We are not using net/http packages because the book goes lower, creating a socket
// and building the HTTP requests by itself and we want to do it here too for fun =).
func Visit(url string) {
	conn, err := net.Dial("tcp", url+":80")
	if err != nil {
		log.Fatalf("Error dialing url %s: %s", url, err)
	}

	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Printf("Status: %s\n", status)
}

func main() {
	Visit("google.com.br")
}
