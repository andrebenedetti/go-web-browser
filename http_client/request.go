package http_client

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

type StatusCode struct {
}

type Response struct {
	statusCode int
}

func isValidHttpMethod(method string) bool {
	return method == "GET" || method == "HEAD" || method == "POST" || method == "PUT" || method == "DELETE" || method == "CONNECT" || method == "OPTIONS" || method == "TRACE" || method == "PATCH"
}

func Request(method string, url string) (Response, error) {
	// Let's copy what go's standard lib does and assume
	// an empty string as being a GET
	if method == "" {
		method = "GET"
	}

	var resp Response
	if !isValidHttpMethod(method) {
		return resp, errors.New("Invalid HTTP method")
	}

	if method != "GET" {
		// Let's first support GET method.
		// TODO: support other methods
		return resp, errors.New("Method not supported")
	}

	conn, err := net.Dial("tcp", url+":80")
	if err != nil {
		log.Fatalf("Error dialing url %s: %s", url, err)
	}

	fmt.Fprintf(conn, "GET /index.html HTTP/1.0\r\nHost: example.org\r\n\r\n")
	reader := bufio.NewReader(conn)
	statusLine, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Request error: %s", err.Error())
	}

	statusCode, err := parseStatusCode(statusLine)
	if err != nil {
		log.Fatalf("Error parsing status line: %s", err.Error())
	}

	resp.statusCode = statusCode

	// Read headers
	for {
		line, err := reader.ReadString('\n')
		// End of headers is marked by a \r\n line
		if line == "\r\n" {
			break
		}

		if err != nil {
			// io.EOF should never be wrapped, so we don't need to use errors.Is,
			// but be aware of that.
			if err == io.EOF {
				break
			} else {
				log.Fatalf("Request error: %s\n", err.Error())
			}
		}

	}

	return resp, nil
}
