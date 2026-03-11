# Setup and Authentication

This guide shows how to migrate setup and authentication code from `github.com/GetStream/stream-chat-go/v8` to `github.com/GetStream/getstream-go/v4`.

## Installation

**Before (stream-chat-go):**

```bash
go get github.com/GetStream/stream-chat-go/v8
```

**After (getstream-go):**

```bash
go get github.com/GetStream/getstream-go/v4
```

**Key changes:**
- Module path changed from `stream-chat-go/v8` to `getstream-go/v4`

## Client Initialization

**Before (stream-chat-go):**

```go
package main

import (
	"time"

	stream "github.com/GetStream/stream-chat-go/v8"
)

func main() {
	client, err := stream.NewClient("your-api-key", "your-api-secret")
	if err != nil {
		panic(err)
	}

	// Or with a custom timeout
	client, err = stream.NewClient("your-api-key", "your-api-secret",
		stream.WithTimeout(5*time.Second))
}
```

**After (getstream-go):**

```go
package main

import (
	"time"

	"github.com/GetStream/getstream-go/v4"
)

func main() {
	client, err := getstream.NewClient("your-api-key", "your-api-secret")
	if err != nil {
		panic(err)
	}

	// Or with a custom timeout
	client, err = getstream.NewClient("your-api-key", "your-api-secret",
		getstream.WithTimeout(5*time.Second))
}
```

**Key changes:**
- Import alias changes from `stream` to `getstream`
- Same `NewClient(key, secret)` signature, just a different package

## Client from Environment Variables

**Before (stream-chat-go):**

```go
package main

import (
	stream "github.com/GetStream/stream-chat-go/v8"
)

func main() {
	// Reads STREAM_KEY, STREAM_SECRET, STREAM_CHAT_TIMEOUT
	client, err := stream.NewClientFromEnvVars()
}
```

**After (getstream-go):**

```go
package main

import (
	"github.com/GetStream/getstream-go/v4"
)

func main() {
	// Reads STREAM_API_KEY, STREAM_API_SECRET
	client, err := getstream.NewClientFromEnvVars()
}
```

**Key changes:**
- Environment variable names changed: `STREAM_KEY` to `STREAM_API_KEY`, `STREAM_SECRET` to `STREAM_API_SECRET`

## Token Generation

**Before (stream-chat-go):**

```go
package main

import (
	"time"

	stream "github.com/GetStream/stream-chat-go/v8"
)

func main() {
	client, _ := stream.NewClient("your-api-key", "your-api-secret")

	// Token with expiration
	expiry := time.Now().Add(24 * time.Hour)
	token, err := client.CreateToken("user-123", expiry)
}
```

**After (getstream-go):**

```go
package main

import (
	"time"

	"github.com/GetStream/getstream-go/v4"
)

func main() {
	client, _ := getstream.NewClient("your-api-key", "your-api-secret")

	// Token with expiration
	token, err := client.CreateToken("user-123",
		getstream.WithExpiration(24*time.Hour))
}
```

**Key changes:**
- `CreateToken` uses functional options (`WithExpiration`) instead of positional `time.Time` arguments
- Expiration is specified as a duration, not an absolute time
