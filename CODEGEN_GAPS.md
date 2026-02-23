# Codegen Gaps: APIs in stream-chat-go missing from getstream-go

These APIs exist in the hand-written `stream-chat-go` SDK but are not available in the auto-generated `getstream-go` SDK. Each entry references the stream-chat-go test file and the tests that cannot be ported.

## 1. BatchUpdateChannels

- **stream-chat-go file**: `channel_batch_updater_test.go`
- **Tests**: `TestClient_UpdateChannelsBatch`, `TestChannelBatchUpdater_AddMembers`, `TestChannelBatchUpdater_RemoveMembers`, `TestChannelBatchUpdater_Archive`
- **Description**: Batch update multiple channels in a single request (add/remove members, archive).

## 2. SendFile / SendImage

- **stream-chat-go file**: `channel_test.go`
- **Tests**: `TestChannel_SendFile`, `TestChannel_SendImage`
- **Description**: Upload files and images to a channel via multipart form upload, and delete them.

## 3. Drafts

- **stream-chat-go file**: `draft_test.go`
- **Tests**: `TestChannel_CreateDraft`, `TestChannel_GetDraft`, `TestChannel_DeleteDraft`, `TestChannel_CreateDraftInThread`, `TestClient_QueryDrafts`, `TestClient_QueryDraftsWithFilters`
- **Description**: Full CRUD for draft messages, including thread drafts and querying with filters/sort/pagination.

## 4. QueryFlagReportsAndReview

- **stream-chat-go file**: `query_test.go`
- **Tests**: `TestClient_QueryFlagReportsAndReview`
- **Description**: Query flag reports on messages and submit review actions on them.

## 5. QueryFutureChannelBans

- **stream-chat-go file**: `ban_test.go`
- **Tests**: `TestQueryFutureChannelBans`
- **Description**: Query channel bans that are scheduled for the future with filter conditions.

## 6. QueryTeamUsageStats

- **stream-chat-go file**: `team_usage_stats_test.go`
- **Tests**: `TestQueryTeamUsageStats_BasicAPI` (11 subtests), `TestQueryTeamUsageStats_Integration` (10 subtests), `TestQueryTeamUsageStats_DataCorrectness` (4 subtests)
- **Description**: Query usage statistics per team, with filtering by month/date range and pagination support.

## 7. CreatePermission

- **stream-chat-go file**: `permission_client_test.go`
- **Tests**: `TestPermissions_PermissionEndpoints` (CreatePermission subtest)
- **Description**: Create custom permissions (ListPermissions and GetPermission are available, but CreatePermission is not).
