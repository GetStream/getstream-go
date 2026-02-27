# Migration Guide: v3 → v4

This guide covers all breaking changes when upgrading from `getstream-go` v3 to v4.

## Overview

v4 is a full OpenAPI-aligned release. The primary change is a **systematic type renaming**: types that appear in API responses now have a `Response` suffix, and input types have a `Request` suffix. There are no removed features — all functionality from v3 is available in v4. Additionally, v4 adds complete coverage of the **Feeds**, **Video**, and **Moderation** product APIs.

## Installation

Update your `go.mod` to use the v4 module path and run `go get`:

```bash
go get github.com/GetStream/getstream-go/v4
```

Update all import paths in your code:

```go
// Before
import "github.com/GetStream/getstream-go/v3"

// After
import "github.com/GetStream/getstream-go/v4"
```

## Naming Conventions

All types use `PascalCase`. The general rules:

- Types returned in API responses: `Foo` → `FooResponse`
- Types used as API inputs: `Foo` → `FooRequest`
- Some moderation action payloads: `FooRequest` → `FooRequestPayload`

## Breaking Changes

### Common / Shared Types

| v3 | v4 | Notes |
| --- | --- | --- |
| `ApplicationConfig` | `AppResponseFields` | App configuration in responses |
| `ChannelPushPreferences` | `ChannelPushPreferencesResponse` | Per-channel push settings |
| `Device` | `DeviceResponse` | Device data (push, voip) |
| `Event` | `WSEvent` | WebSocket event envelope |
| `FeedsPreferences` | `FeedsPreferencesResponse` | Feeds push preferences |
| `ImportV2Task` | `ImportV2TaskItem` | V2 import task |
| `OwnUser` | `OwnUserResponse` | Authenticated user data |
| `Pager` | `PagerResponse` | Now cursor-based (`next`/`prev`) |
| `PushPreferences` | `PushPreferencesResponse` | Push preferences |
| `PushTemplate` | `PushTemplateResponse` | Push template |
| `PrivacySettings` | `PrivacySettingsResponse` | Typing indicators, read receipts |
| `RateLimitInfo` | `LimitInfoResponse` | Rate limit info |
| `SortParam` | `SortParamRequest` | Sort parameter for queries |
| `User` | `UserResponse` | Full user in responses |
| `UserBlock` | `BlockedUserResponse` | Blocked user details |
| `UserCustomEvent` | `CustomEvent` | Custom user event |
| `UserMute` | `UserMuteResponse` | User mute details |

### Event System

The event system was restructured from monolithic presets to composable types:

| Before (v3) | After (v4) | Notes |
| --- | --- | --- |
| `BaseEvent` (field `T`) | `Event` (field `type`) | Base event type |
| `Event` (WS envelope) | `WSEvent` | WebSocket event wrapper |
| `*Preset` embeds | `Has*` composition types | e.g., `HasChannel`, `HasMessage` |
| — | `WHEvent` | New webhook envelope type |

New composition types: `HasOwnUser`, `HasUserCommonFields`, `HasUserPrivacyFields`, `HasOptionalUserCommonFields`, `HasChannel`, `HasOptionalChannel`, `HasMessage`, `HasOptionalMessage`, `HasThreadParticipants`, `HasChannelTypeAndID`.

### Chat Types

| v3 | v4 | Notes |
| --- | --- | --- |
| `Campaign` | `CampaignResponse` | Campaign in responses |
| `CampaignStats` | `CampaignStatsResponse` | Campaign statistics |
| `Channel` | `ChannelResponse` | Channel in responses |
| `ChannelConfigFields` | `ChannelConfigWithInfo` | Channel config + commands/grants |
| `ChannelMember` | `ChannelMemberResponse` | Channel member in responses |
| `ChannelTypeConfigWithInfo` | `ChannelTypeConfig` | Channel type configuration |
| `ConfigOverrides` | `ConfigOverridesRequest` | Channel-level config overrides |
| `DraftMessage` / `DraftMessagePayload` | `DraftResponse` | Two types merged into one |
| `Message` | `MessageResponse` | Message in responses |
| `MessageReminder` | `ReminderResponseData` | Reminder in responses |
| `PendingMessage` | `PendingMessageResponse` | Pending message data |
| `Poll` | `PollResponse` | Poll in responses |
| `PollOption` | `PollOptionResponse` | Poll option in responses |
| `PollVote` | `PollVoteResponse` | Poll vote in responses |
| `Reaction` | `ReactionResponse` | Reaction in responses |
| `ReadState` | `ReadStateResponse` | Read state in responses |
| `Thread` | `ThreadResponse` | Thread in responses |

