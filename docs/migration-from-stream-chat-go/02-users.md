# Users

This guide shows how to migrate user operations from `github.com/GetStream/stream-chat-go/v8` to `github.com/GetStream/getstream-go/v4`.

## Upsert a Single User

**Before (stream-chat-go):**

```go
package main

import (
	"context"

	stream "github.com/GetStream/stream-chat-go/v8"
)

func main() {
	client, _ := stream.NewClient("your-api-key", "your-api-secret")
	ctx := context.Background()

	resp, err := client.UpsertUser(ctx, &stream.User{
		ID:    "user-123",
		Name:  "Alice",
		Image: "https://example.com/alice.jpg",
		Role:  "user",
		ExtraData: map[string]interface{}{
			"country": "NL",
		},
	})
	// resp.User
}
```

**After (getstream-go):**

```go
package main

import (
	"context"

	"github.com/GetStream/getstream-go/v4"
)

func main() {
	client, _ := getstream.NewClient("your-api-key", "your-api-secret")
	ctx := context.Background()

	resp, err := client.UpdateUsers(ctx, &getstream.UpdateUsersRequest{
		Users: map[string]getstream.UserRequest{
			"user-123": {
				ID:    "user-123",
				Name:  getstream.PtrTo("Alice"),
				Image: getstream.PtrTo("https://example.com/alice.jpg"),
				Role:  getstream.PtrTo("user"),
				Custom: map[string]any{
					"country": "NL",
				},
			},
		},
	})
	// resp.Data.Users["user-123"]
}
```

**Key changes:**
- `UpsertUser` becomes `UpdateUsers` with a map of `UserRequest` keyed by user ID
- String fields like `Name`, `Image`, `Role` are pointers; use `getstream.PtrTo()` to set them
- `ExtraData` becomes `Custom`
- Response accessed via `resp.Data.Users[id]` instead of `resp.User`

## Upsert Multiple Users

**Before (stream-chat-go):**

```go
package main

import (
	"context"

	stream "github.com/GetStream/stream-chat-go/v8"
)

func main() {
	client, _ := stream.NewClient("your-api-key", "your-api-secret")
	ctx := context.Background()

	resp, err := client.UpsertUsers(ctx,
		&stream.User{ID: "user-1", Name: "Alice"},
		&stream.User{ID: "user-2", Name: "Bob"},
	)
	// resp.Users is map[string]*User
}
```

**After (getstream-go):**

```go
package main

import (
	"context"

	"github.com/GetStream/getstream-go/v4"
)

func main() {
	client, _ := getstream.NewClient("your-api-key", "your-api-secret")
	ctx := context.Background()

	resp, err := client.UpdateUsers(ctx, &getstream.UpdateUsersRequest{
		Users: map[string]getstream.UserRequest{
			"user-1": {ID: "user-1", Name: getstream.PtrTo("Alice")},
			"user-2": {ID: "user-2", Name: getstream.PtrTo("Bob")},
		},
	})
	// resp.Data.Users is map[string]UserResponse
}
```

**Key changes:**
- `UpsertUsers` (variadic) becomes `UpdateUsers` with a map of `UserRequest`
- Same endpoint, different calling convention

## Query Users

**Before (stream-chat-go):**

```go
package main

import (
	"context"

	stream "github.com/GetStream/stream-chat-go/v8"
)

func main() {
	client, _ := stream.NewClient("your-api-key", "your-api-secret")
	ctx := context.Background()

	resp, err := client.QueryUsers(ctx, &stream.QueryUsersOptions{
		QueryOption: stream.QueryOption{
			Filter: map[string]interface{}{
				"role": map[string]string{"$eq": "admin"},
			},
			Limit:  10,
			Offset: 0,
		},
	}, &stream.SortOption{Field: "created_at", Direction: -1})
	// resp.Users is []*User
}
```

**After (getstream-go):**

```go
package main

import (
	"context"

	"github.com/GetStream/getstream-go/v4"
)

func main() {
	client, _ := getstream.NewClient("your-api-key", "your-api-secret")
	ctx := context.Background()

	resp, err := client.QueryUsers(ctx, &getstream.QueryUsersRequest{
		Payload: &getstream.QueryUsersPayload{
			FilterConditions: map[string]any{
				"role": map[string]any{"$eq": "admin"},
			},
			Limit:  getstream.PtrTo(10),
			Offset: getstream.PtrTo(0),
		},
	})
	// resp.Data.Users
}
```

