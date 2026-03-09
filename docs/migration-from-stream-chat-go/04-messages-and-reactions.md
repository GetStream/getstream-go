# Messages and Reactions

This guide shows how to migrate message and reaction operations from `github.com/GetStream/stream-chat-go/v8` to `github.com/GetStream/getstream-go/v4`.

## Send a Message

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

	resp, err := ch.SendMessage(ctx, &stream.Message{
		Text: "Hello world",
	}, "user-123")
	// resp.Message
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

	resp, err := ch.SendMessage(ctx, &getstream.SendMessageRequest{
		Message: getstream.MessageRequest{
			Text:   getstream.PtrTo("Hello world"),
			UserID: getstream.PtrTo("user-123"),
		},
	})
	// resp.Data.Message
}
```

**Key changes:**
- User ID moves from a separate argument into `MessageRequest.UserID`
- Message content is wrapped in `SendMessageRequest` > `MessageRequest`
- `Text` is a pointer field; use `getstream.PtrTo()`
- Response accessed via `resp.Data.Message`

## Send a Thread Reply

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

	resp, err := ch.SendMessage(ctx, &stream.Message{
		Text:     "This is a reply",
		ParentID: "parent-message-id",
	}, "user-123")
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

	resp, err := ch.SendMessage(ctx, &getstream.SendMessageRequest{
		Message: getstream.MessageRequest{
			Text:     getstream.PtrTo("This is a reply"),
			ParentID: getstream.PtrTo("parent-message-id"),
			UserID:   getstream.PtrTo("user-123"),
		},
	})
}
```

**Key changes:**
- Same pattern as regular messages; `ParentID` is a pointer field on `MessageRequest`

## Get a Message

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

	resp, err := client.GetMessage(ctx, "message-id")
	// resp.Message
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

	resp, err := client.Chat().GetMessage(ctx, "message-id", &getstream.GetMessageRequest{})
	// resp.Data.Message
}
```

**Key changes:**
- Called on `client.Chat()` sub-client
- Requires a `GetMessageRequest` argument (can be empty)

## Update a Message

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

	resp, err := client.UpdateMessage(ctx, &stream.Message{
		Text: "Updated text",
	}, "message-id")
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

	resp, err := client.Chat().UpdateMessage(ctx, "message-id", &getstream.UpdateMessageRequest{
		Message: getstream.MessageRequest{
			Text:   getstream.PtrTo("Updated text"),
			UserID: getstream.PtrTo("user-123"),
		},
	})
}
```

**Key changes:**
- Message ID is a positional argument, not part of the message struct
- Message content wrapped in `UpdateMessageRequest` > `MessageRequest`

## Partial Update a Message

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

	resp, err := client.PartialUpdateMessage(ctx, "message-id", &stream.MessagePartialUpdateRequest{
		PartialUpdate: stream.PartialUpdate{
			Set: map[string]interface{}{
				"text":   "New text",
				"pinned": true,
			},
			Unset: []string{"attachments"},
		},
		UserID: "user-123",
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

	resp, err := client.Chat().UpdateMessagePartial(ctx, "message-id",
		&getstream.UpdateMessagePartialRequest{
			Set: map[string]any{
				"priority": "high",
				"status":   "reviewed",
			},
			Unset:  []string{"old_field"},
			UserID: getstream.PtrTo("user-123"),
		})
}
```

**Key changes:**
- `PartialUpdateMessage` becomes `UpdateMessagePartial`
- `Set`/`Unset` are directly on the request (no nested `PartialUpdate` struct)
- `UserID` is a pointer

## Delete a Message

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

	// Soft delete
	resp, err := client.DeleteMessage(ctx, "message-id")

	// Hard delete
	resp, err = client.DeleteMessage(ctx, "message-id", stream.DeleteMessageWithHard())
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

	// Soft delete
	resp, err := client.Chat().DeleteMessage(ctx, "message-id", &getstream.DeleteMessageRequest{})

	// Hard delete
	resp, err = client.Chat().DeleteMessage(ctx, "message-id", &getstream.DeleteMessageRequest{
		Hard: getstream.PtrTo(true),
	})
}
```

**Key changes:**
- Called on `client.Chat()` sub-client
- Functional options replaced by fields on `DeleteMessageRequest`

## Send a Reaction

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

	resp, err := client.SendReaction(ctx, &stream.Reaction{
		Type: "like",
	}, "message-id", "user-123")
	// resp.Reaction
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

	resp, err := client.Chat().SendReaction(ctx, "message-id", &getstream.SendReactionRequest{
		Reaction: getstream.ReactionRequest{
			Type:   "like",
			UserID: getstream.PtrTo("user-123"),
		},
	})
	// resp.Data.Reaction
}
```

**Key changes:**
- Called on `client.Chat()` sub-client
- Message ID is a positional argument; user ID moves into `ReactionRequest.UserID`
- Reaction wrapped in `SendReactionRequest` > `ReactionRequest`

## List Reactions

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

	resp, err := client.GetReactions(ctx, "message-id", map[string][]string{
		"limit":  {"10"},
		"offset": {"0"},
	})
	// resp.Reactions
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

	resp, err := client.Chat().GetReactions(ctx, "message-id", &getstream.GetReactionsRequest{
		Limit:  getstream.PtrTo(10),
		Offset: getstream.PtrTo(0),
	})
	// resp.Data.Reactions
}
```

**Key changes:**
- Called on `client.Chat()` sub-client
- Query params replaced by typed fields on `GetReactionsRequest`

## Delete a Reaction

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

	resp, err := client.DeleteReaction(ctx, "message-id", "like", "user-123")
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

	resp, err := client.Chat().DeleteReaction(ctx, "message-id", "like",
		&getstream.DeleteReactionRequest{
			UserID: getstream.PtrTo("user-123"),
		})
}
```

**Key changes:**
- Called on `client.Chat()` sub-client
- User ID moves from a positional argument into `DeleteReactionRequest.UserID`
