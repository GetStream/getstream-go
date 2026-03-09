# Migrating from stream-chat-go to getstream-go

## Why Migrate?

- `getstream-go` is the actively developed, long-term-supported SDK
- Covers Chat, Video, Moderation, and Feeds in a single package
- Strongly typed models generated from the official OpenAPI spec
- `stream-chat-go` will enter maintenance mode (critical fixes only)

## Key Differences

| Aspect | stream-chat-go | getstream-go |
|--------|----------------|--------------|
| Import | `github.com/GetStream/stream-chat-go/v8` | `github.com/GetStream/getstream-go/v4` |
| Env vars | `STREAM_KEY`, `STREAM_SECRET` | `STREAM_API_KEY`, `STREAM_API_SECRET` |
| Client init | `stream.NewClient(key, secret)` | `getstream.NewClient(key, secret)` |
| Sub-clients | All methods on root client | `client.Chat()`, `client.Moderation()`, `client.Video()` |
| Channel object | `client.Channel(type, id)` | `client.Chat().Channel(type, id)` |
| Models | Flat structs with `ExtraData` | Generated request/response types with `Custom` maps |
| Optional fields | Value types or functional options | Pointers via `getstream.PtrTo()` |
| Members | `[]string` of user IDs | `[]ChannelMemberRequest` structs |

## Quick Example

**Before:**

```go
package main

import (
	"context"

	stream "github.com/GetStream/stream-chat-go/v8"
)

func main() {
	client, _ := stream.NewClient("key", "secret")
	ctx := context.Background()

	ch := client.Channel("messaging", "general")
	resp, err := ch.SendMessage(ctx, &stream.Message{
		Text: "Hello!",
	}, "user-123")
}
```

**After:**

```go
package main

import (
	"context"

	"github.com/GetStream/getstream-go/v4"
)

func main() {
	client, _ := getstream.NewClient("key", "secret")
	ctx := context.Background()

	ch := client.Chat().Channel("messaging", "general")
	resp, err := ch.SendMessage(ctx, &getstream.SendMessageRequest{
		Message: getstream.MessageRequest{
			Text:   getstream.PtrTo("Hello!"),
			UserID: getstream.PtrTo("user-123"),
		},
	})
}
```

## Migration Guides by Topic

| # | Topic | File |
|---|-------|------|
| 1 | [Setup and Authentication](01-setup-and-auth.md) | Client init, tokens |
| 2 | [Users](02-users.md) | Upsert, query, update, delete |
| 3 | [Channels](03-channels.md) | Create, query, members, update |
| 4 | [Messages and Reactions](04-messages-and-reactions.md) | Send, reply, react |
| 5 | [Moderation](05-moderation.md) | Ban, mute, moderators |
| 6 | [Devices](06-devices.md) | Push device management |

## Notes

- `stream-chat-go` is not going away. Your existing integration will keep working.
- The new SDK uses typed request structs with pointer fields for optional values.
- Use `getstream.PtrTo(value)` to convert any value to a pointer for optional fields.
- If you find a use case missing from this guide, please open an issue.
