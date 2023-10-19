package http_client

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       string
}

func isValidHttpMethod(method string) bool {
	// var validHttpMethods = [9]string{"GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH"}
	// for _, v := range validHttpMethods {
	// 	if v == method {
	// 		return true
	// 	}
	// }

	// return false

	// Out of curiosity, the code block above was benchmarked against the one below.
	// The if chain is 30x faster in the best case (method == "GET"). The worst case also is roughly the same
	return method == "GET" || method == "HEAD" || method == "POST" || method == "PUT" || method == "DELETE" || method == "CONNECT" || method == "OPTIONS" || method == "TRACE" || method == "PATCH"
}

// Parse headers' keys and values into a map.
// Keys are stored in lowercase.
func parseHeaders(raw []string) map[string]string {
	headers := make(map[string]string, len(raw))
	for _, val := range raw {
		kv := strings.Split(val, ":")
		if len(kv) < 2 {
			log.Fatalf("Received malformed header %s", val)
		}

		key := strings.ToLower(kv[0])

		v := strings.Split(kv[1], "\r\n")
		value := strings.Trim(v[0], " ")

		headers[key] = value
	}

	return headers
}

func parseUrlScheme(url string) (string, error) {
	split := strings.Split(url, "://")
	if len(split) > 2 {
		return "", errors.New("malformed url string")
	}

	// Let's assume http if no scheme is set.
	// TODO: assume https
	if len(split) == 1 {
		return "http", nil
	}

	scheme := split[0]
	if scheme != "http" && scheme != "file" {
		return scheme, errors.New("unsupported url scheme")
	}

	return scheme, nil
}

func RetrieveUrl(url string) (Response, error) {
	scheme, err := parseUrlScheme(url)
	if err != nil {
		return Response{}, err
	}

	if scheme == "file" {
		// We only support localhost files, for simplicity and because we don't want to
		// deal with security implications of exposing file urls in this project
		// We will also allow a special host called go-web-browser which will refer to this project's
		// public directory.
		splitUrl := strings.Split(url, "file://")[1]
		host := strings.Split(splitUrl, "/")[0]
		path := strings.Split(splitUrl, host)[1]

		if host != "" && host != "localhost" && host != "go-web-browser" {
			return Response{}, errors.New("file:// scheme pointing to a host different than localhost is not currently supported")
		}
		return fileRequest(host, path)
	}
	return httpRequest("GET", url)

}

func fileRequest(host string, path string) (Response, error) {
	fmt.Println("File request", host, path)
	return Response{}, nil
}

func httpRequest(method string, url string) (Response, error) {
	// Let's copy what go's standard lib does and assume
	// an empty string as being a GET
	if method == "" {
		method = "GET"
	}

	var resp Response
	if !isValidHttpMethod(method) {
		return resp, errors.New("invalid HTTP method")
	}

	if method != "GET" {
		// Let's first support GET method.
		// TODO: support other methods
		return resp, errors.New("method not supported")
	}

	conn, err := net.Dial("tcp", url+":80")
	if err != nil {
		log.Fatalf("Error dialing url %s: %s", url, err)
	}

	fmt.Fprintf(conn, "%s /index.html HTTP/1.0\r\nHost: %s\r\n\r\n", method, url)
	reader := bufio.NewReader(conn)
	statusLine, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Request error: %s", err.Error())
	}

	statusCode, err := parseStatusCode(statusLine)
	if err != nil {
		log.Fatalf("Error parsing status line: %s", err.Error())
	}

	resp.StatusCode = statusCode

	// Read headers
	rawHeaders := make([]string, 0, 32)
	for {
		line, err := reader.ReadString('\n')
		// End of headers is marked by a \r\n line
		if line == "\r\n" {
			break
		}

		rawHeaders = append(rawHeaders, line)

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

	resp.Headers = parseHeaders(rawHeaders)
	if resp.Headers["content-encoding"] != "" || resp.Headers["transfer-encoding"] != "" {
		return resp, errors.New("request encoding not implemented")
	}

	// Using 1 MB buffer size to read the page, which is a lot for a web page
	// TODO: This can be revisited later
	buffer := make([]byte, 10000)
	for {
		_, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalf("Error reading body: %s", err.Error())
			}
		}
	}
	resp.Body = string(buffer)
	return resp, nil
}
