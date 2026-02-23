# Codegen Gaps: APIs in stream-chat-go missing from getstream-go

These APIs exist in the hand-written `stream-chat-go` SDK but are not available in the auto-generated `getstream-go` SDK. Each gap was verified by inspecting the backend controllers — all are intentionally excluded from the OpenAPI spec via `Ignore: true`, `Beta: true`, or `ClientSideOnly()`.

Re-generating the client from the latest backend spec (`make openapi` + `generate-client --language go-serverside`) produces identical API methods — no new endpoints appear.

## Intentionally excluded (Ignore: true / Beta)

### 1. BatchUpdateChannels

- **stream-chat-go file**: `channel_batch_updater_test.go`
- **Tests**: `TestClient_UpdateChannelsBatch`, `TestChannelBatchUpdater_AddMembers`, `TestChannelBatchUpdater_RemoveMembers`, `TestChannelBatchUpdater_Archive`
- **Description**: Batch update multiple channels in a single request (add/remove members, archive).
- **Backend**: `controllers/chat/v1/channel_batch_update.go` — `Ignore: true`, `Beta: true`

### 2. QueryFlagReports / ReviewFlagReport

- **stream-chat-go file**: `query_test.go`
- **Tests**: `TestClient_QueryFlagReportsAndReview`
- **Description**: Query flag reports on messages and submit review actions on them.
- **Backend**: `controllers/chat/v1/moderation_report.go` — `Ignore: true` on both `QueryFlagReports` and `ReviewFlagReport` controllers

### 3. QueryTeamUsageStats

- **stream-chat-go file**: `team_usage_stats_test.go`
- **Tests**: `TestQueryTeamUsageStats_BasicAPI` (11 subtests), `TestQueryTeamUsageStats_Integration` (10 subtests), `TestQueryTeamUsageStats_DataCorrectness` (4 subtests)
- **Description**: Query usage statistics per team, with filtering by month/date range and pagination support.
- **Backend**: `controllers/chat/v1/query_team_usage_stats.go` — `Ignore: true`, `Beta: true`

### 4. CreatePermission

- **stream-chat-go file**: `permission_client_test.go`
- **Tests**: `TestPermissions_PermissionEndpoints` (CreatePermission subtest)
- **Description**: Create custom permissions (ListPermissions and GetPermission are available, but CreatePermission is not).
- **Backend**: `controllers/create_permission.go` — `Ignore: true`

## Client-side only (excluded from server-side SDKs by design)

### 5. CreateDraft

- **stream-chat-go file**: `draft_test.go`
- **Tests**: `TestChannel_CreateDraft`, `TestChannel_CreateDraftInThread`
- **Description**: Create/upsert draft messages in a channel.
- **Backend**: `controllers/chat/v1/create_draft.go` — `ClientSideOnly()`, so excluded from server-side SDK generation
- **Note**: `GetDraft`, `DeleteDraft`, and `QueryDrafts` ARE available in getstream-go.

## Previously misidentified as gaps (actually exist in getstream-go)

These were initially listed as gaps but are available in the SDK:

- **SendFile / SendImage** — Available as `client.Chat().UploadChannelFile()` and `client.Chat().UploadChannelImage()` (plus `client.UploadFile()` / `client.UploadImage()` for generic uploads)
- **QueryFutureChannelBans** — Available as `client.Chat().QueryFutureChannelBans()`
- **Drafts (Get/Delete/Query)** — Available as `ch.GetDraft()`, `ch.DeleteDraft()`, `client.Chat().QueryDrafts()`
