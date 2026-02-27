# Changelog

All notable changes to this project will be documented in this file. See [standard-version](https://github.com/conventional-changelog/standard-version) for commit guidelines.

## [4.0.0-beta.1](https://github.com/GetStream/getstream-go/compare/v3.8.0...v4.0.0-beta.1) (2026-02-27)

### Breaking Changes

- Type names across all products now follow the OpenAPI spec naming convention: response types are suffixed with `Response`, input types with `Request`. See [MIGRATION_v3_to_v4.md](./MIGRATION_v3_to_v4.md) for the complete rename mapping.
- `Event` (WebSocket envelope type) renamed to `WSEvent`. Base event type renamed from `BaseEvent` to `Event` (with field `type` instead of `T`).
- Event composition changed from monolithic `*Preset` embeds to modular `Has*` types (`HasChannel`, `HasMessage`, `HasUserCommonFields`, etc.).
- `Pager` renamed to `PagerResponse` and migrated from offset-based to cursor-based pagination (`next`/`prev` tokens).
- Module path changed from `github.com/GetStream/getstream-go/v3` to `github.com/GetStream/getstream-go/v4`.

### Added

- Full product coverage: Chat, Video, Moderation, and Feeds APIs are all supported in a single SDK.
- **Feeds**: activities, feeds, feed groups, follows, comments, reactions, collections, bookmarks, membership levels, feed views, and more.
- **Video**: calls, recordings, transcription, closed captions, SFU, call statistics, user feedback analytics, and more.
- **Moderation**: flags, review queue, moderation rules, config, appeals, moderation logs, and more.
- Push notification types, preferences, and templates.
- Webhook support: `WHEvent` envelope type for receiving webhook payloads, utility functions for decoding and verifying webhook signatures, and a full set of individual typed event structs for every event across all products (Chat, Video, Moderation, Feeds) usable as discriminated event types.
- Cursor-based pagination across all list endpoints.

## [3.8.0](https://github.com/GetStream/getstream-go/compare/v3.7.0...v3.8.0) (2026-02-03)

## [3.7.0](https://github.com/GetStream/getstream-go/compare/v3.6.0...v3.7.0) (2025-11-28)

## [3.6.0](https://github.com/GetStream/getstream-go/compare/v3.5.0...v3.6.0) (2025-11-18)

## [3.5.0](https://github.com/GetStream/getstream-go/compare/v3.4.0...v3.5.0) (2025-11-13)

## [3.4.0](https://github.com/GetStream/getstream-go/compare/v3.3.0...v3.4.0) (2025-11-12)

## [3.3.0](https://github.com/GetStream/getstream-go/compare/v3.2.0...v3.3.0) (2025-10-10)

## [3.2.0](https://github.com/GetStream/getstream-go/compare/v3.1.2...v3.2.0) (2025-09-30)

### [3.1.2](https://github.com/GetStream/getstream-go/compare/v3.1.1...v3.1.2) (2025-09-19)

### [3.1.1](https://github.com/GetStream/getstream-go/compare/v3.1.0...v3.1.1) (2025-09-02)


### Features

* add SendClosedCaption + other updates ([#45](https://github.com/GetStream/getstream-go/issues/45)) ([21e8050](https://github.com/GetStream/getstream-go/commit/21e8050fbffe283fe50b3b3dab562c3ff2511461))

## [3.1.0-feeds](https://github.com/GetStream/getstream-go/compare/v3.0.1...v3.1.0-feeds) (2025-08-25)

## [3.0.0-feeds](https://github.com/GetStream/getstream-go/compare/v3.0.1...v3.0.0-feeds) (2025-08-13)

### [3.0.1-feeds](https://github.com/GetStream/getstream-go/compare/v3.0.0...v3.0.1-feeds) (2025-08-13)

## [3.0.0-feeds](https://github.com/GetStream/getstream-go/compare/v3.0.1-feeds...v3.0.0-feeds) (2025-08-05)

## [2.1.0](https://github.com/GetStream/getstream-go/compare/v1.2.0...v2.1.0) (2025-05-06)

## [2.0.0](https://github.com/GetStream/getstream-go/compare/v1.2.0...v2.0.0) (2025-04-30)

## 1.2.0 (2025-02-10)

## 1.2.0 (2025-02-10)

## 1.1.0 (2025-01-13)

## 1.0.0 (2024-12-31)

### 0.0.1: Initial release of the package (2024-10-28)
