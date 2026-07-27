# Migration Guide: v4 → v5

v5 is a small release. The import path changes, four response fields that the API no longer returns are gone, and `CheckResponse` can no longer be compared with `==`. Everything else is additive, so most codebases only need to update the import path.

## Start at v5.1.0

v5.1.0 is the first v5 release. v5.0.0 was withdrawn, and everything in it is included in v5.1.0.

## Installation

```bash
go get github.com/GetStream/getstream-go/v5
```

Update all import paths in your code:

```go
// Before
import "github.com/GetStream/getstream-go/v4"

// After
import "github.com/GetStream/getstream-go/v5"
```

## Breaking Changes

### Removed fields

These fields were dropped from the API spec and have no replacement. They were never populated in a way applications could rely on.

| Type | Field | JSON |
| --- | --- | --- |
| `AppResponseFields` | `ModerationAudioFileEnabled` | `moderation_audio_file_enabled` |
| `ModerationDashboardPreferences` | `AnalyzeMaxImageSizeBytes` | `analyze_max_image_size_bytes` |
| `ModerationDashboardPreferences` | `AnalyzeMaxKeyframeSizeBytes` | `analyze_max_keyframe_size_bytes` |
| `ModerationDashboardPreferences` | `WebhookHeaderClientRequestIDKey` | `webhook_header_client_request_id_key` |

### `CheckResponse` is no longer comparable

`CheckResponse` gained `TriggeredRules []TriggeredRuleResponse`, which lists every rule a moderation check triggered. A struct containing a slice is not comparable in Go, so `==` on it no longer compiles:

```go
// v4: compiled
if got == want { ... }

// v5: compile error
// invalid operation: got == want (struct containing []getstream.TriggeredRuleResponse cannot be compared)
if reflect.DeepEqual(got, want) { ... }
```

This propagates to anything embedding it, most commonly `StreamResponse[CheckResponse]`. Field-by-field comparison works too and is usually what you want in tests.

The singular `TriggeredRule *TriggeredRuleResponse` is still present and still populated with the first triggered rule, so existing reads keep working. Prefer `TriggeredRules` for new code: it is the only field that surfaces content, user, and call rules together.

## New in v5

- `ChatClient.GetChannel(ctx, type, id, request)` and `Channels.Get(ctx, request)`: fetch channel state without a `GetOrCreate` write.
- `CheckResponse.TriggeredRules`: all rules triggered by a moderation check, with their resolved actions.
- `ModerationCallResponse`: the moderation call payload has its own type. The Video `CallResponse` is unchanged.
- Feeds translation: `TranslateActivity`, `TranslateComment`, `I18n` on activity and comment responses, plus `Language` and `TranslateText` parameters on the feeds read endpoints.
- Activity shares: `QueryActivityShares`, `ShareResponse`, `FeedsShareResponse`.
- `WithRetry(RetryConfig{...})`: opt-in retry for rate-limited and transport-failed requests (v5.1.0).
- `WithLogBodies(true)` and structured client log events with secret redaction (v5.1.0).

## Verifying your upgrade

The removals and the comparability change are all compile-time failures, so the compiler finds every affected line for you:

```bash
go build ./... && go vet ./...
```

## Getting Help

- [Stream documentation](https://getstream.io/docs/)
- [GitHub Issues](https://github.com/GetStream/getstream-go/issues)
- [Stream support](https://getstream.io/contact/support/)
