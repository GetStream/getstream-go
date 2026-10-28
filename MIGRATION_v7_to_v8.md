# Migrating from v7 to v8

v8 regenerates the OpenAPI client from `openapi-v239.93.2` (FLU-1105). JSON on the wire is unchanged; several request and response fields are named enum types instead of `string` / `*string`.

## Module path

```bash
go get github.com/GetStream/getstream-go/v8
```

```go
- import "github.com/GetStream/getstream-go/v7"
+ import "github.com/GetStream/getstream-go/v8"
```

## Typed enum fields

Fields such as `PushLevel`, `VotingVisibility`, `PushProviderName`, and `BlocklistBehavior` are distinct types. Use the generated constants or convert with `getstream.PushLevel("all")` where a plain string worked in v7.

See [CHANGELOG.md](./CHANGELOG.md) for the full breaking-change entry when v8.0.0 is released.
