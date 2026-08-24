# Moderation

This guide shows how to migrate moderation operations from `github.com/GetStream/stream-chat-go/v8` to `github.com/GetStream/getstream-go/v5`.

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

	"github.com/GetStream/getstream-go/v5"
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

	"github.com/GetStream/getstream-go/v5"
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

	"github.com/GetStream/getstream-go/v5"
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

	"github.com/GetStream/getstream-go/v5"
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

	"github.com/GetStream/getstream-go/v5"
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

	"github.com/GetStream/getstream-go/v5"
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

	"github.com/GetStream/getstream-go/v5"
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

	"github.com/GetStream/getstream-go/v5"
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

## Block a User

User blocking is a per-user block list, separate from banning.

**Before (stream-chat-go):**

```go
resp, err := client.BlockUser(ctx, "target-user", "acting-user")

resp, err = client.UnblockUser(ctx, "target-user", "acting-user")

blocked, err := client.GetBlockedUser(ctx, "acting-user")
```

**After (getstream-go):**

```go
resp, err := client.BlockUsers(ctx, &getstream.BlockUsersRequest{
	BlockedUserID: "target-user",
	UserID:        getstream.PtrTo("acting-user"),
})

resp, err = client.UnblockUsers(ctx, &getstream.UnblockUsersRequest{
	BlockedUserID: "target-user",
	UserID:        getstream.PtrTo("acting-user"),
})

blocked, err := client.GetBlockedUsers(ctx, &getstream.GetBlockedUsersRequest{
	UserID: getstream.PtrTo("acting-user"),
})
```

**Key changes:**
- Method names are plural: `BlockUser` becomes `BlockUsers`, `GetBlockedUser` becomes `GetBlockedUsers`
- Positional arguments become fields on a request struct
- Do not confuse these with `client.Video().UnblockUser(...)`, which removes a user from a call

## Blocklists

**Before (stream-chat-go):**

```go
err := client.CreateBlocklist(ctx, &stream.BlocklistCreateRequest{
	BlocklistBase: stream.BlocklistBase{
		Name:  "profanity",
		Words: []string{"badword"},
	},
})

list, err := client.GetBlocklist(ctx, "profanity")
_, err = client.UpdateBlocklist(ctx, "profanity", []string{"badword", "worse"})
all, err := client.ListBlocklists(ctx)
_, err = client.DeleteBlocklist(ctx, "profanity")
```

**After (getstream-go):**

```go
resp, err := client.CreateBlockList(ctx, &getstream.CreateBlockListRequest{
	Name:  "profanity",
	Words: []string{"badword"},
})

list, err := client.GetBlockList(ctx, "profanity", &getstream.GetBlockListRequest{})
_, err = client.UpdateBlockList(ctx, "profanity", &getstream.UpdateBlockListRequest{
	Words: []string{"badword", "worse"},
})
all, err := client.ListBlockLists(ctx, &getstream.ListBlockListsRequest{})
_, err = client.DeleteBlockList(ctx, "profanity", &getstream.DeleteBlockListRequest{})
```

**Key changes:**
- The type is spelled `BlockList` (capital L), not `Blocklist`. Every method renames accordingly: `CreateBlocklist` to `CreateBlockList`, `ListBlocklists` to `ListBlockLists`, and so on
- Every method takes a request struct, even where the legacy call took only a name
- `UpdateBlockList` carries `Words` on the request instead of a positional slice
- The new request supports matching options the legacy SDK did not expose: `Type` (`word`, `regex`, `domain`, `email`, and the allowlist variants), `IsSubstringMatchingEnabled`, `IsLeetCheckEnabled`, `IsPluralCheckEnabled`, `IsConfusableFoldingEnabled`

## Flagging Content

> **Read this before migrating flags.** The legacy flag methods write to the **v1 chat flags** store. `getstream-go` exposes the **v2 moderation** API, and the two stores are not the same: content flagged through `Moderation().Flag()` may not appear in `Chat().QueryMessageFlags()`, which still reads v1. Swapping the call alone can therefore leave a flagging workflow that writes to one store and reads from another. Migrate the whole workflow (flag, query, review) to v2 together, or keep using v1 until you can.

**Before (stream-chat-go):**

```go
_, err := client.FlagMessage(ctx, "message-id", "acting-user")
_, err = client.FlagUser(ctx, "target-user", "acting-user")

flags, err := client.QueryMessageFlags(ctx, &stream.QueryOption{
	Filter: map[string]interface{}{"channel_cid": "messaging:general"},
})
```

