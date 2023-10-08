package http_client

import "testing"

func BenchmarkIsValidHttpHeader(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isValidHttpMethod("FLANDERFELTER")
	}
}
