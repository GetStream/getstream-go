# Channels

This guide shows how to migrate channel operations from `github.com/GetStream/stream-chat-go/v8` to `github.com/GetStream/getstream-go/v4`.

## Create a Channel

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

	resp, err := client.CreateChannel(ctx, "messaging", "general", "user-123",
		&stream.ChannelRequest{
			Members: []string{"user-123", "user-456"},
		})
	// resp.Channel
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

	ch := client.Chat().Channel("messaging", "general")
	resp, err := ch.GetOrCreate(ctx, &getstream.GetOrCreateChannelRequest{
		Data: &getstream.ChannelInput{
			CreatedByID: getstream.PtrTo("user-123"),
			Members: []getstream.ChannelMemberRequest{
				{UserID: "user-123"},
				{UserID: "user-456"},
			},
		},
	})
	// resp.Data.Channel
}
```

**Key changes:**
- `CreateChannel(type, id, creatorID, request)` becomes `client.Chat().Channel(type, id).GetOrCreate(ctx, request)`
- Members are `[]ChannelMemberRequest` objects instead of `[]string`
- Creator is set via `CreatedByID` field on `ChannelInput`

## Create a Distinct (1:1) Channel

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

	// Empty ID creates a distinct channel based on members
	resp, err := client.CreateChannel(ctx, "messaging", "", "alice",
		&stream.ChannelRequest{
			Members: []string{"alice", "bob"},
		})
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

	resp, err := client.Chat().GetOrCreateDistinctChannel(ctx, "messaging",
		&getstream.GetOrCreateDistinctChannelRequest{
			Data: &getstream.ChannelInput{
				CreatedByID: getstream.PtrTo("alice"),
				Members: []getstream.ChannelMemberRequest{
					{UserID: "alice"},
					{UserID: "bob"},
				},
			},
		})
	// resp.Data.Channel
}
```

**Key changes:**
- Distinct channels use a dedicated `GetOrCreateDistinctChannel` method instead of passing an empty ID

## Query Channels

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

	resp, err := client.QueryChannels(ctx, &stream.QueryOption{
		Filter: map[string]interface{}{
			"members": map[string]interface{}{"$in": []string{"user-123"}},
		},
		Limit:  10,
		Offset: 0,
	}, &stream.SortOption{Field: "created_at", Direction: -1})
	// resp.Channels is []*Channel
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

	resp, err := client.Chat().QueryChannels(ctx, &getstream.QueryChannelsRequest{
		FilterConditions: map[string]any{
			"members": map[string]any{"$in": []string{"user-123"}},
		},
	})
	// resp.Data.Channels
}
```

**Key changes:**
- Called on `client.Chat()` sub-client
- `Filter` becomes `FilterConditions` directly on the request (no wrapper `QueryOption`)
- Sort is a separate field on the request instead of a separate argument

## Add and Remove Members

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
	ch := client.Channel("messaging", "general")

	// Add members
	resp, err := ch.AddMembers(ctx, []string{"user-789", "user-012"})

	// Remove members
	resp, err = ch.RemoveMembers(ctx, []string{"user-012"}, nil)
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
	ch := client.Chat().Channel("messaging", "general")

	// Add members
	resp, err := ch.Update(ctx, &getstream.UpdateChannelRequest{
		AddMembers: []getstream.ChannelMemberRequest{
			{UserID: "user-789"},
			{UserID: "user-012"},
		},
	})

	// Remove members
	resp, err = ch.Update(ctx, &getstream.UpdateChannelRequest{
		RemoveMembers: []string{"user-012"},
	})
}
```

**Key changes:**
- No separate `AddMembers`/`RemoveMembers` methods; use `ch.Update()` with `AddMembers`/`RemoveMembers` fields
- `AddMembers` takes `[]ChannelMemberRequest` instead of `[]string`

## Update Channel

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
	ch := client.Channel("messaging", "general")

	// Full update
	resp, err := ch.Update(ctx, map[string]interface{}{
		"color": "blue",
	}, &stream.Message{Text: "Channel updated"})
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
	ch := client.Chat().Channel("messaging", "general")

	// Full update
	resp, err := ch.Update(ctx, &getstream.UpdateChannelRequest{
		Data: &getstream.ChannelInputRequest{
			Custom: map[string]any{"color": "blue"},
		},
		Message: &getstream.MessageRequest{
			Text:   getstream.PtrTo("Channel updated"),
			UserID: getstream.PtrTo("admin-user"),
		},
	})
}
```

**Key changes:**
- Update takes a single `UpdateChannelRequest` struct instead of separate map and message arguments
- Custom data goes under `Data.Custom`
- System message is a `MessageRequest` with pointer fields

## Partial Update Channel

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
	ch := client.Channel("messaging", "general")

	resp, err := ch.PartialUpdate(ctx, stream.PartialUpdate{
		Set:   map[string]interface{}{"name": "New Name"},
		Unset: []string{"description"},
	})
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
	ch := client.Chat().Channel("messaging", "general")

	resp, err := ch.UpdateChannelPartial(ctx, &getstream.UpdateChannelPartialRequest{
		Set:   map[string]any{"name": "New Name"},
		Unset: []string{"description"},
	})
}
```

**Key changes:**
- `PartialUpdate` becomes `UpdateChannelPartial`
- Same `Set`/`Unset` pattern, wrapped in `UpdateChannelPartialRequest`

## Delete Channel

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
	ch := client.Channel("messaging", "general")

	// Soft delete
	resp, err := ch.Delete(ctx)
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
	ch := client.Chat().Channel("messaging", "general")

	// Soft delete
	resp, err := ch.Delete(ctx, &getstream.DeleteChannelRequest{})

	// Hard delete
	resp, err = ch.Delete(ctx, &getstream.DeleteChannelRequest{
		HardDelete: getstream.PtrTo(true),
	})
}
```

**Key changes:**
- `Delete` requires a `DeleteChannelRequest` argument (can be empty for soft delete)
- Hard delete is a field on the request instead of separate options

## Batch Delete Channels

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

	cids := []string{"messaging:ch1", "messaging:ch2"}
	resp, err := client.DeleteChannels(ctx, cids, true) // hardDelete = true
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

	resp, err := client.Chat().DeleteChannels(ctx, &getstream.DeleteChannelsRequest{
		Cids:       []string{"messaging:ch1", "messaging:ch2"},
		HardDelete: getstream.PtrTo(true),
	})
	// resp.Data.TaskID for async task tracking
}
```

**Key changes:**
- Called on `client.Chat()` sub-client
- CIDs and hard delete wrapped in `DeleteChannelsRequest`

## Query Members

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
	ch := client.Channel("messaging", "general")

	resp, err := ch.QueryMembers(ctx, &stream.QueryOption{
		Filter: map[string]interface{}{},
		Limit:  10,
	})
	// resp.Members is []*ChannelMember
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

	resp, err := client.Chat().QueryMembers(ctx, &getstream.QueryMembersRequest{
		Payload: &getstream.QueryMembersPayload{
			Type:             "messaging",
			ID:               getstream.PtrTo("general"),
			FilterConditions: map[string]any{},
		},
	})
	// resp.Data.Members
}
```

**Key changes:**
- Called on `client.Chat()` instead of the channel object
- Channel type and ID are fields on the `QueryMembersPayload` instead of being implicit from the channel object