**After (getstream-go):**

```go
_, err := client.Moderation().Flag(ctx, &getstream.FlagRequest{
	EntityType: "message",
	EntityID:   "message-id",
	UserID:     getstream.PtrTo("acting-user"),
})

_, err = client.Moderation().Flag(ctx, &getstream.FlagRequest{
	EntityType: "user",
	EntityID:   "target-user",
	UserID:     getstream.PtrTo("acting-user"),
})

flags, err := client.Chat().QueryMessageFlags(ctx, &getstream.QueryMessageFlagsRequest{
	Payload: &getstream.QueryMessageFlagsPayload{
		FilterConditions: map[string]any{"channel_cid": "messaging:general"},
	},
})
```

**Key changes:**
- `FlagMessage` and `FlagUser` collapse into one `Moderation().Flag()` call; the target is described by `EntityType` plus `EntityID` instead of a dedicated method
- `QueryMessageFlags` stays on the Chat sub-client and takes its filter under `Payload`
- Flag reads and writes can cross the v1/v2 boundary described above; verify your flags are visible where you expect before relying on them

## Reviewing Flagged Content

The v1 flag-report workflow has no drop-in replacement. `QueryFlagReports` and `ReviewFlagReport` are replaced by the v2 **review queue**, which is a different model rather than a rename, so this needs rework instead of a call swap.

**Before (stream-chat-go):**

```go
reports, err := client.QueryFlagReports(ctx, &stream.QueryFlagReportsRequest{})
_, err = client.ReviewFlagReport(ctx, "report-id", &stream.ReviewFlagReportRequest{
	ReviewResult: "reviewed",
})
```

**After (getstream-go):**

```go
queue, err := client.Moderation().QueryReviewQueue(ctx, &getstream.QueryReviewQueueRequest{})

item, err := client.Moderation().GetReviewQueueItem(ctx, "item-id", &getstream.GetReviewQueueItemRequest{})

_, err = client.Moderation().SubmitAction(ctx, &getstream.SubmitActionRequest{
	ItemID:     getstream.PtrTo("item-id"),
	ActionType: "mark_reviewed",
})
```

**Key changes:**
- Flag reports become review-queue items: query with `QueryReviewQueue`, read one with `GetReviewQueueItem`
- Reviewing is an action submitted against an item via `SubmitAction` rather than a single review call
- Because the underlying model changed, treat this as a redesign of the moderation workflow and confirm the behavior you need against the moderation documentation

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
| `client.BlockUser(ctx, target, by)` | `client.BlockUsers(ctx, &BlockUsersRequest{...})` |
| `client.UnblockUser(ctx, target, by)` | `client.UnblockUsers(ctx, &UnblockUsersRequest{...})` |
| `client.GetBlockedUser(ctx, by)` | `client.GetBlockedUsers(ctx, &GetBlockedUsersRequest{...})` |
| `client.CreateBlocklist(ctx, req)` | `client.CreateBlockList(ctx, &CreateBlockListRequest{...})` |
| `client.GetBlocklist(ctx, name)` | `client.GetBlockList(ctx, name, &GetBlockListRequest{})` |
| `client.UpdateBlocklist(ctx, name, words)` | `client.UpdateBlockList(ctx, name, &UpdateBlockListRequest{Words: words})` |
| `client.ListBlocklists(ctx)` | `client.ListBlockLists(ctx, &ListBlockListsRequest{})` |
| `client.DeleteBlocklist(ctx, name)` | `client.DeleteBlockList(ctx, name, &DeleteBlockListRequest{})` |
| `client.FlagMessage(ctx, msgID, by)` | `client.Moderation().Flag(ctx, &FlagRequest{EntityType: "message", ...})` (v1 to v2, see note) |
| `client.FlagUser(ctx, target, by)` | `client.Moderation().Flag(ctx, &FlagRequest{EntityType: "user", ...})` (v1 to v2, see note) |
| `client.QueryMessageFlags(ctx, q)` | `client.Chat().QueryMessageFlags(ctx, &QueryMessageFlagsRequest{...})` |
| `client.QueryFlagReports(ctx, req)` | `client.Moderation().QueryReviewQueue(ctx, ...)` (different model) |
| `client.ReviewFlagReport(ctx, id, req)` | `client.Moderation().SubmitAction(ctx, ...)` (different model) |
