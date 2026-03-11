# Moderation

This guide shows how to migrate moderation operations from `github.com/GetStream/stream-chat-go/v8` to `github.com/GetStream/getstream-go/v4`.

## Add Moderators

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

	resp, err := ch.AddModerators(ctx, "user-123", "user-456")
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

	resp, err := ch.Update(ctx, &getstream.UpdateChannelRequest{
		AddModerators: []string{"user-123", "user-456"},
	})
}
```

**Key changes:**
- No dedicated `AddModerators` method; use `ch.Update()` with `AddModerators` field
- `DemoteModerators` is also a field on `UpdateChannelRequest`

## Ban User (App Level)

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

	resp, err := client.BanUser(ctx, "target-user", "moderator-user",
		stream.BanWithReason("Spam"),
		stream.BanWithExpiration(60))
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

	resp, err := client.Moderation().Ban(ctx, &getstream.BanRequest{
		TargetUserID: "target-user",
		BannedByID:   getstream.PtrTo("moderator-user"),
		Reason:       getstream.PtrTo("Spam"),
		Timeout:      getstream.PtrTo(60),
	})
}
```

**Key changes:**
- Called on `client.Moderation()` sub-client instead of the root client
- Functional options replaced by fields on `BanRequest`
- `BanWithExpiration` becomes `Timeout`

## Ban User (Channel Level)

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

	resp, err := ch.BanUser(ctx, "target-user", "moderator-user",
		stream.BanWithReason("Inappropriate content"))
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

	resp, err := client.Moderation().Ban(ctx, &getstream.BanRequest{
		TargetUserID: "target-user",
		BannedByID:   getstream.PtrTo("moderator-user"),
		ChannelCid:   getstream.PtrTo("messaging:general"),
		Reason:       getstream.PtrTo("Inappropriate content"),
	})
}
```

**Key changes:**
- No channel-level `BanUser` method; use `client.Moderation().Ban()` with `ChannelCid` field
- Channel CID format is `type:id` (e.g., `messaging:general`)

## Unban User

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

	// App-level unban
	resp, err := client.UnBanUser(ctx, "target-user")

	// Channel-level unban
	ch := client.Channel("messaging", "general")
	resp, err = ch.UnBanUser(ctx, "target-user")
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

	// App-level unban
	resp, err := client.Moderation().Unban(ctx, &getstream.UnbanRequest{
		TargetUserID: "target-user",
	})

	// Channel-level unban
	resp, err = client.Moderation().Unban(ctx, &getstream.UnbanRequest{
		TargetUserID: "target-user",
		ChannelCid:   getstream.PtrTo("messaging:general"),
	})
}
```

**Key changes:**
- Both app-level and channel-level unbans use `client.Moderation().Unban()`
- Channel scope specified via `ChannelCid` field

## Shadow Ban

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

	// App-level shadow ban
	resp, err := client.ShadowBan(ctx, "target-user", "moderator-user")

	// Channel-level shadow ban
	ch := client.Channel("messaging", "general")
	resp, err = ch.ShadowBan(ctx, "target-user", "moderator-user")
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

	// App-level shadow ban
	resp, err := client.Moderation().Ban(ctx, &getstream.BanRequest{
		TargetUserID: "target-user",
		BannedByID:   getstream.PtrTo("moderator-user"),
		Shadow:       getstream.PtrTo(true),
	})

	// Channel-level shadow ban
	resp, err = client.Moderation().Ban(ctx, &getstream.BanRequest{
		TargetUserID: "target-user",
		BannedByID:   getstream.PtrTo("moderator-user"),
		ChannelCid:   getstream.PtrTo("messaging:general"),
		Shadow:       getstream.PtrTo(true),
	})
}
```

**Key changes:**
- No separate `ShadowBan` method; use `Ban()` with `Shadow: true`

## Query Banned Users

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

	resp, err := client.QueryBannedUsers(ctx, &stream.QueryBannedUsersOptions{
		QueryOption: stream.QueryOption{
			Filter: map[string]interface{}{
				"channel_cid": "messaging:general",
			},
			Limit: 10,
		},
	})
	// resp.Bans
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

	resp, err := client.Chat().QueryBannedUsers(ctx, &getstream.QueryBannedUsersRequest{
		Payload: &getstream.QueryBannedUsersPayload{
			FilterConditions: map[string]any{
				"channel_cid": "messaging:general",
			},
		},
	})
	// resp.Data.Bans
}
```

**Key changes:**
- Called on `client.Chat()` sub-client
- Filter wrapped in `Payload.FilterConditions`

## Mute User

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

	resp, err := client.MuteUser(ctx, "target-user", "moderator-user",
		stream.MuteWithExpiration(60))
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

	resp, err := client.Moderation().Mute(ctx, &getstream.MuteRequest{
		TargetIds: []string{"target-user"},
		UserID:    getstream.PtrTo("moderator-user"),
		Timeout:   getstream.PtrTo(60),
	})
}
```

**Key changes:**
- Called on `client.Moderation()` sub-client
- `MuteUser` becomes `Mute` with `MuteRequest`
- `TargetIds` is a slice, allowing batch muting
- Expiration specified via `Timeout` (minutes)

## Unmute User

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

	resp, err := client.UnmuteUser(ctx, "target-user", "moderator-user")
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

	resp, err := client.Moderation().Unmute(ctx, &getstream.UnmuteRequest{
		TargetIds: []string{"target-user"},
		UserID:    getstream.PtrTo("moderator-user"),
	})
}
```

**Key changes:**
- Called on `client.Moderation()` sub-client
- `UnmuteUser` becomes `Unmute` with `UnmuteRequest`
- `TargetIds` is a slice, allowing batch unmuting

## Method Mapping Summary

| Legacy (stream-chat-go) | New (getstream-go) |
|---|---|
| `ch.AddModerators(ctx, ids...)` | `ch.Update(ctx, &UpdateChannelRequest{AddModerators: ids})` |
| `ch.DemoteModerators(ctx, ids...)` | `ch.Update(ctx, &UpdateChannelRequest{DemoteModerators: ids})` |
| `client.BanUser(ctx, target, by, opts...)` | `client.Moderation().Ban(ctx, &BanRequest{...})` |
| `ch.BanUser(ctx, target, by, opts...)` | `client.Moderation().Ban(ctx, &BanRequest{ChannelCid: ...})` |
| `client.UnBanUser(ctx, target, opts...)` | `client.Moderation().Unban(ctx, &UnbanRequest{...})` |
| `client.ShadowBan(ctx, target, by)` | `client.Moderation().Ban(ctx, &BanRequest{Shadow: true})` |
| `client.MuteUser(ctx, target, by, opts...)` | `client.Moderation().Mute(ctx, &MuteRequest{...})` |
| `client.UnmuteUser(ctx, target, by)` | `client.Moderation().Unmute(ctx, &UnmuteRequest{...})` |
| `client.QueryBannedUsers(ctx, opts)` | `client.Chat().QueryBannedUsers(ctx, &QueryBannedUsersRequest{...})` |
