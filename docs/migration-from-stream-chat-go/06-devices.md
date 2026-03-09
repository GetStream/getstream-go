# Devices

This guide shows how to migrate device (push notification) operations from `github.com/GetStream/stream-chat-go/v8` to `github.com/GetStream/getstream-go/v4`.

## Add a Device

**Before (stream-chat-go):**

```go
package main

import (
	"context"

	stream "github.com/GetStream/stream-chat-go/v8"
)

func main() {
	client, _ := stream.NewClient("your-api-key", "your-api-secret")
	ctx := context.Background()

	resp, err := client.AddDevice(ctx, &stream.Device{
		ID:           "device-token-123",
		UserID:       "user-123",
		PushProvider: stream.PushProviderFirebase,
	})
}
```

**After (getstream-go):**

```go
package main

import (
	"context"

	"github.com/GetStream/getstream-go/v4"
)

func main() {
	client, _ := getstream.NewClient("your-api-key", "your-api-secret")
	ctx := context.Background()

	resp, err := client.CreateDevice(ctx, &getstream.CreateDeviceRequest{
		ID:           "device-token-123",
		UserID:       getstream.PtrTo("user-123"),
		PushProvider: "firebase",
	})
}
```

**Key changes:**
- `AddDevice` becomes `CreateDevice`
- `Device` struct becomes `CreateDeviceRequest`
- `UserID` is a pointer field
- `PushProvider` is a plain string (`"firebase"`, `"apn"`) instead of a typed constant

## Add an APN Device

**Before (stream-chat-go):**

```go
package main

import (
	"context"

	stream "github.com/GetStream/stream-chat-go/v8"
)

func main() {
	client, _ := stream.NewClient("your-api-key", "your-api-secret")
	ctx := context.Background()

	resp, err := client.AddDevice(ctx, &stream.Device{
		ID:           "apn-device-token",
		UserID:       "user-123",
		PushProvider: stream.PushProviderAPN,
	})
}
```

**After (getstream-go):**

```go
package main

import (
	"context"

	"github.com/GetStream/getstream-go/v4"
)

func main() {
	client, _ := getstream.NewClient("your-api-key", "your-api-secret")
	ctx := context.Background()

	resp, err := client.CreateDevice(ctx, &getstream.CreateDeviceRequest{
		ID:           "apn-device-token",
		UserID:       getstream.PtrTo("user-123"),
		PushProvider: "apn",
	})
}
```

**Key changes:**
- `PushProviderAPN` constant becomes the string `"apn"`

## List Devices

**Before (stream-chat-go):**

```go
package main

import (
	"context"

	stream "github.com/GetStream/stream-chat-go/v8"
)

func main() {
	client, _ := stream.NewClient("your-api-key", "your-api-secret")
	ctx := context.Background()

	resp, err := client.GetDevices(ctx, "user-123")
	// resp.Devices is []*Device
}
```

**After (getstream-go):**

```go
package main

import (
	"context"

	"github.com/GetStream/getstream-go/v4"
)

func main() {
	client, _ := getstream.NewClient("your-api-key", "your-api-secret")
	ctx := context.Background()

	resp, err := client.ListDevices(ctx, &getstream.ListDevicesRequest{
		UserID: getstream.PtrTo("user-123"),
	})
	// resp.Data.Devices
}
```

**Key changes:**
- `GetDevices` becomes `ListDevices`
- User ID moves from a positional argument into `ListDevicesRequest.UserID`
- Response accessed via `resp.Data.Devices`

## Delete a Device

**Before (stream-chat-go):**

```go
package main

import (
	"context"

	stream "github.com/GetStream/stream-chat-go/v8"
)

func main() {
	client, _ := stream.NewClient("your-api-key", "your-api-secret")
	ctx := context.Background()

	resp, err := client.DeleteDevice(ctx, "user-123", "device-token-123")
}
```

**After (getstream-go):**

```go
package main

import (
	"context"

	"github.com/GetStream/getstream-go/v4"
)

func main() {
	client, _ := getstream.NewClient("your-api-key", "your-api-secret")
	ctx := context.Background()

	resp, err := client.DeleteDevice(ctx, &getstream.DeleteDeviceRequest{
		ID:     "device-token-123",
		UserID: getstream.PtrTo("user-123"),
	})
}
```

**Key changes:**
- Positional arguments become fields on `DeleteDeviceRequest`
- `UserID` is a pointer field

## Method Mapping Summary

| Legacy (stream-chat-go) | New (getstream-go) |
|---|---|
| `client.AddDevice(ctx, &Device{...})` | `client.CreateDevice(ctx, &CreateDeviceRequest{...})` |
| `client.GetDevices(ctx, userID)` | `client.ListDevices(ctx, &ListDevicesRequest{UserID: ...})` |
| `client.DeleteDevice(ctx, userID, deviceID)` | `client.DeleteDevice(ctx, &DeleteDeviceRequest{ID: ..., UserID: ...})` |
