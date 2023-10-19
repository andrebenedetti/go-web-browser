package http_client

import (
	"errors"
	"strings"
	"unicode/utf8"
)

// ** NOTICE **
// Of course, the Go Standard Library implements full url parsing in conformation
// to the WHATWG spec https://url.spec.whatwg.org/#urls.
// Since the goal of this project is learning and teaching as much as possible, we're doing
// our own implementation of an url parser here.

// We're following the basic URL parser spec https://url.spec.whatwg.org/#concept-basic-url-parser

var ErrMalformedUrl = errors.New("malformed url string")

type UrlRecord struct {
	scheme   string
	username string
	password string
	host     string
	port     uint16
	path     []string
	query    string
	fragment string
	// blobUrlEntry
}

// think if we need to separate Scheme from SpecialScheme
type Scheme struct {
	scheme      string
	defaultPort int
}

var ErrInvalidUrlUnit = errors.New("invalid_url_unit")

func validateUrl(url string) error {
	// SPEC
	// If input contains any leading or trailing C0 control or space, invalid-URL-unit validation error.
	firstChar, _ := utf8.DecodeRune([]byte(url))
	lastChar, _ := utf8.DecodeLastRune([]byte(url))
	if firstChar <= 0x001F || lastChar <= 0x001F {
		return ErrInvalidUrlUnit
	}

	// If input contains any ASCII tab or newline, invalid-URL-unit validation error.
	for _, c := range url {
		if c == 0x9 || c == 0xA || c == 0xD {
			return ErrInvalidUrlUnit
		}
	}

	// WIP now parse scheme (item 4 of https://url.spec.whatwg.org/#concept-basic-url-parser)

	return nil
}

// ://
// len 2  value ["",""]

// asdasdas://
// len2, value [asdasdas, ""]

// Assumes url is a valid domain
// TODO: actually follow the spec when validating an url
func parseUrlScheme(url string) (string, error) {
	split := strings.Split(url, "://")
	if len(split) > 2 || split[0] == "" {
		return "", ErrMalformedUrl
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

type parsedUrl struct {
	scheme string
	host   string
	path   string
}

// func parseUrl(rawUrl string) (parsedUrl, error) {
// 	// parsed := parsedUrl{}
// 	scheme, error := parseUrlScheme(rawUrl)

// 	return parsedUrl{}, nil
// }