### Video Types

| v3 | v4 | Notes |
| --- | --- | --- |
| `AudioSettings` | `AudioSettingsResponse` | |
| `BackstageSettings` | `BackstageSettingsResponse` | |
| `BroadcastSettings` | `BroadcastSettingsResponse` | |
| `Call` | `CallResponse` | |
| `CallEgress` | `EgressResponse` | |
| `CallMember` | `MemberResponse` | Note: not `CallMemberResponse` |
| `CallParticipant` | `CallParticipantResponse` | |
| `CallParticipantFeedback` | *(removed)* | Use `CollectUserFeedbackRequest` |
| `CallSession` | `CallSessionResponse` | |
| `CallSettings` | `CallSettingsResponse` | |
| `CallType` | `CallTypeResponse` | |
| `EventNotificationSettings` | `EventNotificationSettingsResponse` | |
| `FrameRecordSettings` | `FrameRecordingSettingsResponse` | `Recording` inserted in name |
| `GeofenceSettings` | `GeofenceSettingsResponse` | |
| `HLSSettings` | `HLSSettingsResponse` | |
| `IndividualRecordSettings` | `IndividualRecordingSettingsResponse` | `Recording` inserted in name |
| `IngressSettings` | `IngressSettingsResponse` | |
| `IngressSource` | `IngressSourceResponse` | |
| `IngressAudioEncodingOptions` | `IngressAudioEncodingResponse` | Shortened name |
| `IngressVideoEncodingOptions` | `IngressVideoEncodingResponse` | Shortened name |
| `IngressVideoLayer` | `IngressVideoLayerResponse` | |
| `LimitsSettings` | `LimitsSettingsResponse` | |
| `NotificationSettings` | `NotificationSettingsResponse` | |
| `RawRecordSettings` | `RawRecordingSettingsResponse` | `Recording` inserted in name |
| `RecordSettings` | `RecordSettingsResponse` | |
| `RingSettings` | `RingSettingsResponse` | |
| `RTMPSettings` | `RTMPSettingsResponse` | |
| `ScreensharingSettings` | `ScreensharingSettingsResponse` | |
| `SessionSettings` | `SessionSettingsResponse` | |
| `SIPCallConfigs` | `SIPCallConfigsResponse` | |
| `SIPCallerConfigs` | `SIPCallerConfigsResponse` | |
| `SIPDirectRoutingRuleCallConfigs` | `SIPDirectRoutingRuleCallConfigsResponse` | |
| `SIPInboundRoutingRules` | `SIPInboundRoutingRuleResponse` | Plural → singular |
| `SIPPinProtectionConfigs` | `SIPPinProtectionConfigsResponse` | |
| `SIPTrunk` | `SIPTrunkResponse` | |
| `ThumbnailsSettings` | `ThumbnailsSettingsResponse` | |
| `TranscriptionSettings` | `TranscriptionSettingsResponse` | |
| `VideoSettings` | `VideoSettingsResponse` | |

### Moderation Types