**Key changes:**
- Filter and pagination are wrapped in a `Payload` field of type `QueryUsersPayload`
- `Filter` becomes `FilterConditions`
- `Limit`/`Offset` are pointers
- Sort is not a separate argument; it goes inside the payload if needed

## Partial Update User

**Before (stream-chat-go):**

```go
package main

import (
	"context"

	stream "github.com/GetStream/stream-chat-go/v8"
)

func main() {
	client, _ := stream.NewClient("your-api-key", "your-api-secret")
	ctx := context.Background()

	user, err := client.PartialUpdateUser(ctx, stream.PartialUserUpdate{
		ID: "user-123",
		Set: map[string]interface{}{
			"name":    "Alice Updated",
			"country": "US",
		},
		Unset: []string{"image"},
	})
	// user is *User
}
```

**After (getstream-go):**

```go
package main

import (
	"context"

	"github.com/GetStream/getstream-go/v4"
)

func main() {
	client, _ := getstream.NewClient("your-api-key", "your-api-secret")
	ctx := context.Background()

	resp, err := client.UpdateUsersPartial(ctx, &getstream.UpdateUsersPartialRequest{
		Users: []getstream.UpdateUserPartialRequest{
			{
				ID: "user-123",
				Set: map[string]any{
					"name":    "Alice Updated",
					"country": "US",
				},
				Unset: []string{"image"},
			},
		},
	})
	// resp.Data.Users["user-123"]
}
```

**Key changes:**
- `PartialUpdateUser` becomes `UpdateUsersPartial` with a slice of `UpdateUserPartialRequest`
- Always a batch operation; single user is a slice of one
- Returns `resp.Data.Users` map instead of a single `*User`

## Deactivate User

**Before (stream-chat-go):**

```go
package main

import (
	"context"

	stream "github.com/GetStream/stream-chat-go/v8"
)

func main() {
	client, _ := stream.NewClient("your-api-key", "your-api-secret")
	ctx := context.Background()

	resp, err := client.DeactivateUser(ctx, "user-123",
		stream.DeactivateUserWithMarkMessagesDeleted())
}
```

**After (getstream-go):**

```go
package main

import (
	"context"

	"github.com/GetStream/getstream-go/v4"
)

func main() {
	client, _ := getstream.NewClient("your-api-key", "your-api-secret")
	ctx := context.Background()

	resp, err := client.DeactivateUser(ctx, "user-123", &getstream.DeactivateUserRequest{
		MarkMessagesDeleted: getstream.PtrTo(true),
	})
}
```

**Key changes:**
- Functional options like `DeactivateUserWithMarkMessagesDeleted()` become fields on a `DeactivateUserRequest` struct
- User ID is still a positional argument

## Delete Users

**Before (stream-chat-go):**

```go
package main

import (
	"context"

	stream "github.com/GetStream/stream-chat-go/v8"
)

func main() {
	client, _ := stream.NewClient("your-api-key", "your-api-secret")
	ctx := context.Background()

	resp, err := client.DeleteUser(ctx, "user-123",
		stream.DeleteUserWithHardDelete(),
		stream.DeleteUserWithMarkMessagesDeleted())
}
```

**After (getstream-go):**

```go
package main

import (
	"context"

	"github.com/GetStream/getstream-go/v4"
)

func main() {
	client, _ := getstream.NewClient("your-api-key", "your-api-secret")
	ctx := context.Background()

	resp, err := client.DeleteUsers(ctx, &getstream.DeleteUsersRequest{
		UserIds:  []string{"user-123"},
		User:     getstream.PtrTo("hard"),
		Messages: getstream.PtrTo("hard"),
	})
	// resp.Data.TaskID for async task tracking
}
```

**Key changes:**
- `DeleteUser` (single, sync) becomes `DeleteUsers` (batch, async)
- Options expressed as string fields (`"hard"` or `"soft"`) instead of functional options
- Returns a `TaskID` for polling completion
