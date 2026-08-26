# Migrating from v5 to v6

v6 changes the module path and six generated types. There are no removed endpoints, types or fields.

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
