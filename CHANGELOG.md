# Changelog

All notable changes to this project will be documented in this file. See [standard-version](https://github.com/conventional-changelog/standard-version) for commit guidelines.

### [4.0.6](https://github.com/GetStream/getstream-go/compare/v4.0.4...v4.0.6) (2026-04-23)


### Features

* add bookmark for comments ([#97](https://github.com/GetStream/getstream-go/issues/97)) ([7e41662](https://github.com/GetStream/getstream-go/commit/7e41662f20af4df9bf30040da2448c44d594b8e8))
* regenerate from latest OpenAPI spec, keep only retention runs test ([5d543c4](https://github.com/GetStream/getstream-go/commit/5d543c48dc4b14fb3d95f538031f0e3029c2618c))
* update opnepai specs ([3f28e35](https://github.com/GetStream/getstream-go/commit/3f28e35d9eb9c193222724e748f92e41d29e2f4c))


### Bug Fixes

* skip retention tests when endpoints unavailable (404) ([faff5e6](https://github.com/GetStream/getstream-go/commit/faff5e66c79c1f30a7adddefa8d59c7ad1c25bbb))

### [4.0.4](https://github.com/GetStream/getstream-go/compare/v4.0.3...v4.0.4) (2026-03-31)

### [4.0.3](https://github.com/GetStream/getstream-go/compare/v4.0.2...v4.0.3) (2026-03-23)

### [4.0.2](https://github.com/GetStream/getstream-go/compare/v4.0.1...v4.0.2) (2026-03-19)


### Bug Fixes

* add omitempty on pointer types ([#87](https://github.com/GetStream/getstream-go/issues/87)) ([e98f4c9](https://github.com/GetStream/getstream-go/commit/e98f4c9b020dc854b53a728297a8674713da0d16))

### [4.0.1](https://github.com/GetStream/getstream-go/compare/v4.0.0...v4.0.1) (2026-03-19)

## [4.0.0](https://github.com/GetStream/getstream-go/compare/v3.9.0...v4.0.0) (2026-03-05)

### Breaking Changes

- Type names across all products now follow the OpenAPI spec naming convention: response types are suffixed with `Response`, input types with `Request`. See [MIGRATION_v3_to_v4.md](./MIGRATION_v3_to_v4.md) for the complete rename mapping.
- `Event` (WebSocket envelope type) renamed to `WSEvent`. Base event type renamed from `BaseEvent` to `Event` (with field `type` instead of `T`).
- Event composition changed from monolithic `*Preset` embeds to modular `Has*` types (`HasChannel`, `HasMessage`, `HasUserCommonFields`, etc.).
- `Pager` renamed to `PagerResponse` and migrated from offset-based to cursor-based pagination (`next`/`prev` tokens).
- Module path changed from `github.com/GetStream/getstream-go/v4` to `github.com/GetStream/getstream-go/v4`.

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
