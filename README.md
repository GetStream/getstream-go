# Official Go SDK for [Stream](https://getstream.io/)

[![Build Status](https://img.shields.io/github/actions/workflow/status/GetStream/getstream-go/ci.yml?branch=main&style=flat-square)](https://github.com/GetStream/getstream-go/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/GetStream/getstream-go?style=flat-square)](https://goreportcard.com/report/github.com/GetStream/getstream-go)
[![Godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/GetStream/getstream-go)
[![GitHub release](https://img.shields.io/github/release/GetStream/getstream-go.svg?style=flat-square)](https://github.com/GetStream/getstream-go/releases/latest)
[![Go Version](https://img.shields.io/badge/go%20version-%3E%3D1.19-61CFDD.svg?style=flat-square)](https://golang.org/doc/devel/release.html)
[![codecov](https://img.shields.io/codecov/c/github/GetStream/getstream-go.svg?style=flat-square)](https://codecov.io/gh/GetStream/getstream-go)

<p align="center">
    <img src="./assets/logo.svg" width="50%" height="50%">
</p>
<p align="center">
    Official Go API client for Stream Chat and Video, a service for building chat and video applications.
    <br />
    <a href="https://getstream.io/chat/docs/"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/GetStream/getstream-go/issues">Report Bug</a>
    ·
    <a href="https://github.com/GetStream/getstream-go/issues">Request Feature</a>
</p>

## Features

- Video call creation and management
- Chat session creation and management
- Token generation for user authentication
- **Timeout management**: Customizable timeouts for API requests.
- **Rate Limiting**: Rate limits are automatically parsed from responses and accessible in API errors and responses.
- **Customizable Logging with Support for Popular Logging Libraries**

## 📝 About Stream

You can sign up for a Stream account at our [Get Started](https://getstream.io/chat/get_started/) page.

You can use this library to access chat API endpoints server-side.


## ⚙️ Installation

```shell
go get github.com/GetStream/getstream-go
```

## ✨ Getting Started

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"

    stream "github.com/GetStream/getstream-go"
)

func main() {
    apiKey := os.Getenv("STREAM_API_KEY")
    apiSecret := os.Getenv("STREAM_API_SECRET")
    userID := "your-user-id" // Replace with your server user ID

    // Initialize client
    client, err := stream.NewClient(apiKey, apiSecret)
    if err != nil {
        fmt.Printf("Error initializing client: %v\n", err)
        return
    }

    // Or initialize using only environmental variables:
    // (required) STREAM_API_KEY, (required) STREAM_API_SECRET
    client, err = stream.NewClientFromEnvVars()
    if err != nil {
        fmt.Printf("Error initializing client from env vars: %v\n", err)
        return
    }

    // Define a context
    ctx := context.Background()

    // Create a call
    call := client.Video.Call("default", "unique-call-id")

    // Create or get a call
    response, err := call.GetOrCreate(ctx, &stream.CallRequest{
        CreatedBy: stream.UserRequest{ID: userID},
    })
    if err != nil {
        fmt.Printf("Error creating/getting call: %v\n", err)
        return
    }

    fmt.Printf("Call created/retrieved: %s\n", response.Call.ID)

    // Update call settings
    _, err = call.Update(ctx, &stream.UpdateCallRequest{
        SettingsOverride: &stream.CallSettingsRequest{
            Audio: &stream.AudioSettings{
                MicDefaultOn: stream.PtrTo(true),
            },
        },
    })
    if err != nil {
        fmt.Printf("Error updating call settings: %v\n", err)
        return
    }

    // Create a token for client-side use
    token, err := client.CreateToken(userID, nil)
    if err != nil {
        fmt.Printf("Error creating token: %v\n", err)
        return
    }

    fmt.Printf("Token for user %s: %s\n", userID, token)

    // Query calls
    callsResponse, err := client.Video.QueryCalls(ctx, &stream.QueryCallsRequest{
        FilterConditions: map[string]interface{}{
            "created_by_user_id": userID,
        },
    })
    if err != nil {
        fmt.Printf("Error querying calls: %v\n", err)
        return
    }

    fmt.Printf("Found %d calls\n", len(callsResponse.Calls))
}

// Helper function to create a pointer to a value
func PtrTo[T any](v T) *T {
    return &v
}
```

## 🚦 Rate Limits

Rate limit information is parsed from the response headers and is accessible in both API errors and successful responses. This helps you manage your API usage efficiently and avoid being throttled.

Rate limits can be accessed from:
- **Errors**: The `StreamError` struct includes `RateLimitInfo` in case of rate-limited requests.
- **Responses**: Successful responses also contain `RateLimitInfo` under the `RateLimitInfo` field.

## 🚀 Timeout Configuration

You can customize the timeout for all API requests. By default, the client has a timeout of 6 seconds, but you can set a custom timeout using the `WithTimeout` option during the client initialization.

```go
client, err := stream.NewClient(apiKey, apiSecret, stream.WithTimeout(10*time.Second))
if err != nil {
    fmt.Printf("Error initializing client: %v\n", err)
    return
}
```

Alternatively, you can configure the timeout via the `STREAM_HTTP_TIMEOUT` environment variable:

```bash
export STREAM_HTTP_TIMEOUT=10
```


## 📄 Custom Logging

The SDK provides flexible logging capabilities by allowing you to pass your own logger. This means you can integrate your preferred logging library (e.g., `logrus`, `zap`, `zerolog`) with the SDK to have consistent logging across your application.

### **Using a Custom Logger**

To use a custom logger with the SDK, you need to implement the `Logger` interface provided by the SDK and pass it when initializing the client.

#### **Logger Interface**

```go
type Logger interface {
    Debug(format string, v ...interface{})
    Info(format string, v ...interface{})
    Warn(format string, v ...interface{})
    Error(format string, v ...interface{})
}
```

#### **Example with `logrus`**

```go
package main

import (
    "context"
    "fmt"
    "os"

    stream "github.com/GetStream/getstream-go"
    "github.com/sirupsen/logrus"
)

type LogrusLogger struct {
    logger *logrus.Logger
}

func (l *LogrusLogger) Debug(format string, v ...interface{}) {
    l.logger.Debugf(format, v...)
}

func (l *LogrusLogger) Info(format string, v ...interface{}) {
    l.logger.Infof(format, v...)
}

func (l *LogrusLogger) Warn(format string, v ...interface{}) {
    l.logger.Warnf(format, v...)
}

func (l *LogrusLogger) Error(format string, v ...interface{}) {
    l.logger.Errorf(format, v...)
}

func main() {
    apiKey := os.Getenv("STREAM_API_KEY")
    apiSecret := os.Getenv("STREAM_API_SECRET")

    // Initialize logrus logger
    logrusLogger := logrus.New()
    logrusLogger.SetLevel(logrus.DebugLevel)
    logrusLogger.SetFormatter(&logrus.TextFormatter{
        FullTimestamp: true,
    })

    // Wrap logrus.Logger with LogrusLogger
    logger := &LogrusLogger{logger: logrusLogger}

    // Initialize client with custom logger
    client, err := stream.NewClient(apiKey, apiSecret, stream.WithLogger(logger))
    if err != nil {
        fmt.Printf("Error initializing client: %v\n", err)
        return
    }

    // Use the client
    ctx := context.Background()
    // ... your code ...

    client.Logger().Info("Custom logger integrated with SDK")
}
```

#### **Example with `zap`**

```go
package main

import (
    "context"
    "fmt"
    "os"

    stream "github.com/GetStream/getstream-go"
    "go.uber.org/zap"
)

type ZapLogger struct {
    logger *zap.SugaredLogger
}

func (l *ZapLogger) Debug(format string, v ...interface{}) {
    l.logger.Debugf(format, v...)
}

func (l *ZapLogger) Info(format string, v ...interface{}) {
    l.logger.Infof(format, v...)
}

func (l *ZapLogger) Warn(format string, v ...interface{}) {
    l.logger.Warnf(format, v...)
}

func (l *ZapLogger) Error(format string, v ...interface{}) {
    l.logger.Errorf(format, v...)
}

func main() {
    apiKey := os.Getenv("STREAM_API_KEY")
    apiSecret := os.Getenv("STREAM_API_SECRET")

    // Create a zap configuration
    config := zap.NewProductionConfig()
    config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

    // Build the logger
    zapLogger, err := config.Build()
    if err != nil {
        fmt.Printf("Error initializing zap logger: %v\n", err)
        return
    }
    defer zapLogger.Sync()

    // Wrap zap.SugaredLogger with ZapLogger
    logger := &ZapLogger{logger: zapLogger.Sugar()}

    // Initialize client with custom logger
    client, err := stream.NewClient(apiKey, apiSecret, stream.WithLogger(logger))
    if err != nil {
        fmt.Printf("Error initializing client: %v\n", err)
        return
    }

    // Use the client
    ctx := context.Background()
    // ... your code ...

    client.Logger().Info("Custom zap logger integrated with SDK")
}
```

#### **Example with `zerolog`**

```go
package main

import (
    "context"
    "fmt"
    "os"

    stream "github.com/GetStream/getstream-go"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

type ZerologLogger struct {
    logger zerolog.Logger
}

func (l *ZerologLogger) Debug(format string, v ...interface{}) {
    l.logger.Debug().Msgf(format, v...)
}

func (l *ZerologLogger) Info(format string, v ...interface{}) {
    l.logger.Info().Msgf(format, v...)
}

func (l *ZerologLogger) Warn(format string, v ...interface{}) {
    l.logger.Warn().Msgf(format, v...)
}

func (l *ZerologLogger) Error(format string, v ...interface{}) {
    l.logger.Error().Msgf(format, v...)
}

func main() {
    apiKey := os.Getenv("STREAM_API_KEY")
    apiSecret := os.Getenv("STREAM_API_SECRET")

    // Initialize zerolog logger
    zerolog.SetGlobalLevel(zerolog.DebugLevel)
    logger := &ZerologLogger{logger: log.Logger}

    // Initialize client with custom logger
    client, err := stream.NewClient(apiKey, apiSecret, stream.WithLogger(logger))
    if err != nil {
        fmt.Printf("Error initializing client: %v\n", err)
        return
    }

    // Use the client
    ctx := context.Background()
    // ... your code ...

    client.Logger().Info("Custom zerolog logger integrated with SDK")
}
```

### **Setting Log Levels in Custom Loggers**

The SDK delegates log level management to the custom logger implementations. Here's how you can set log levels in different logging libraries:

- **logrus**: Use `SetLevel` method.
  ```go
  logrusLogger.SetLevel(logrus.InfoLevel)
  ```

- **zap**: Configure the log level in the configuration before building the logger.
  ```go
  config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
  ```

- **zerolog**: Use `SetGlobalLevel` or set the level on the logger instance.
  ```go
  zerolog.SetGlobalLevel(zerolog.WarnLevel)
  ```

### **Using the Default Logger**

If you don't provide a custom logger, the SDK uses a default logger that writes to `os.Stderr`. You can configure the default logger's log level:

```go
stream.SetDefaultLogLevel(stream.LogLevelDebug)
```

## ✍️ Contributing

We welcome code changes that improve this library or fix a problem, please make sure to follow all best practices and add tests if applicable before submitting a Pull Request on Github. We are very happy to merge your code in the official repository. Make sure to sign our [Contributor License Agreement (CLA)](https://docs.google.com/forms/d/e/1FAIpQLScFKsKkAJI7mhCr7K9rEIOpqIDThrWxuvxnwUq2XkHyG154vQ/viewform) first. See our [license file](./LICENSE) for more details.

Head over to [CONTRIBUTING.md](./CONTRIBUTING.md) for some development tips.

### Generate Code from Spec

To regenerate the Go source from OpenAPI, just run the `./generate.sh` script from this repo.

> **Note**
> Code generation currently relies on tooling that is not publicly available. Only Stream developers can regenerate SDK source code from the OpenAPI spec.

## 🧑‍💻 We Are Hiring!

We've recently closed a [$38 million Series B funding round](https://techcrunch.com/2021/03/04/stream-raises-38m-as-its-chat-and-activity-feed-apis-power-communications-for-1b-users/) and we keep actively growing.
Our APIs are used by more than a billion end-users, and you'll have a chance to make a huge impact on the product within a team of the strongest engineers all over the world.

Check out our current openings and apply via [Stream's website](https://getstream.io/team/#jobs).
