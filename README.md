# Official Go SDK for [Stream](https://getstream.io/)

[![build](https://github.com/GetStream/getstream-go/workflows/build/badge.svg)](https://github.com/GetStream/getstream-go/actions)
[![godoc](https://pkg.go.dev/badge/GetStream/getstream-go)](https://pkg.go.dev/github.com/GetStream/getstream-go?tab=doc)

<p align="center">
    <img src="./assets/logo.svg" width="50%" height="50%">
</p>
<p align="center">
    Official Go API client for Stream Chat and Video, a service for building chat and video applications.
    <br />
    <a href="https://getstream.io/chat/docs/"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/GetStream/stream-go/issues">Report Bug</a>
    ·
    <a href="https://github.com/GetStream/stream-go/issues">Request Feature</a>
</p>

## Features

- Video call creation and management
- Chat session creation and management
- Token generation for user authentication

## 📝 About Stream

You can sign up for a Stream account at our [Get Started](https://getstream.io/chat/get_started/) page.

You can use this library to access chat API endpoints server-side.

For the client-side integrations (web and mobile) have a look at the JavaScript, iOS and Android SDK libraries ([docs](https://getstream.io/chat/)).

## ⚙️ Installation

```shell
go get github.com/GetStream/stream-go
```

## ✨ Getting started

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
	userID := "your-user-id" // Replace with your server user id

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

## ✍️ Contributing

We welcome code changes that improve this library or fix a problem, please make sure to follow all best practices and add tests if applicable before submitting a Pull Request on Github. We are very happy to merge your code in the official repository. Make sure to sign our [Contributor License Agreement (CLA)](https://docs.google.com/forms/d/e/1FAIpQLScFKsKkAJI7mhCr7K9rEIOpqIDThrWxuvxnwUq2XkHyG154vQ/viewform) first. See our [license file](./LICENSE) for more details.

Head over to [CONTRIBUTING.md](./CONTRIBUTING.md) for some development tips.


### Generate code from spec

To regenerate the Go source from OpenAPI, just run the `./generate.sh` script from this repo.

> [!NOTE]
> Code generation currently relies on tooling that is not publicly available, only Stream devs can regenerate SDK source code from the OpenAPI spec.


## 🧑‍💻 We are hiring!

We've recently closed a [$38 million Series B funding round](https://techcrunch.com/2021/03/04/stream-raises-38m-as-its-chat-and-activity-feed-apis-power-communications-for-1b-users/) and we keep actively growing.
Our APIs are used by more than a billion end-users, and you'll have a chance to make a huge impact on the product within a team of the strongest engineers all over the world.

Check out our current openings and apply via [Stream's website](https://getstream.io/team/#jobs).
