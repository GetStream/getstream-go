package getstream

import (
	"math/rand"
	"time"
)

func PtrTo[T any](v T) *T {
	return &v
}

func randomString(n int) string {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, n)
	for i := 0; i < n; i++ {
		bytes[i] = byte(65 + r.Intn(26)) // should be 26 to include 'Z'
	}
	return string(bytes)
}
