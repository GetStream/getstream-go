# Official Go SDK for [Stream](https://getstream.io/)


<div align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://github.com/user-attachments/assets/19dd203e-c84a-4015-9c90-1a54212fc2e2">
    <source media="(prefers-color-scheme: light)" srcset="https://github.com/user-attachments/assets/fb6b6686-ce5d-4c8f-87b7-9bb495b6ce66">
    <img src="https://github.com/user-attachments/assets/19dd203e-c84a-4015-9c90-1a54212fc2e2" width="360" alt="Stream logo">
  </picture>
  <h3>Official Go API client for Stream Chat and Video, a service for building chat and video applications.</h3>
  
[![Build Status](https://img.shields.io/github/actions/workflow/status/GetStream/getstream-go/ci.yml?branch=main&style=flat-square)](https://github.com/GetStream/getstream-go/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/GetStream/getstream-go?style=flat-square)](https://goreportcard.com/report/github.com/GetStream/getstream-go)
[![Godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/GetStream/getstream-go)
[![GitHub release](https://img.shields.io/github/release/GetStream/getstream-go.svg?style=flat-square)](https://github.com/GetStream/getstream-go/releases/latest)
[![Go Version](https://img.shields.io/badge/go%20version-%3E%3D1.19-61CFDD.svg?style=flat-square)](https://golang.org/doc/devel/release.html)
[![codecov](https://img.shields.io/codecov/c/github/GetStream/getstream-go.svg?style=flat-square)](https://codecov.io/gh/GetStream/getstream-go)
</div>

<p align="center">
    <a href="https://getstream.io/chat/docs/"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/GetStream/getstream-go/issues">Report Bug</a>
    ·
    <a href="https://github.com/GetStream/getstream-go/issues">Request Feature</a>
</p>

## **Quick Links**

- [Register](https://getstream.io/chat/trial/) to get an API key for Stream
- [Docs](https://getstream.io/video/docs/api/)

## What is Stream?

Stream allows developers to rapidly deploy scalable feeds, chat messaging and video with an industry leading 99.999% uptime SLA guarantee.

Stream provides UI components and state handling that make it easy to build video calling for your app. All calls run on Stream's network of edge servers around the world, ensuring optimal latency and reliability.

## 👩‍💻 Free for Makers 👨‍💻

Stream is free for most side and hobby projects. To qualify, your project/company needs to have < 5 team members and < $10k in monthly revenue. Makers get $100 in monthly credit for video for free.

## 😎 Repo Overview 😎

This repo contains the Golang server-side SDK developed by the team and Stream community. For a feature overview please visit our [roadmap](https://github.com/GetStream/protocol/discussions/177).

## ✍️ Contributing

We welcome code changes that improve this library or fix a problem, please make sure to follow all best practices and add tests if applicable before submitting a Pull Request on Github. We are very happy to merge your code in the official repository. Make sure to sign our [Contributor License Agreement (CLA)](https://docs.google.com/forms/d/e/1FAIpQLScFKsKkAJI7mhCr7K9rEIOpqIDThrWxuvxnwUq2XkHyG154vQ/viewform) first. See our [license file](./LICENSE) for more details.

Head over to [CONTRIBUTING.md](./CONTRIBUTING.md) for some development tips.

### Generate Code from Spec

To regenerate the Go source from OpenAPI, just run the `./generate.sh` script from this repo.

> **Note**
> Code generation currently relies on tooling that is not publicly available. Only Stream developers can regenerate SDK source code from the OpenAPI spec.

## 🔍 Why Choose Stream?

[Stream](https://getstream.io) powers chat and activity feeds for over 1 billion end users. Our robust, scalable platform helps you build:

- **In-app Messaging** - Group chats, direct messaging, channels
- **Team Collaboration** - Slack-like workspaces and threaded conversations
- **Live Streaming** - Interactive live streams with chat
- **Virtual Events** - Engaging event platforms with rich chat features
- **Gaming Communities** - Community building with real-time chat
- **Support Channels** - Customer support with integrated chat

## 🧑‍💻 We Are Hiring!

We've recently closed a [$38 million Series B funding round](https://techcrunch.com/2021/03/04/stream-raises-38m-as-its-chat-and-activity-feed-apis-power-communications-for-1b-users/) and we keep actively growing.
Our APIs are used by more than a billion end-users, and you'll have a chance to make a huge impact on the product within a team of the strongest engineers all over the world.

Check out our current openings and apply via [Stream's website](https://getstream.io/team/#jobs).
