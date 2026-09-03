# Migrating from v6 to v7

v7 changes the module path. `CreateReminder` now returns `CreateReminderResponse` with a nested `Reminder` instead of `ReminderResponseData` at the top level.

## Module path

```bash
go get github.com/GetStream/getstream-go/v7
```

```go
- import "github.com/GetStream/getstream-go/v6"
+ import "github.com/GetStream/getstream-go/v7"
```

## CreateReminder wrapper

```go
- createResp.Data.MessageID
+ createResp.Data.Reminder.MessageID
```
