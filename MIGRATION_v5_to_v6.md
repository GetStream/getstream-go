# Migrating from v5 to v6

v6 changes the module path, six generated types, and every optional array and map on a request. There are no removed endpoints, types or fields.

## Module path

```bash
go get github.com/GetStream/getstream-go/v6
```

Update every import:

```go
- import "github.com/GetStream/getstream-go/v5"
+ import "github.com/GetStream/getstream-go/v6"
```

## `ChannelMemberRequest.UserID` is now optional

`UserID` changed from `string` to `*string`. Use the `PtrTo` helper:

```go
- members := []ChannelMemberRequest{{UserID: userID}}
+ members := []ChannelMemberRequest{{UserID: PtrTo(userID)}}
```

## `ChannelMemberRequest.User` takes a request type

The field changed from `*UserResponse` to `*MemberUserRequest`, a new type carrying only the fields a request may set. If you were passing a `UserResponse` you received from the API, build a `MemberUserRequest` instead and set the fields you mean to change.

```go
- ChannelMemberRequest{User: &UserResponse{ID: id, Name: PtrTo("Alice")}}
+ ChannelMemberRequest{User: &MemberUserRequest{ID: id, Name: PtrTo("Alice")}}
```

## Reminder events carry a value, not a pointer

`Reminder` changed from `*ReminderResponseData` to `ReminderResponseData` on `ReminderCreatedEvent`, `ReminderUpdatedEvent`, `ReminderDeletedEvent` and `ReminderNotificationEvent`. Drop the nil check and the dereference:

```go
- if event.Reminder != nil { use(*event.Reminder) }
+ use(event.Reminder)
```

## Optional arrays and maps on requests are now pointers

Every non-required array or map field on a request type changed from `[]T` to `*[]T` and from `map[string]V` to `*map[string]V`, and gained `,omitempty`. Required fields are unchanged.

This fixes silent data loss. A bare slice with no `omitempty` marshals to `null` when you leave it unset, and `UpdateChannelType` reads that `null` as "clear", so omitting `Commands` wiped every command on the channel type. The same applied to `Blocklists` and `AllowedFlagReasons`.

The pointer lets you say all three things:

```go
// keep the current value: omit the field
client.Chat().UpdateChannelType(ctx, "messaging", &UpdateChannelTypeRequest{
    Automod: "simple",
})

// clear it: point at an empty slice
client.Chat().UpdateChannelType(ctx, "messaging", &UpdateChannelTypeRequest{
    Commands: PtrTo([]string{}),
})

// replace it
client.Chat().UpdateChannelType(ctx, "messaging", &UpdateChannelTypeRequest{
    Commands: PtrTo([]string{"giphy", "ban"}),
})
```

To migrate, wrap the value in `PtrTo`:

```go
- &QueryChannelsRequest{FilterConditions: map[string]any{"type": "messaging"}}
+ &QueryChannelsRequest{FilterConditions: PtrTo(map[string]any{"type": "messaging"})}
```

The compiler finds every site for you. If a field was set to an empty slice only to satisfy the type, drop the field instead: that now means "keep", which is almost always what the code wanted.
