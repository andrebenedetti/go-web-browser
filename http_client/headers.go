package http_client

import (
	"log"
	"strconv"
	"strings"
)

// Take the status line (first line of the http response) and parse the
// status code. If the format does not conform to the specification, return
// an error.
func parseStatusCode(statusLine string) (int, error) {
	split := strings.Split(statusLine, " ")
	// We expect the format "HTTP1.0 200 OK"
	if len(split) != 3 {
		log.Fatalf("Malformed status header: %s", statusLine)
	}

	code, err := strconv.Atoi(split[1])
	if err != nil || code < 100 || code > 599 {
		log.Fatalf("Received invalid status code %d", code)
	}

	return code, nil
}
