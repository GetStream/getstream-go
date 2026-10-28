# Migrating from v7 to v8

v8 regenerates the OpenAPI client from protocol tag `openapi-v239.93.2` (chat **v239.93.2**, FLU-1105). **HTTP JSON is unchanged** — the server still accepts and returns the same string values. The Go SDK types many of those fields as **named enums** (`type X string`) instead of `string` / `*string`, so **v7.1.1 keeps working against the new server** until you choose to upgrade.

Upgrading the import path alone is not enough: expect **compile errors** anywhere you assign, compare, or pass those fields as plain strings.

## Module path

```bash
go get github.com/GetStream/getstream-go/v8@v8.0.0
```

```go
- import "github.com/GetStream/getstream-go/v7"
+ import "github.com/GetStream/getstream-go/v8"
```

Update every import of this module in your repo (including tests).

## Upgrade workflow

1. Bump the module path and run `go mod tidy`.
2. Run `go build ./...` (or your usual CI) and fix compile errors field by field.
3. Prefer **generated constants** over raw strings so renames and typos fail at compile time.
4. Run integration tests against your app; wire behavior should match v7 unless you relied on omitted-vs-empty JSON edge cases (see [Campaign sender fields](#campaign-sender-fields)).

Full API diff between releases: [v7.1.1...v8.0.0](https://github.com/GetStream/getstream-go/compare/v7.1.1...v8.0.0).

## Enum migration patterns

All new enums are **string-backed**. JSON still uses the same words (`"all"`, `"flag"`, `"firebase"`, …).

### Request literals

```go
// v7
created, err := client.Chat().CreateChannelType(ctx, &CreateChannelTypeRequest{
    Name:            "team",
    AutomodBehavior: "flag",
    PushLevel:       PtrTo("all"),
})

// v8 — use constants (names are generated from enum values)
created, err := client.Chat().CreateChannelType(ctx, &CreateChannelTypeRequest{
    Name:            "team",
    AutomodBehavior: ChannelModerationBehaviorFlag,
    PushLevel:       PtrTo(PushLevelAll),
})
```

When no constant exists or you have a dynamic value:

```go
level := PushLevel(userConfiguredLevel) // same string as on the wire
```

### Response reads and comparisons

```go
// v7
if resp.Data.PushLevel != nil && *resp.Data.PushLevel == "none" { ... }

// v8
if resp.Data.PushLevel != nil && *resp.Data.PushLevel == PushLevelNone { ... }
// or
if resp.Data.PushLevel != nil && resp.Data.PushLevel.String() == "none" { ... }
```

### Pointers

Many fields move from `*string` to `*SomeEnum`. Keep using `PtrTo`:

```go
VotingVisibility: PtrTo(VotingVisibilityAnonymous),
```

### Push diagnostics

`CheckPushRequest.PushProviderType` is `*PushProviderName` (not `*string`). Device and template types use `PushProviderName` where the spec names the provider.

```go
_, err := client.CheckPush(ctx, &CheckPushRequest{
    UserID:           PtrTo(userID),
    PushProviderType: PtrTo(PushProviderNameFirebase),
})
```

## Campaign sender fields

On `CampaignResponse`, `sender_mode` and `sender_visibility` are `*CampaignSenderMode` and `*CampaignSenderVisibility`, with **`omitempty`**. Default campaigns may **omit** those keys instead of sending `""`. Do not treat missing keys as empty string on read; use nil checks.

## New named enum types in v8

These `type … string` values were added or are now used on requests/responses (constants live next to each type in `models.go`):

`ActivitySelectorType`, `ActivityVisibility`, `AudioDefaultDevice`, `AutomodMode`, `BlocklistBehavior`, `CallLifecycleMode`, `CameraFacing`, `CampaignSenderMode`, `CampaignSenderVisibility`, `ChannelModerationBehavior`, `ChannelOwnCapability`, `CollectionStatus`, `CommentStatus`, `DeleteMessages`, `FeatureActivationMode`, `FeedMemberStatus`, `FeedOwnCapability`, `FeedVisibility`, `FollowPushPreference`, `FollowStatus`, `FriendReactionType`, `LayoutName`, `ModerationAction`, `OwnCapability`, `PushLevel`, `PushProviderName`, `RecordingType`, `RestrictReplies`, `RetentionPolicyType`, `ScheduledMessageFailureReason`, `ScheduledMessageQueryStatus`, `ScheduledMessageStatus`, `SetupStatus`, `SetupStep`, `SortParamType`, `TranscriptionLanguage`, `VideoCodec`, `VotingVisibility`.

Not every struct field uses every type yet; search your broken builds for the type name or run the [compare view](https://github.com/GetStream/getstream-go/compare/v7.1.1...v8.0.0) for the exact field you touch.

## Highlights by area

| Area | v7 | v8 |
|------|----|----|
| Channel type / metadata `push_level` | `*string` | `*PushLevel` |
| Channel type `automod_behavior` | `string` | `ChannelModerationBehavior` |
| Blocklist / config overrides `blocklist_behavior` | `*string` | `*BlocklistBehavior` |
| Feeds `visibility`, polls `voting_visibility` | `*string` | `*FeedVisibility`, `VotingVisibility`, … |
| Devices / push | `string` provider fields | `PushProviderName` where typed in spec |
| Campaigns | `string` sender fields | `*CampaignSenderMode`, `*CampaignSenderVisibility`, omitempty |

## When to stay on v7

Stay on **v7.1.1** if you are not ready to fix compile-time typing. The server release that shipped these spec names does **not** require v8 on the client.

## Changelog

See [CHANGELOG.md](./CHANGELOG.md) for the v8.0.0 release entry.