| v3 | v4 | Notes |
| --- | --- | --- |
| `ActionLog` | `ActionLogResponse` | |
| `Appeal` | `AppealItemResponse` | |
| `AutomodDetails` | `AutomodDetailsResponse` | |
| `Ban` | `BanInfoResponse` | |
| `BanOptions` | *(removed)* | Merged into `BanActionRequestPayload` |
| `BanActionRequest` | `BanActionRequestPayload` | `Payload` suffix added |
| `BlockActionRequest` | `BlockActionRequestPayload` | |
| `BlockedMessage` | *(removed)* | Internal only |
| `CustomActionRequest` | `CustomActionRequestPayload` | |
| `DeleteMessageRequest` | `DeleteMessageRequestPayload` | |
| `DeleteUserRequest` | `DeleteUserRequestPayload` | |
| `EntityCreator` | `EntityCreatorResponse` | |
| `Evaluation` | `EvaluationResponse` | |
| `FeedsModerationTemplate` | `QueryFeedModerationTemplate` | No `Response` suffix |
| `FeedsModerationTemplateConfig` | `FeedsModerationTemplateConfigPayload` | |
| `Flag` | *(removed)* | Use `ModerationFlagResponse` |
| `Flag2` | `ModerationFlagResponse` | |
| `FlagDetails` | `FlagDetailsResponse` | |
| `FlagFeedback` | `FlagFeedbackResponse` | |
| `FlagMessageDetails` | `FlagMessageDetailsResponse` | |
| `FlagReport` | *(removed)* | Internal only |
| `FutureChannelBan` | `FutureChannelBanResponse` | |
| `MarkReviewedRequest` | `MarkReviewedRequestPayload` | |
| `Match` | `MatchResponse` | |
| `ModerationActionConfig` | `ModerationActionConfigResponse` | |
| `ModerationAnalytics` | *(removed)* | Internal only |
| `ModerationAnalyticsOverview` | *(removed)* | Internal only |
| `ModerationBulkSubmitActionRequest` | `BulkSubmitActionRequest` | `Moderation` prefix dropped |
| `ModerationConfig` | `ConfigResponse` | |
| `ModerationFlags` | *(removed)* | Use `[]*ModerationFlagResponse` |
| `ModerationLog` | *(removed)* | Use `ActionLogResponse` |
| `ModerationLogResponse` | *(removed)* | Use `QueryModerationLogsResponse` |
| `ModerationUsageStats` | `ModerationUsageStatsResponse` | |
| `RestoreActionRequest` | `RestoreActionRequestPayload` | |
| `ReviewQueueItem` | `ReviewQueueItemResponse` | |
| `Rule` | `RuleResponse` | |
| `ShadowBlockActionRequest` | `ShadowBlockActionRequestPayload` | |
| `Task` | `TaskResponse` | |
| `Trigger` | `TriggerResponse` | |
| `UnbanActionRequest` | `UnbanActionRequestPayload` | |
| `UnblockActionRequest` | `UnblockActionRequestPayload` | |
| `VideoEndCallRequest` | `VideoEndCallRequestPayload` | |
| `VideoKickUserRequest` | `VideoKickUserRequestPayload` | |

### Feeds Types

| v3 | v4 | Notes |
| --- | --- | --- |
| `Activity` | `ActivityResponse` | |
| `ActivityFeedback` | `ActivityFeedbackRequest` | Request-only type (no `Response` suffix) |
| `ActivityMark` | `MarkActivityRequest` | |
| `ActivityPin` | `ActivityPinResponse` | |
| `AggregatedActivity` | `AggregatedActivityResponse` | |
| `Bookmark` | `BookmarkResponse` | |
| `BookmarkFolder` | `BookmarkFolderResponse` | |
| `Collection` | `CollectionResponse` | |
| `Comment` | `CommentResponse` | |
| `CommentMedia` | *(removed)* | Embedded inline in `CommentResponse` |
| `CommentMention` | *(removed)* | Embedded inline in `CommentResponse` |
| `DenormalizedFeedsReaction` | *(removed)* | Internal only |
| `Feed` | `FeedResponse` | |
| `FeedGroup` | `FeedGroupResponse` | |
| `FeedMember` | `FeedMemberResponse` | |
| `FeedsReaction` | `FeedsReactionResponse` | |
| `FeedsReactionGroup` | `FeedsReactionGroupResponse` | |
| `FeedSuggestion` | `FeedSuggestionResponse` | |
| `FeedView` | `FeedViewResponse` | |
| `FeedVisibilityInfo` | `FeedVisibilityResponse` | |
| `Follow` | `FollowResponse` | |
| `MembershipLevel` | `MembershipLevelResponse` | |
| `ThreadedComment` | `ThreadedCommentResponse` | |

## Getting Help

- [Stream documentation](https://getstream.io/docs/)
- [GitHub Issues](https://github.com/GetStream/getstream-go/issues)
- [Stream support](https://getstream.io/contact/support/)
